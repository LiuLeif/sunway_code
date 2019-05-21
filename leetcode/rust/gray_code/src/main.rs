// Accepted
// @lc id=89 lang=rust
// problem: gray_code
struct Solution {}
impl Solution {
    pub fn gray_code(n: i32) -> Vec<i32> {
        let mut ret = vec![0];
        for i in 0..n {
            ret.extend(ret.iter().rev().map(|x| x | 1 << i).collect::<Vec<i32>>());
        }
        ret
    }
}
fn main() {
    println!("{:?}", Solution::gray_code(2));
}
