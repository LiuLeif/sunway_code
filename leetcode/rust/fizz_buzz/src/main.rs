// 2019-01-03 14:28
// @lc id=412 lang=rust
// problem: fizz_buzz
struct Solution {}
impl Solution {
    pub fn fizz_buzz(n: i32) -> Vec<String> {
        (1..n + 1)
            .map(|i| {
                let mut ret = "".to_owned();
                if i % 3 == 0 {
                    ret.push_str("Fizz");
                }
                if i % 5 == 0 {
                    ret.push_str("Buzz");
                }
                if ret == "" {
                    ret = i.to_string();
                }
                ret
            })
            .collect()
    }
}

fn main() {
    println!("{:?}", Solution::fizz_buzz(15));
}
