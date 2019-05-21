// Accepted
// @lc id=260 lang=rust
// problem: single_number_iii
struct Solution {}
impl Solution {
    pub fn single_number(nums: Vec<i32>) -> Vec<i32> {
        // 1,2,1,3,2,5
        // 011
        // 101
        // 110 = 6
        let tmp = nums.iter().fold(0, |x, &i| x ^ i);
        let tmp = tmp & (-tmp);
        let mut ret = vec![0, 0];
        for i in nums {
            if i & tmp != 0 {
                ret[0] ^= i;
            } else {
                ret[1] ^= i;
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::single_number(vec![1, 2, 1, 3, 2, 5]));
}
