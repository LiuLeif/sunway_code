// Accepted
// @lc id=492 lang=rust
// problem: construct_the_rectangle
struct Solution {}
impl Solution {
    pub fn construct_rectangle(area: i32) -> Vec<i32> {
        let max = (area as f64).sqrt() as i32;
        let x = (0..max + 1)
            .rev()
            .skip_while(|i| area % i != 0)
            .next()
            .unwrap_or(0);
        vec![area / x, x]
    }
}

fn main() {
    println!("{:?}", Solution::construct_rectangle(2));
    println!("{:?}", Solution::construct_rectangle(4));
}
