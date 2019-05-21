// Accepted
// @lc id=566 lang=rust
// problem: reshape_the_matrix
struct Solution {}
impl Solution {
    pub fn matrix_reshape(nums: Vec<Vec<i32>>, r: i32, c: i32) -> Vec<Vec<i32>> {
        if nums.is_empty() {
            return nums;
        }
        let (m, n) = (nums.len() as i32, nums[0].len() as i32);
        if m * n != r * c {
            return nums;
        }
        let mut ret = vec![vec![0; c as usize]; r as usize];
        for i in 0..r {
            for j in 0..c {
                let ii = (i * c + j) / n;
                let jj = (i * c + j) % n;
                ret[i as usize][j as usize] = nums[ii as usize][jj as usize];
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::matrix_reshape(vec![vec![1, 2], vec![3, 4]], 2, 4)
    );
}
