// 2019-09-04 15:26
use lazy_static::*;
use std::thread;
use std::time::Duration;
use zmq;

lazy_static! {
    static ref CONTEXT: zmq::Context = zmq::Context::new();
}

fn start_subscriber() {
    let socket = CONTEXT.socket(zmq::SUB).unwrap();
    // assert!(socket.connect("ipc:///tmp/hello.ipc").is_ok());
    assert!(socket.connect("inproc://hello").is_ok());
    socket.set_subscribe("alert".as_bytes());
    let mut msg = zmq::Message::new();

    socket.recv(&mut msg, 0).unwrap();

    println!("received: {}", msg.as_str().unwrap());
}

fn start_publisher() {
    let socket = CONTEXT.socket(zmq::PUB).unwrap();
    assert!(socket.bind("inproc://hello").is_ok());

    let mut msg = zmq::Message::new();
    socket.send("info server", 0).unwrap();
    socket.send("alert server", 0).unwrap();
}

pub fn run() {
    let subscribers = (0..5)
        .map(|_| {
            thread::spawn(|| {
                start_subscriber();
            })
        })
        .collect::<Vec<_>>();

    thread::sleep(Duration::from_secs(1));

    let publisher = thread::spawn(|| {
        start_publisher();
    });

    subscribers.into_iter().for_each(|x| {
        x.join();
    });
    publisher.join();
}
