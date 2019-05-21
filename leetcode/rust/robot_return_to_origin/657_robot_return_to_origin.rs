// 2018-12-25 14:16
// @lc id=657 lang=rust
struct Solution {}
impl Solution {
    pub fn judge_circle(moves: String) -> bool {
        let (mut h, mut v) = moves.chars().fold((0, 0), |(mut h, mut v), d| {
            match d {
                'U' => v += 1,
                'D' => v -= 1,
                'L' => h += 1,
                'R' => h -= 1,
                _ => (),
            }
            (h, v)
        });
        return (h, v) == (0, 0);
    }
}

fn main() {
    println!("{:?}", Solution::judge_circle("UDLR".to_owned()));
}
