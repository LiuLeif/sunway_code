// Accepted
// 2019-01-06 01:39
// @lc id=739 lang=rust
// problem: daily_temperatures
struct Solution {}
impl Solution {
    pub fn daily_temperatures(t: Vec<i32>) -> Vec<i32> {
        let mut ret = vec![0; t.len()];
        let mut stack: Vec<usize> = vec![];
        for i in 0..t.len() {
            while !stack.is_empty() && t[*stack.last().unwrap()] < t[i] {
                let index = stack.pop().unwrap();
                ret[index] = (i - index) as i32;
            }
            stack.push(i);
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::daily_temperatures(vec![73, 74, 75, 71, 69, 72, 76, 73])
    );
}
