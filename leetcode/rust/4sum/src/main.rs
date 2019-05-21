// Accepted
// @lc id=18 lang=rust
// problem: 4sum
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<i32>>,
        curr: &mut Vec<i32>,
        curr_sum: i32,
        nums: &[i32],
        index: usize,
        target: i32,
    ) -> () {
        if index >= nums.len() {
            return;
        }
        if curr.len() == 2 {
            // two sum: [index..]
            let (mut lo, mut hi) = (index, nums.len() - 1);
            while lo < hi {
                if nums[lo] + nums[hi] + curr_sum == target {
                    let mut tmp = curr.clone();
                    tmp.push(nums[lo]);
                    tmp.push(nums[hi]);
                    ret.push(tmp);
                    lo += 1;
                    while lo < hi && nums[lo] == nums[lo - 1] {
                        lo += 1;
                    }
                } else if nums[lo] + nums[hi] + curr_sum > target {
                    hi -= 1;
                } else {
                    lo += 1;
                }
            }
            return;
        }

        for i in index..nums.len() {
            if curr_sum + nums[i] > target && nums[i] > 0 {
                continue;
            }
            if i > index && nums[i] == nums[i - 1] {
                continue;
            }
            curr.push(nums[i]);
            Solution::backtrack(ret, curr, curr_sum + nums[i], nums, i + 1, target);
            curr.pop();
        }
    }

    pub fn four_sum(nums: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut nums = nums;
        nums.sort();
        Solution::backtrack(&mut ret, &mut vec![], 0, &nums, 0, target);
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::four_sum(vec![-3, -2, -1, 0, 0, 1, 2, 3], 0)
    );
}
