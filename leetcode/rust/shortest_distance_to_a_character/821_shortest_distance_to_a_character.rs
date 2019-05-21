// 2018-12-28 17:55
// @lc id=821 lang=rust
struct Solution {}
impl Solution {
    pub fn shortest_to_char(s: String, c: char) -> Vec<i32> {
        let mut ret = vec![0; s.len()];
        let mut prev_e = -1;
        let s = s.as_bytes();
        for i in 0..s.len() {
            if s[i] == c as u8 {
                ret[i] = 0;
                for j in prev_e as i32 + 1..i as i32 {
                    if prev_e == -1 {
                        ret[j as usize] = i as i32 - j;
                    } else {
                        ret[j as usize] = std::cmp::min(i as i32 - j, j - prev_e);
                    }
                }
                prev_e = i as i32;
            }
        }
        if prev_e != s.len() as i32 - 1 {
            for i in prev_e + 1..s.len() as i32 {
                ret[i as usize] = i - prev_e
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::shortest_to_char("aaab".to_owned(), 'b'));
}
