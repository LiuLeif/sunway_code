// Accepted
// @lc id=263 lang=rust
// problem: ugly_number
struct Solution {}
impl Solution {
    pub fn is_ugly(num: i32) -> bool {
        // 6 = 2 * 3
        if num <= 0 {
            return false;
        }
        let mut num = num;
        for d in &[2, 3, 5] {
            while num != 1 {
                num /= d;
            }
        }
        num == 1
    }
}

fn main() {
    println!("{:?}", Solution::is_ugly(8));
}
