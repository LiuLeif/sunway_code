// Accepted
// @lc id=506 lang=rust
// problem: relative_ranks

struct Solution {}
impl Solution {
    fn partition(
        nums: &mut [(usize, i32)],
        ranks: &mut [String],
        left: usize,
        right: usize,
    ) -> i32 {
        let mut p = left;
        let mut left = left + 1;

        while left <= right {
            if nums[left].1 <= nums[p].1 {
                left += 1;
            } else {
                nums.swap(p + 1, left);
                ranks.swap(nums[p + 1].0, nums[left].0);
                nums.swap(p, p + 1);
                ranks.swap(nums[p].0, nums[p + 1].0);
                p += 1;
                left += 1
            }
        }
        p as i32
    }

    fn quick_sort(nums: &mut [(usize, i32)], ranks: &mut [String], left: i32, right: i32) {
        if left >= right {
            return;
        }
        let p = Solution::partition(nums, ranks, left as usize, right as usize);
        Solution::quick_sort(nums, ranks, left, p - 1);
        Solution::quick_sort(nums, ranks, p + 1, right);
    }

    pub fn find_relative_ranks(nums: Vec<i32>) -> Vec<String> {
        // 1 2 3 4 5
        // 4 3 2 1 0
        let N = nums.len();
        let mut ranks = (0..N)
            .map(|i| match i {
                0 => "Gold Medal".to_owned(),
                1 => "Silver Medal".to_owned(),
                2 => "Bronze Medal".to_owned(),
                _ => (i + 1).to_string(),
            })
            .collect::<Vec<String>>();
        let mut nums = nums.into_iter().enumerate().collect::<Vec<(usize, i32)>>();
        Solution::quick_sort(&mut nums, &mut ranks, 0, N as i32 - 1);
        ranks
    }
}

fn main() {
    println!("{:?}", Solution::find_relative_ranks(vec![2, 3]));
}
