// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1137 lang=rust
 *
 * [1137] N-th Tribonacci Number
 */

// @lc code=start
impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        // 0 1 1 2 4 7
        let mut value = vec![0; 38];
        value[1] = 1;
        value[2] = 1;
        for i in 3..38 {
            value[i] = value[i - 3] + value[i - 2] + value[i - 1];
        }
        value[n as usize]
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::tribonacci(4));
}
