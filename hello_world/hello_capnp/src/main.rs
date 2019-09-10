// 2019-09-10 12:29
// http://bspeice.github.io/captains-cookbook-part-1.html
// https://github.com/capnproto/capnproto-rust/tree/master/example
use capnp;
use capnp::traits::ToU16;
pub mod data_capnp {
    include! {"../schema/data_capnp.rs"}
}

fn main() {
    let mut builder = capnp::message::Builder::new_default();

    let mut data_msg = builder.init_root::<data_capnp::data::Builder>();
    data_msg.set_x(12);
    data_msg.set_y(14);

    let mut inner_data_msg = data_msg.reborrow().init_inner_data();
    inner_data_msg.set_x("hello");
    inner_data_msg.set_y(10);

    let mut list_data = data_msg.reborrow().init_list_data(1);
    list_data.set(0, 100);

    data_msg.set_color(data_capnp::data::Color::Green);

    //  ___________
    // < serialize >
    //  -----------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    let mut buffer = Vec::new();
    capnp::serialize::write_message(&mut buffer, &builder).unwrap();

    //  _____________
    // < deserialize >
    //  -------------
    //         \   ^__^
    //          \  (oo)\_______
    //             (__)\       )\/\
    //                 ||----w |
    //                 ||     ||
    let deserialized = capnp::serialize::read_message(
        &mut buffer.as_slice(),
        capnp::message::ReaderOptions::new(),
    )
    .unwrap();
    let data_reader = deserialized.get_root::<data_capnp::data::Reader>().unwrap();
    println!("{:?}", data_reader.get_x());
    println!("{:?}", data_reader.get_inner_data().unwrap().get_x());
    println!("{:?}", data_reader.get_list_data().unwrap().get(0));
    println!("{:?}", data_reader.get_color().unwrap().to_u16());
}
