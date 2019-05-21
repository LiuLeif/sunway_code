// 2018-12-25 15:28
// @lc id=461 lang=rust
struct Solution {}
impl Solution {
    pub fn hamming_distance(x: i32, y: i32) -> i32 {
        let mut z = x ^ y;
        let mut ret = 0;
        while z != 0 {
            z = z & (z - 1);
            ret += 1;
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::hamming_distance(1, 4));
}
