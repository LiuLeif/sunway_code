// 2019-09-04 15:14
use lazy_static::*;
use rand::prelude::*;
use std::thread;
use std::time::Duration;
use zmq;

lazy_static! {
    static ref CONTEXT: zmq::Context = zmq::Context::new();
}

fn start_worker(identity: i32) {
    let socket = CONTEXT.socket(zmq::PULL).unwrap();
    assert!(socket.connect("inproc://balancer").is_ok());

    loop {
        let msg = {
            let mut msg = zmq::Message::new();
            socket.recv(&mut msg, 0).unwrap();
            msg
        };

        println!("{:?} in {}", msg.as_str().unwrap(), identity);
        // NOTE: PULL socket 无法用来 send
        // socket
        //     .send(format!("done from {}", identity).as_bytes(), 0)
        //     .unwrap();
    }
}

fn start_balancer() {
    let backend = CONTEXT.socket(zmq::PUSH).expect("server backend");
    let frontend = CONTEXT.socket(zmq::DEALER).expect("server backend");
    assert!(backend.bind("inproc://balancer").is_ok());
    assert!(frontend.bind("inproc://task").is_ok());

    loop {
        let request = frontend.recv_string(0).unwrap().unwrap();
        backend.send(request.as_str(), 0);
    }
}

fn start_client() {
    let socket = CONTEXT.socket(zmq::DEALER).expect("server socket");
    assert!(socket.connect("inproc://task").is_ok());

    // 直接在 client 使用 PUSH 是一样的效果
    // let socket = CONTEXT.socket(zmq::PUSH).expect("server socket");
    // assert!(socket.bind("inproc://balancer").is_ok());

    for i in 0..20 {
        socket.send("request", 0);
    }
}

pub fn run() {
    let workers = (0..5)
        .map(|i| {
            thread::spawn(move || {
                start_worker(i);
            })
        })
        .collect::<Vec<_>>();

    let balancer = thread::spawn(|| {
        start_balancer();
    });

    thread::sleep(Duration::from_secs(1));

    let client = thread::spawn(|| {
        start_client();
    });

    workers.into_iter().for_each(|x| {
        x.join();
    });

    balancer.join();
}
