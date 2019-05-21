// Accepted
// @lc id=48 lang=rust
// problem: rotate_image
struct Solution {}
impl Solution {
    fn transpose(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for i in 0..n {
            for j in i..n {
                let tmp = matrix[i][j];
                matrix[i][j] = matrix[j][i];
                matrix[j][i] = tmp;
            }
        }
    }

    fn reverse(matrix: &mut Vec<Vec<i32>>) {
        let n = matrix.len();
        for j in 0..n / 2 {
            for i in 0..n {
                let tmp = matrix[i][j];
                matrix[i][j] = matrix[i][n - j - 1];
                matrix[i][n - j - 1] = tmp;
            }
        }
    }

    pub fn rotate(matrix: &mut Vec<Vec<i32>>) {
        Solution::transpose(matrix);
        Solution::reverse(matrix);
    }
}

fn main() {
    let mut v = vec![
        vec![5, 1, 9, 11],
        vec![2, 4, 8, 10],
        vec![13, 3, 6, 7],
        vec![15, 14, 12, 16],
    ];
    Solution::rotate(&mut v);
    println!("{:?}", v);
}
