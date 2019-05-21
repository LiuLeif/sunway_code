// Accepted
// @lc id=70 lang=rust
// problem: climbing_stairs
struct Solution {}
impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let (mut x, mut y) = (1, 1);
        for _ in 0..n - 1 {
            let tmp = x;
            x = y;
            y += tmp;
        }
        y
    }
}

fn main() {
    println!("{:?}", Solution::climb_stairs(3));
}
