// Accepted
// 2019-01-05 20:09
// @lc id=762 lang=rust
// problem: prime_number_of_set_bits_in_binary_representation
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn count_prime_set_bits(l: i32, r: i32) -> i32 {
        // 10: 1010
        // 11: 1011
        // 12: 1100
        // ...
        // 15: 1111
        let primes = vec![2, 3, 5, 7, 11, 13, 17, 19, 23]
            .into_iter()
            .collect::<HashSet<_>>();
        
        (l..r + 1)
            .map(|mut i| {
                let mut count = 0;
                while i != 0 {
                    count += 1;
                    i = i & (i - 1);
                }
                count
            })
            .filter(|i| primes.contains(i))
            .count() as i32
    }
}

fn main() {
    println!("{:?}", Solution::count_prime_set_bits(6, 10));
}
