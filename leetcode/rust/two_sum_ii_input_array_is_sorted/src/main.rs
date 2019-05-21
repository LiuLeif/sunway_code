// Accepted
// @lc id=167 lang=rust
// problem: two_sum_ii_input_array_is_sorted
struct Solution {}
impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let (mut lo, mut hi) = (0, numbers.len() - 1);
        while lo < hi {
            if numbers[lo] + numbers[hi] == target {
                return vec![lo as i32 + 1, hi as i32 + 1];
            }
            if numbers[lo] + numbers[hi] > target {
                hi -= 1;
            } else {
                lo += 1;
            }
        }
        return vec![];
    }
}

fn main() {
    println!("{:?}", Solution::two_sum(vec![2, 7, 11, 15], 9));
}
