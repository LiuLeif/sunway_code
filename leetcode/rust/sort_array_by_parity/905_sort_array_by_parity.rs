// 2018-12-25 11:29
// @lc id=905 lang=rust
struct Solution {}
impl Solution {
    pub fn sort_array_by_parity(a: Vec<i32>) -> Vec<i32> {
        let mut a = a;
        a.sort_by(|x, y| (x % 2).cmp(&(y % 2)));
        a
    }
}

fn main() {
    println!("{:?}", Solution::sort_array_by_parity(vec![3, 1, 2, 4]));
}
