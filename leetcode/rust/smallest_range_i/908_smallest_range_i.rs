// 2018-12-26 16:02
// @lc id=908 lang=rust
struct Solution {}
impl Solution {
    pub fn smallest_range_i(a: Vec<i32>, k: i32) -> i32 {
        let (lo, hi) = (a.iter().min().unwrap(), a.iter().max().unwrap());
        if lo + k >= hi - k {
            return 0;
        } else {
            return hi - k - lo - k;
        }
    }
}

fn main() {
    println!("{:?}", Solution::smallest_range_i(vec![0, 10], 2));
}
