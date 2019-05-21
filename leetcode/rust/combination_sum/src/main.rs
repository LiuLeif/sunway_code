// Accepted
// @lc id=39 lang=rust
// problem: combination_sum
struct Solution {}
impl Solution {
    fn backtrack(
        candidate: &[i32],
        ret: &mut Vec<Vec<i32>>,
        current: &mut Vec<i32>,
        current_sum: i32,
        index: usize,
        target: i32,
    ) {
        if current_sum == target {
            ret.push(current.clone());
            return;
        }
        if index >= candidate.len() {
            return;
        }
        for i in index..candidate.len() {
            if current_sum + candidate[i] <= target {
                current.push(candidate[i]);
                Solution::backtrack(
                    candidate,
                    ret,
                    current,
                    current_sum + candidate[i],
                    i,
                    target,
                );
                current.pop();
            }
        }
    }

    pub fn combination_sum(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        Solution::backtrack(&candidates, &mut ret, &mut vec![], 0, 0, target);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::combination_sum(vec![2, 3, 5], 8));
}
