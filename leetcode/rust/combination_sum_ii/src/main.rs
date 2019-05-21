// Accepted
// @lc id=40 lang=rust
// problem: combination_sum_ii
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<i32>>,
        curr: &mut Vec<i32>,
        index: usize,
        target: i32,
        candidates: &[i32],
    ) {
        if target == 0 {
            ret.push(curr.clone());
            return;
        }
        if index == candidates.len() {
            return;
        }
        for i in index..candidates.len() {
            if candidates[i] > target {
                break;
            }
            if i != index && candidates[i] == candidates[i - 1] {
                continue;
            }
            curr.push(candidates[i]);
            Solution::backtrack(ret, curr, i + 1, target - candidates[i], candidates);
            curr.pop();
        }
    }

    pub fn combination_sum2(candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut candidates = candidates;
        candidates.sort();
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], 0, target, &candidates);
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::combination_sum2(vec![10, 1, 2, 7, 6, 1, 5], 8)
    );
}
