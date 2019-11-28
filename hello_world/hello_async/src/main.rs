use futures::executor::block_on;
//  _________
// < trivial >
//  ---------
//         \   ^__^
//          \  (oo)\_______
//             (__)\       )\/\
//                 ||----w |
//                 ||     ||
async fn hello() -> () {
    println!("hello");
    world().await;
}

async fn world() -> () {
    println!("world");
}

fn trivial() -> () {
    println!("------ trivial ------");
    block_on(hello());
}
//  ______
// < poll >
//  ------
//         \   ^__^
//          \  (oo)\_______
//             (__)\       )\/\
//                 ||----w |
//                 ||     ||

async fn process_data() -> () {
    let data = poll_data().await;
    println!("{:?}", data);
}

fn poll() -> () {
    println!("------ poll ------");
    block_on(accept());
}

fn main() {
    // trivial();
    poll();
}
