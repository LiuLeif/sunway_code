// Accepted
// @lc id=79 lang=rust
// problem: word_search
struct Solution {}
impl Solution {
    fn backtrack(
        visited: &mut Vec<Vec<bool>>,
        board: &Vec<Vec<char>>,
        x: usize,
        y: usize,
        word: &[u8],
    ) -> bool {
        if visited[x][y] {
            return false;
        }
        if board[x][y] != word[0] as char {
            return false;
        }
        if word.len() == 1 {
            return true;
        }
        visited[x][y] = true;
        let (m, n) = (board.len(), board[0].len());
        for delta_x in vec![-1, 0, 1] {
            for delta_y in vec![-1, 0, 1] {
                if (delta_x == 0 && delta_y == 0) || (delta_x != 0 && delta_y != 0) {
                    continue;
                }
                let (newx, newy) = (x as i32 + delta_x, y as i32 + delta_y);
                if newx < 0 || newx >= m as i32 || newy < 0 || newy >= n as i32 {
                    continue;
                }
                if Solution::backtrack(visited, board, newx as usize, newy as usize, &word[1..]) {
                    return true;
                }
            }
        }
        visited[x][y] = false;
        false
    }

    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        let (m, n) = (board.len(), board[0].len());
        for i in 0..m {
            for j in 0..n {
                if Solution::backtrack(&mut vec![vec![false; n]; m], &board, i, j, &word.as_bytes())
                {
                    return true;
                }
            }
        }
        false
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::exist(
            vec![
                vec!['A', 'B', 'C', 'E'],
                vec!['S', 'F', 'C', 'S'],
                vec!['A', 'D', 'E', 'E'],
            ],
            "AB".to_owned()
        ),
    );
}
