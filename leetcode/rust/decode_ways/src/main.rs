// Accepted
// @lc id=91 lang=rust
// problem: decode_ways
struct Solution {}
impl Solution {
    fn backtrack(s: &str) -> i32 {
        if s.is_empty() {
            return 1;
        }
        let mut ret = 0;
        let a = *&s[0..1].parse::<i32>().unwrap();
        if a == 0 {
            return ret;
        }
        ret += Solution::backtrack(&s[1..]);
        if s.len() < 2 {
            return ret;
        }
        let a = *&s[0..2].parse::<i32>().unwrap();
        if a > 26 {
            return ret;
        }
        ret += Solution::backtrack(&s[2..]);
        ret
    }

    // A-1 Z->26
    pub fn num_decodings(s: String) -> i32 {
        Solution::backtrack(&s)
    }
}

fn main() {
    println!("{:?}", Solution::num_decodings("234".to_owned()));
}
