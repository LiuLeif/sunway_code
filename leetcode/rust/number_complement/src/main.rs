// 2018-12-30 19:35
// @lc id=476 lang=rust
struct Solution {}
impl Solution {
    pub fn find_complement(num: i32) -> i32 {
        // 101
        let mut ret = 0;
        let mut power = 1;
        let mut num = num;
        while num != 0 {
            let n = num & 1;
            num = num >> 1;
            ret += (n + 1) % 2 * power;
            power = power << 1;
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::find_complement(5));
}
