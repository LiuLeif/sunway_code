// Accepted
// @lc id=1018 lang=rust
// problem: binary_prefix_divisible_by_5
struct Solution {}
impl Solution {
    pub fn prefixes_div_by5(a: Vec<i32>) -> Vec<bool> {
        let mut ret = vec![];
        let mut curr = 0;
        for bit in a {
            curr = (curr * 2 + bit) % 5;
            ret.push(curr == 0);
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::prefixes_div_by5(vec![
            1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0,
            0, 0, 1, 1, 0, 1, 0, 0, 0, 1
        ])
    );
}
