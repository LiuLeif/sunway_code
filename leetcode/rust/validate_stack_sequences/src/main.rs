// Accepted
// @lc id=946 lang=rust
// problem: validate_stack_sequences
struct Solution {}
impl Solution {
    pub fn validate_stack_sequences(pushed: Vec<i32>, popped: Vec<i32>) -> bool {
        // 1 2 3 4 5
        // 4 5 3 2 1
        let mut popped = popped;
        popped.reverse();

        let mut stack = vec![];
        for i in pushed {
            stack.push(i);
            while !stack.is_empty()
                && !popped.is_empty()
                && stack.last().unwrap() == popped.last().unwrap()
            {
                stack.pop();
                popped.pop();
            }
        }
        stack.is_empty() && popped.is_empty()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::validate_stack_sequences(vec![1, 2, 3, 4, 5], vec![4, 5, 3, 2, 1])
    );
    println!(
        "{:?}",
        Solution::validate_stack_sequences(vec![1, 2, 3, 4, 5], vec![4, 3, 5, 1, 2])
    );
}
