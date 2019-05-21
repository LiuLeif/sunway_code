// Accepted
// @lc id=74 lang=rust
// problem: search_a_2d_matrix
struct Solution {}
impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        if matrix.is_empty() {
            return false;
        }
        let (m, n) = (matrix.len() as i32, matrix[0].len() as i32);
        let (mut lo, mut hi) = (0 as i32, (m * n) as i32 - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            let (r, c) = ((mid / n) as usize, (mid % n) as usize);
            if matrix[r][c] == target {
                return true;
            }
            if matrix[r][c] > target {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return false;
    }
}

fn main() {
    // println!(
    //     "{:?}",
    //     Solution::search_matrix(
    //         vec![vec![1, 3, 5, 7], vec![10, 11, 16, 20], vec![23, 30, 34, 50],],
    //         30
    //     )
    // );
    println!("{:?}", Solution::search_matrix(vec![vec![]], 30));
}
