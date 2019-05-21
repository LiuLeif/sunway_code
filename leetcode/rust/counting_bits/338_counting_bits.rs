// 2018-12-26 16:53
// @lc id=338 lang=rust
struct Solution {}
impl Solution {
    pub fn count_bits(num: i32) -> Vec<i32> {
        let mut ret = vec![0; num as usize + 1];
        for i in 1..num as usize + 1 {
            ret[i] = ret[i >> 1];
            ret[i] += i as i32 & 1
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::count_bits(5));
}
