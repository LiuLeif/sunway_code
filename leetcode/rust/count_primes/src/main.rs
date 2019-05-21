// Accepted
// @lc id=204 lang=rust
// problem: count_primes
struct Solution {}
impl Solution {
    pub fn count_primes(n: i32) -> i32 {
        let mut is_prime = vec![true; n as usize];
        if let Some(element) = is_prime.get_mut(0) {
            *element = false;
        }
        if let Some(element) = is_prime.get_mut(1) {
            *element = false;
        }
        for i in 2..((n as f64).sqrt() as usize) + 1 {
            if !is_prime[i] {
                continue;
            }
            for j in (i..n as usize).take_while(|x| x * i < n as usize) {
                is_prime[i * j] = false;
            }
        }
        is_prime.into_iter().filter(|&x| x == true).count() as i32
    }
}

fn main() {
    println!("{:?}", Solution::count_primes(100));
}
