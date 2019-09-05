mod dealer_dealer;
mod dealer_router;
mod load_balance;
mod pub_sub;
mod push_pull;
mod req_rep;

fn main() {
    // req_rep::run();
    // pub_sub::run();
    // dealer_dealer::run();
    // dealer_router::run();
    // load_balance::run();
    push_pull::run();
}
