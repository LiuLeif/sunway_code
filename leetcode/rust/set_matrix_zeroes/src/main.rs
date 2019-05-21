// Accepted
// @lc id=73 lang=rust
// problem: set_matrix_zeroes
struct Solution {}
impl Solution {
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let (m, n) = (matrix.len(), matrix[0].len());
        let (mut x, mut y) = (-1, -1);
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    x = i as i32;
                    y = j as i32;
                    break;
                }
            }
        }
        if x == -1 {
            return;
        }
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == 0 {
                    matrix[x as usize][j] = 0;
                    matrix[i][y as usize] = 0;
                }
            }
        }
        for i in 0..m {
            if i == x as usize {
                continue;
            }
            if matrix[i][y as usize] == 0 {
                for j in 0..n {
                    matrix[i][j] = 0;
                }
            }
        }
        for j in 0..n {
            if j == y as usize {
                continue;
            }
            if matrix[x as usize][j] == 0 {
                for i in 0..m {
                    matrix[i][j] = 0;
                }
            }
        }
        for i in 0..m {
            matrix[i][y as usize] = 0;
        }
        for j in 0..n {
            matrix[x as usize][j] = 0;
        }
    }
}

fn main() {
    let mut v = vec![vec![0, 1, 2, 0], vec![3, 4, 5, 2], vec![1, 3, 1, 5]];
    Solution::set_zeroes(&mut v);
    println!("{:?}", v);
}
