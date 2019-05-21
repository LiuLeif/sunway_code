// Accepted
// @lc id=565 lang=rust
// problem: array_nesting
struct Solution {}
impl Solution {
    fn dfs(curr: usize, nums: &[i32], visited: &mut [bool], length: i32) -> (i32) {
        if visited[curr] {
            return length;
        }
        visited[curr] = true;
        return Solution::dfs(nums[curr] as usize, nums, visited, length + 1);
    }

    pub fn array_nesting(nums: Vec<i32>) -> i32 {
        let mut visited = vec![false; nums.len()];
        let mut ret = 0;
        for i in 0..nums.len() {
            if visited[i] {
                continue;
            }
            ret = std::cmp::max(ret, Solution::dfs(i, &nums, &mut visited, 0));
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::array_nesting(vec![5, 4, 0, 3, 1, 6, 2]));
}
