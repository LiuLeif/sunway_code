// Accepted
// @lc id=926 lang=rust
// problem: flip_string_to_monotone_increasing
use std::collections::HashMap;
struct Solution {}
impl Solution {
    fn dfs(cache: &mut HashMap<(usize, u8), i32>, s: &[u8], index: usize, prev: u8) -> i32 {
        if index == s.len() {
            return 0;
        }
        if cache.contains_key(&(index, prev)) {
            return cache[&(index, prev)];
        }
        let mut ret = std::i32::MAX;
        let curr = s[index];
        if curr >= prev {
            ret = std::cmp::min(ret, Solution::dfs(cache, s, index + 1, curr));
        }
        let curr = {
            if curr == b'1' {
                b'0'
            } else {
                b'1'
            }
        };
        if curr >= prev {
            ret = std::cmp::min(ret, Solution::dfs(cache, s, index + 1, curr) + 1);
        }
        cache.insert((index, prev), ret);
        return ret;
    }

    pub fn min_flips_mono_incr(s: String) -> i32 {
        let mut cache = HashMap::new();
        std::cmp::min(
            Solution::dfs(&mut cache, &s.as_bytes(), 0, 0),
            Solution::dfs(&mut cache, &s.as_bytes(), 0, 1),
        )
    }
}

fn main() {
    println!("{:?}", Solution::min_flips_mono_incr("00110".to_owned()));
    println!("{:?}", Solution::min_flips_mono_incr("00011000".to_owned()));
    println!("{:?}", Solution::min_flips_mono_incr("111".to_owned()));
}
