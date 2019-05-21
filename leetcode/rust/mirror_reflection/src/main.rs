// Accepted
// @lc id=858 lang=rust
// problem: mirror_reflection
struct Solution {}
impl Solution {
    pub fn mirror_reflection(p: i32, q: i32) -> i32 {
        let mut q_total = 0;
        let mut ret = 2;
        loop {
            q_total += q;
            q_total %= 2 * p;
            if q_total == 0 {
                return 0;
            }
            ret = (!ret) & 3;
            if q_total == p {
                return ret;
            }
        }
    }
}

fn main() {
    println!("{:?}", Solution::mirror_reflection(2, 1));
    println!("{:?}", Solution::mirror_reflection(2, 2));
}
