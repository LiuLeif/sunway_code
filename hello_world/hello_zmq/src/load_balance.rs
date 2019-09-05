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
    let socket = CONTEXT.socket(zmq::DEALER).unwrap();
    socket
        .set_identity(identity.to_string().as_bytes())
        .unwrap();
    assert!(socket.connect("inproc://balancer").is_ok());

    loop {
        let msg = {
            let mut msg = zmq::Message::new();
            socket.recv(&mut msg, 0).unwrap();
            msg
        };

        socket
            .send(format!("done from {}", identity).as_bytes(), 0)
            .unwrap();
    }
}

fn start_balancer() {
    let backend = CONTEXT.socket(zmq::ROUTER).expect("server backend");
    let frontend = CONTEXT.socket(zmq::REP).expect("server backend");
    assert!(backend.bind("inproc://balancer").is_ok());
    assert!(frontend.bind("inproc://task").is_ok());

    loop {
        let mut items = [
            frontend.as_poll_item(zmq::POLLIN),
            backend.as_poll_item(zmq::POLLIN),
        ];
        zmq::poll(&mut items, -1).expect("failed polling");

        if items[0].is_readable() {
            let request = frontend.recv_string(0).unwrap().unwrap();

            let mut rng = rand::thread_rng();
            let identity = rng.gen_range(0, 5);

            backend.send(identity.to_string().as_str(), zmq::SNDMORE);
            backend.send(request.as_str(), 0);
        }
        if items[1].is_readable() {
            let response = backend.recv_string(0).unwrap().unwrap();
            let response = backend.recv_string(0).unwrap().unwrap();
            frontend.send(response.as_str(), 0);
        }
    }
}

fn start_client() {
    let socket = CONTEXT.socket(zmq::REQ).expect("server socket");
    assert!(socket.connect("inproc://task").is_ok());

    for i in 0..20 {
        socket.send("request", 0);

        let msg = {
            let mut msg = zmq::Message::new();
            socket.recv(&mut msg, 0);
            msg
        };

        println!("cleint: received {} from balancer", msg.as_str().unwrap());
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
