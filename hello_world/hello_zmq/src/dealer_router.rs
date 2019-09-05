// 2019-09-04 15:14
use lazy_static::*;
use std::thread;
use zmq;

lazy_static! {
    static ref CONTEXT: zmq::Context = zmq::Context::new();
}

// router 用来解决一个 server 对应多个 client 的情形:
// router 作为 server, dealer 作为 client
// router -> dealer 时, 需要 router 自己附加一个 identity, 以便 frame 会发给特定的 dealer
// dealer -> router 时, 会自动加上自己的 identity, 以便 router 能识别出 frame 来自哪个 dealer

fn start_client() {
    let socket = CONTEXT.socket(zmq::DEALER).unwrap();
    socket.set_identity(b"A");
    assert!(socket.connect("inproc://hello").is_ok());

    socket.send("request", 0).expect("send hello from client");

    let msg = {
        let mut msg = zmq::Message::new();
        socket.recv(&mut msg, 0);
        msg
    };

    println!("received {} from server", msg.as_str().unwrap());
}

fn start_server() {
    let socket = CONTEXT.socket(zmq::ROUTER).expect("server socket");
    assert!(socket.bind("inproc://hello").is_ok());

    let msg = {
        let mut msg = zmq::Message::new();
        socket.recv(&mut msg, 0);
        msg
    };

    let client_identity = msg.as_str().unwrap().to_string();

    let msg = {
        let mut msg = zmq::Message::new();
        socket.recv(&mut msg, 0);
        msg
    };

    println!(
        "received {} from client {}",
        msg.as_str().unwrap(),
        client_identity
    );

    socket.send("A", zmq::SNDMORE);
    socket.send("response", 0);
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
