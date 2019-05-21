// Accepted
// @lc id=696 lang=rust
// problem: count_binary_substrings
struct Solution {}
impl Solution {
    pub fn count_binary_substrings(s: String) -> i32 {
        let mut curr = 0;
        let mut prev = 'x';
        let mut counter = vec![];
        for c in s.chars() {
            if prev == c {
                curr += 1;
            } else {
                counter.push(curr);
                curr = 1;
            }
            prev = c;
        }
        counter.push(curr);
        counter[1..]
            .windows(2)
            .map(|v| std::cmp::min(v[0], v[1]))
            .sum()
    }
    // pub fn count_binary_substrings(s: String) -> i32 {
    //     let (mut one, mut zero) = (0, 0);
    //     let mut ret = 0;
    //     let mut prev = 'x';
    //     for c in s.chars() {
    //         if c == '0' {
    //             if prev != '0' {
    //                 zero = 0;
    //             }
    //             prev = '0';
    //             zero += 1;
    //             if one != 0 {
    //                 ret += 1;
    //                 one -= 1;
    //             }
    //         } else {
    //             if prev != '1' {
    //                 one = 0;
    //             }
    //             prev = '1';
    //             one += 1;
    //             if zero != 0 {
    //                 ret += 1;
    //                 zero -= 1;
    //             }
    //         }
    //     }
    //     return ret;
    // }
}

fn main() {
    println!(
        "{:?}",
        Solution::count_binary_substrings("00110011".to_owned())
    );
    println!(
        "{:?}",
        Solution::count_binary_substrings("10101".to_owned())
    );
}
