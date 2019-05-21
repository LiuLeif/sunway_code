// Accepted
// @lc id=202 lang=rust
// problem: happy_number
struct Solution {}
impl Solution {
    fn square_sum(mut x: i32) -> i32 {
        let mut ret = 0;
        while x != 0 {
            ret += (x % 10).pow(2);
            x /= 10;
        }
        ret
    }

    pub fn is_happy(n: i32) -> bool {
        let mut slow = n;
        let mut fast = n;
        while {
            slow = Solution::square_sum(slow);
            fast = Solution::square_sum(Solution::square_sum(fast));
            slow != fast
        } {}
        slow == 1
    }
}

fn main() {
    println!("{:?}", Solution::is_happy(19));
}
