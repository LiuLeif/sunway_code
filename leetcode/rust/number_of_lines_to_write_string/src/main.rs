// 2019-01-03 13:48
// @lc id=806 lang=rust
// problem: number_of_lines_to_write_string
struct Solution {}
impl Solution {
    pub fn number_of_lines(widths: Vec<i32>, s: String) -> Vec<i32> {
        // 10 20 25 40 98 102
        let mut row = 1;
        let mut col = 0;
        for i in s.chars() {
            let index = i as usize - 'a' as usize;
            if col + widths[index] <= 100 {
                col += widths[index];
            } else {
                row += 1;
                col = widths[index];
            }
        }
        return vec![row, col];
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::number_of_lines(
            vec![
                10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
                10, 10, 10, 10, 10,
            ],
            "abcdefghijklmnopqrstuvwxyz".to_owned()
        )
    );
}
