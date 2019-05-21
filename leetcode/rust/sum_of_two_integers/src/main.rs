// Accepted
// @lc id=371 lang=rust
// problem: sum_of_two_integers
struct Solution {}
impl Solution {
    pub fn get_sum(a: i32, b: i32) -> i32 {
        let mut ret = 0;
        let mut carry = 0;
        for bit in 0..32 {
            let bit_a = a >> bit & 1;
            let bit_b = b >> bit & 1;

            let bit_x = bit_a ^ bit_b ^ carry;
            if (carry & bit_a == 1) || (carry & bit_b) == 1 || (bit_a & bit_b) == 1 {
                carry = 1;
            } else {
                carry = 0;
            }
            ret = ret | (bit_x << bit);
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::get_sum(4, 5));
}
