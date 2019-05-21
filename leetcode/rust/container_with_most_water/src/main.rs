// Accepted
// @lc id=11 lang=rust
// problem: container_with_most_water
struct Solution {}
impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let (mut lo, mut hi) = (0, height.len() - 1);
        let mut ret = 0;
        while lo < hi {
            ret = std::cmp::max(
                ret,
                ((hi - lo) as i32) * std::cmp::min(height[lo], height[hi]),
            );
            if height[lo] <= height[hi] {
                lo += 1;
            } else {
                hi -= 1;
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::max_area(vec![1, 8, 6, 2, 5, 4, 8, 3, 7]));
}
