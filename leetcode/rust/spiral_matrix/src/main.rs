// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=54 lang=rust
 *
 * [54] Spiral Matrix
 */

// @lc code=start
impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        if matrix.is_empty() {
            return vec![];
        }

        let m = matrix.len();
        let n = matrix[0].len();
        let mut ret = vec![];

        let (mut left, mut right, mut top, mut bottom) =
            (0 as i32, n as i32 - 1, 0 as i32, m as i32 - 1);
        loop {
            for j in left..=right {
                ret.push(matrix[top as usize][j as usize]);
            }
            top += 1;
            if left > right || top > bottom {
                break;
            }
            for i in top..=bottom {
                ret.push(matrix[i as usize][right as usize]);
            }
            right -= 1;
            if left > right || top > bottom {
                break;
            }
            for j in (left..=right).rev() {
                ret.push(matrix[bottom as usize][j as usize]);
            }
            bottom -= 1;
            if left > right || top > bottom {
                break;
            }
            for i in (top..=bottom).rev() {
                ret.push(matrix[i as usize][left as usize]);
            }
            left += 1;
            if left > right || top > bottom {
                break;
            }
        }
        ret
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::spiral_order(vec![vec![1], vec![2]]));
}
