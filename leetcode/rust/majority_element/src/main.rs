// Accepted
// @lc id=169 lang=rust
// problem: majority_element
struct Solution {}
impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let mut x = 0;
        let mut count = 0;
        for i in nums {
            if count == 0 {
                x = i;
            }
            if i == x {
                count += 1
            } else {
                count -= 1;
            }
        }
        return x;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::majority_element(vec![2, 2, 1, 1, 1, 2, 2])
    );
    println!("{:?}", Solution::majority_element(vec![2, 3, 3]));
}
