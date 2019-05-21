// Accepted
// @lc id=69 lang=rust
// problem: sqrt_x
struct Solution {}
impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        let (mut lo, mut hi) = (1, 46340);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if mid * mid > x {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        lo - 1
    }
}

fn main() {
    println!("{:?}", Solution::my_sqrt(2147395599));
}
