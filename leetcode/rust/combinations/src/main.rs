// Accepted
// @lc id=77 lang=rust
// problem: combinations
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut Vec<Vec<i32>>, curr: &mut Vec<i32>, start: usize, n: usize, k: i32) {
        if k == 0 {
            ret.push(curr.to_vec());
            return;
        }
        for i in start..n {
            curr.push(i as i32 + 1);
            Solution::backtrack(ret, curr, i + 1, n, k - 1);
            curr.pop();
        }
    }

    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        // 1,2,3,4
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], 0, n as usize, k);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::combine(4, 3));
}
