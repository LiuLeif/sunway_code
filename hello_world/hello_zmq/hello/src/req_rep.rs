// 2019-09-04 15:14
use lazy_static::*;
use std::thread;
use zmq;

lazy_static! {
    static ref CONTEXT: zmq::Context = zmq::Context::new();
}

fn start_client() {
    let socket = CONTEXT.socket(zmq::REQ).unwrap();
    // assert!(socket.connect("ipc:///tmp/hello.ipc").is_ok());
    assert!(socket.connect("inproc://hello").is_ok());

    let mut msg = zmq::Message::new();

    socket.send("hello", 0).expect("send hello from client");
    socket.recv(&mut msg, 0).expect("recv from client");

    println!("received {}", msg.as_str().unwrap());
}

fn start_server() {
    let socket = CONTEXT.socket(zmq::REP).expect("server socket");
    // assert!(socket.bind("ipc:///tmp/hello.ipc").is_ok());
    assert!(socket.bind("inproc://hello").is_ok());

    let mut msg = zmq::Message::new();
    socket.recv(&mut msg, 0).unwrap();
    socket
        .send("hello from server", 0)
        .expect("send from server");
}

pub fn run() {
    let client = thread::spawn(|| {
        start_client();
    });

    let server = thread::spawn(|| {
        start_server();
    });

    client.join();
    server.join();
}
