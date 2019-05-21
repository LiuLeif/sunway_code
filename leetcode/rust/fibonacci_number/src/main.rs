// Accepted
// @lc id=509 lang=rust
// problem: fibonacci_number
struct Solution {}
struct Fibonacci {
    prev: Vec<i32>,
}
impl Fibonacci {
    fn new() -> Self {
        Fibonacci { prev: vec![] }
    }
}
impl Iterator for Fibonacci {
    type Item = i32;
    fn next(&mut self) -> (Option<i32>) {
        if self.prev.is_empty() {
            self.prev.insert(0, 0);
        } else if self.prev.len() == 1 {
            self.prev.insert(0, 1);
        } else {
            let tmp = self.prev[0] + self.prev[1];
            self.prev[1] = self.prev[0];
            self.prev[0] = tmp;
        }
        return Some(self.prev[0]);
    }
}

impl Solution {
    pub fn fib(n: i32) -> i32 {
        // 0,1,1,2,3
        Fibonacci::new().nth(n as usize).unwrap()
    }
}

fn main() {
    println!("{:?}", Solution::fib(4));
}
