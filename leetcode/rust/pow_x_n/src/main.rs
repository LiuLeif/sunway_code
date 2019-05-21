// Accepted
// @lc id=50 lang=rust
// problem: pow_x_n
struct Solution {}
impl Solution {
    fn pow(x: f64, n: i32) -> f64 {
        // 2,10
        if n == 0 {
            return 1.0;
        }
        if n == 1 {
            return x;
        }
        let tmp = Solution::pow(x, n / 2);
        if n % 2 == 0 {
            tmp * tmp
        } else {
            tmp * tmp * x
        }
    }

    pub fn my_pow(x: f64, n: i32) -> f64 {
        if n > 0 {
            Solution::pow(x, n)
        } else {
            1.0 / Solution::pow(x, -n)
        }
    }
}

fn main() {
    println!("{:?}", Solution::my_pow(2.0, 100));
}
