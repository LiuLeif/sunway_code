// 2019-01-03 16:59
// @lc id=766 lang=rust
// problem: toeplitz_matrix
struct Solution {}
impl Solution {
    pub fn is_toeplitz_matrix(matrix: Vec<Vec<i32>>) -> bool {
        let (m, n) = (matrix.len() as i32, matrix[0].len() as i32);
        for j in -m + 1..n {
            let mut target = None;
            for i in 0..m {
                let (x, y) = (i, i + j);
                if x < 0 || x >= m || y < 0 || y >= n {
                    continue;
                }
                match target {
                    Some(v) => {
                        if v != matrix[x as usize][y as usize] {
                            return false;
                        }
                    }
                    None => target = Some(matrix[x as usize][y as usize]),
                }
            }
        }
        return true;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_toeplitz_matrix(vec![vec![1, 2, 3, 4], vec![5, 1, 2, 3], vec![9, 5, 1, 2],])
    );
    println!(
        "{:?}",
        Solution::is_toeplitz_matrix(vec![vec![1, 2], vec![2, 2]])
    );
}
