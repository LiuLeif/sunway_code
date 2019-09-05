// 2019-09-04 15:14
use lazy_static::*;
use std::thread;
use zmq;

lazy_static! {
    static ref CONTEXT: zmq::Context = zmq::Context::new();
}

fn start_client() {
    let socket = CONTEXT.socket(zmq::DEALER).unwrap();
    // assert!(socket.connect("ipc:///tmp/hello.ipc").is_ok());
    assert!(socket.connect("inproc://hello").is_ok());

    let mut msg = zmq::Message::new();

    // dealer - dealer 不需要像 rep - req 那样严格的遵守 `send - reply` 或
    // `reply - send` 的模式, 任意一方 dealer 都可以任意 send 或 reply, 所以
    // dealer 被称为异步的, 而 req, rep 称为同步
    socket.send("hello", 0).expect("send hello from client");
    socket.send("hello", 0).expect("send hello from client");
    socket.recv(&mut msg, 0).expect("recv from client");

    println!("received {} from server", msg.as_str().unwrap());
}

fn start_server() {
    let socket = CONTEXT.socket(zmq::DEALER).expect("server socket");
    // assert!(socket.bind("ipc:///tmp/hello.ipc").is_ok());
    assert!(socket.bind("inproc://hello").is_ok());

    let mut msg = zmq::Message::new();
    socket.recv(&mut msg, 0).unwrap();
    println!("received {} from client", msg.as_str().unwrap());
    socket.recv(&mut msg, 0).unwrap();
    println!("received {} from client", msg.as_str().unwrap());

    socket.send("hello", 0).expect("send from server");
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
