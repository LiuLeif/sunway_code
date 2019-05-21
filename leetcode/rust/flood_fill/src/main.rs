// Accepted
// @lc id=733 lang=rust
// problem: flood_fill
struct Solution {}
impl Solution {
    fn dfs(image: &mut Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32, orig_color: i32) {
        if image[sr as usize][sc as usize] != orig_color {
            return;
        }
        image[sr as usize][sc as usize] = new_color;
        for (x, y) in vec![(1, 0), (-1, 0), (0, 1), (0, -1)] {
            let (sr, sc) = (sr + x, sc + y);
            if sr < 0 || sr >= image.len() as i32 || sc < 0 || sc >= image[0].len() as i32 {
                continue;
            }
            Solution::dfs(image, sr, sc, new_color, orig_color);
        }
    }

    pub fn flood_fill(image: Vec<Vec<i32>>, sr: i32, sc: i32, new_color: i32) -> Vec<Vec<i32>> {
        let mut image = image;
        let orig_color = image[sr as usize][sc as usize];
        if new_color == orig_color {
            return image;
        }
        Solution::dfs(&mut image, sr, sc, new_color, orig_color);
        return image;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::flood_fill(vec![vec![1, 1, 1], vec![1, 1, 0], vec![1, 0, 1],], 1, 1, 2)
    );
}
