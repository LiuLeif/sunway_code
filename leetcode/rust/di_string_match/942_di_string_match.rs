// 2018-12-26 14:50
// @lc id=942 lang=rust
struct Solution {}
impl Solution {
    pub fn di_string_match(s: String) -> Vec<i32> {
        let (mut lo, mut hi) = (0 as i32, s.len() as i32);
        let mut ret: Vec<_> = vec![];
        s.chars().for_each(|c| match c {
            'I' => {
                ret.push(lo);
                lo += 1;
            }
            'D' => {
                ret.push(hi);
                hi -= 1;
            }
            _ => (),
        });
        ret.push(lo);
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::di_string_match("IDID".to_owned()));
}
