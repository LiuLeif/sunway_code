mod pub_sub;
mod req_rep;
mod dealer_dealer;
mod dealer_router;
mod load_balance;

fn main() {
    // req_rep::run();
    // pub_sub::run();
    // dealer_dealer::run();
    // dealer_router::run();
    load_balance::run();
}
