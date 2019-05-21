// Accepted
// @lc id=419 lang=rust
// problem: battleships_in_a_board
struct Solution {}
impl Solution {
    pub fn count_battleships(board: Vec<Vec<char>>) -> i32 {
        let (m, n) = (board.len(), board[0].len());
        let mut ret = 0;
        for i in 0..m {
            for j in 0..n {
                if board[i][j] == '.' {
                    continue;
                }
                ret += 2;
                for (x, y) in vec![(-1, 0), (1, 0), (0, 1), (0, -1)] {
                    let (i, j) = (i as i32 + x, j as i32 + y);
                    if i < 0
                        || i >= m as i32
                        || j < 0
                        || j >= n as i32
                        || board[i as usize][j as usize] == '.'
                    {
                        continue;
                    }
                    ret -= 1;
                }
            }
        }
        return ret / 2;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::count_battleships(vec![
            "X..X".chars().collect::<Vec<_>>(),
            "....".chars().collect::<Vec<_>>(),
            "X..X".chars().collect::<Vec<_>>(),
        ])
    );
}
