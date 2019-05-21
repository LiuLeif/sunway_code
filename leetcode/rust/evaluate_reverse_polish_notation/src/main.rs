// Accepted
// @lc id=150 lang=rust
// problem: evaluate_reverse_polish_notation
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack = vec![];
        let mut ops = HashMap::<&str, &dyn Fn(i32, i32) -> i32>::new();
        ops.insert("+", &|a, b| a + b);
        ops.insert("-", &|a, b| a - b);
        ops.insert("*", &|a, b| a * b);
        ops.insert("/", &|a, b| a / b);
        for token in tokens {
            match token.as_str() {
                "+" | "-" | "*" | "/" => {
                    let (b, a) = (stack.pop().unwrap(), stack.pop().unwrap());
                    let op = ops.get(token.as_str()).unwrap();
                    stack.push(op(a, b));
                }
                _ => stack.push(token.parse().unwrap()),
            }
        }
        stack.pop().unwrap()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::eval_rpn(vec![
            "2".to_owned(),
            "1".to_owned(),
            "+".to_owned(),
            "3".to_owned(),
            "*".to_owned()
        ])
    );
}
