// Accepted
// @lc id=216 lang=rust
// problem: combination_sum_iii
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<i32>>,
        curr: &mut Vec<i32>,
        k: i32,
        sum: i32,
        start: i32,
        n: i32,
    ) {
        if curr.len() == k as usize {
            if sum == n {
                ret.push(curr.clone());
            }
            return;
        }
        if sum + start > n {
            return;
        }
        for i in start..10 {
            curr.push(i);
            Solution::backtrack(ret, curr, k, sum + i, i + 1, n);
            curr.pop();
        }
    }

    pub fn combination_sum3(k: i32, n: i32) -> Vec<Vec<i32>> {
        // 1,2,3,4,5,6,7,8,9
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], k, 0, 1, n);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::combination_sum3(3, 9));
}
