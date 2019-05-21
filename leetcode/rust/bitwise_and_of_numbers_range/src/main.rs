// Accepted
// @lc id=201 lang=rust
// problem: bitwise_and_of_numbers_range
struct Solution {}
impl Solution {
    pub fn range_bitwise_and(m: i32, n: i32) -> i32 {
        // 6 7 8 9 10 11 12
        //  110
        //  111
        // 1000
        // 1001
        // 1010
        // 1011
        // 1100
        let mut n = n;
        while m < n {
            n = n & (n - 1);
        }
        n
    }
}

fn main() {
    println!("{:?}", Solution::range_bitwise_and(1, 2));
}
