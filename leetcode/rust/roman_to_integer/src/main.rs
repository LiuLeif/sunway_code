// Accepted
// @lc id=13 lang=rust
// problem: roman_to_integer
use std::collections::HashMap;
macro_rules! hashmap {
    ($( $key: expr => $val: expr ),*) => {{
         let mut map = ::std::collections::HashMap::new();
         $( map.insert($key, $val); )*
         map
    }}
}
struct Solution {}
impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let m = hashmap![
            b'I' => 1,
            b'V' => 5,
            b'X' => 10,
            b'L' => 50,
            b'C' => 100,
            b'D' => 500,
            b'M' => 1000
        ];
        let mut ret = 0;
        let mut max_val = 0;
        for c in s.bytes().rev() {
            let x = m[&c];
            if x < max_val {
                ret -= x;
            } else {
                ret += x;
                max_val = x;
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::roman_to_int("MCMXCIV".to_owned()));
}
