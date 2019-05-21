// Accepted
// @lc id=526 lang=rust
// problem: beautiful_arrangement
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut i32, index: i32, mask: i32, n: i32) -> () {
        if index == n {
            *ret += 1;
            return;
        }
        for i in 0..n {
            if mask & (1 << i) != 0 {
                continue;
            }
            if (index + 1) % (i + 1) != 0 && (i + 1) % (index + 1) != 0 {
                continue;
            }
            Solution::backtrack(ret, index + 1, mask | (1 << i), n);
        }
    }

    pub fn count_arrangement(n: i32) -> i32 {
        let mut ret = 0;
        Solution::backtrack(&mut ret, 0, 0, n);
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::count_arrangement(2));
}
