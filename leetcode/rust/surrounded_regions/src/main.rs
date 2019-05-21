// Accepted
// @lc id=130 lang=rust
// problem: surrounded_regions
struct Solution {}
impl Solution {
    fn dfs(board: &mut Vec<Vec<char>>, x: usize, y: usize) {
        if board[x][y] != 'O' {
            return;
        }
        let (m, n) = (board.len() as i32, board[0].len() as i32);
        board[x][y] = 'P';
        for i in vec![-1, 0, 1] {
            for j in vec![-1, 0, 1] {
                if (i != 0 && j != 0) || (i == 0 && j == 0) {
                    continue;
                }
                let (x, y) = (x as i32 + i, y as i32 + j);
                if x < 0 || x >= m || y < 0 || y >= n {
                    continue;
                }
                Solution::dfs(board, x as usize, y as usize);
            }
        }
    }

    pub fn solve(board: &mut Vec<Vec<char>>) {
        if board.is_empty() {
            return;
        }
        let (m, n) = (board.len(), board[0].len());
        for i in vec![0, m - 1] {
            for j in 0..n {
                Solution::dfs(board, i, j);
            }
        }
        for j in vec![0, n - 1] {
            for i in 0..m {
                Solution::dfs(board, i, j);
            }
        }
        for x in 0..m {
            for y in 0..n {
                if board[x][y] == 'O' {
                    board[x][y] = 'X'
                }
            }
        }
        for x in 0..m {
            for y in 0..n {
                if board[x][y] == 'P' {
                    board[x][y] = 'O'
                }
            }
        }
    }
}

fn main() {
    let mut v = vec![
        vec!['X', 'O', 'X', 'O', 'X', 'O'],
        vec!['O', 'X', 'O', 'X', 'O', 'X'],
        vec!['X', 'O', 'X', 'O', 'X', 'O'],
        vec!['O', 'X', 'O', 'X', 'O', 'X'],
    ];
    Solution::solve(&mut v);
    println!("{:?}", v);
}
