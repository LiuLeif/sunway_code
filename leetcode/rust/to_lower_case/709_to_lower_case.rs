// 2018-12-24 17:13
// @lc id=709 lang=rust
struct Solution {}
impl Solution {
    pub fn to_lower_case(str: String) -> String {
        return str.to_lowercase();
    }
}

fn main() {
    println!("{:?}",Solution::to_lower_case("Hello".to_owned()));

}
