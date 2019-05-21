// Accepted
// @lc id=171 lang=rust
// problem: excel_sheet_column_number

struct Solution {}
impl Solution {
    pub fn title_to_number(s: String) -> i32 {
        // A-1
        // Z-26
        // ZY
        // 26*26+25
        s.chars()
            .rev()
            .map(|c| c as i32 - 'A' as i32 + 1)
            .fold((0, 1), |(sum, step), n| (sum + n * step, step * 26))
            .0
    }
}

fn main() {
    println!("{:?}", Solution::title_to_number("AB".to_owned()));
}
