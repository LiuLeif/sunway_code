// Accepted
// @lc id=29 lang=rust
// problem: divide_two_integers
struct Solution {}
impl Solution {
    fn mul(a: i64, b: i64) -> i64 {
        let mut ret = 0;
        for i in 0..a {
            ret += b;
        }
        ret
    }

    pub fn divide(dividend: i32, divisor: i32) -> i32 {
        if dividend == 0 {
            return 0;
        }

        let mut negative = false;
        if !(dividend < 0 && divisor < 0) && !(dividend > 0 && divisor > 0) {
            negative = true;
        }
        let dividend = (dividend as i64).abs();
        let divisor = (divisor as i64).abs();
        let (mut lo, mut hi) = (1, dividend);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            let mul = Solution::mul(divisor, mid);
            if mul > dividend {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        if negative {
            hi = -hi;
        }
        hi = std::cmp::min(std::i32::MAX as i64, hi);

        return hi as i32;
    }
}

fn main() {
    println!("{:?}", Solution::divide(-2147483648, 1));
}
