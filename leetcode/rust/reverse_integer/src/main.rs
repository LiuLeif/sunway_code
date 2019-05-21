// Accepted
// @lc id=7 lang=rust
// problem: reverse_integer
struct Solution {}
impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut ret = 0i64;
        let mut x = x;
        let mut sign = 1;
        if x < 0 {
            x = -x;
            sign = -1;
        }
        while x != 0 {
            ret = ret * 10 + x as i64 % 10;
            x /= 10;
        }
        ret *= sign;
        if ret > std::i32::MAX as i64 || ret < std::i32::MIN as i64 {
            return 0;
        }
        ret as i32
    }
}

fn main() {
    println!("{:?}", Solution::reverse(-128));
}
