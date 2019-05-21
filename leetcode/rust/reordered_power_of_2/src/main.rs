// Accepted
// @lc id=869 lang=rust
// problem: reordered_power_of_2
struct Solution {}
impl Solution {
    fn parse(mut x: i32) -> [i32; 10] {
        let mut ret = [0; 10];
        while x != 0 {
            let index = x % 10;
            ret[index as usize] += 1;
            x /= 10;
        }
        return ret;
    }

    pub fn reordered_power_of2(n: i32) -> bool {
        let n = Solution::parse(n);
        (0..31).map(|i| Solution::parse(1 << i)).any(|x| x == n)
    }
}

fn main() {
    println!("{:?}", Solution::reordered_power_of2(16));
}
