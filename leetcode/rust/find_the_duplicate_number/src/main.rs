// Accepted
// @lc id=287 lang=rust
// problem: find_the_duplicate_number
struct Solution {}
impl Solution {
    fn count(nums: &[i32], target: i32) -> (i32, i32) {
        nums.iter().fold((0, 0), |(a, b), &i| {
            let mut smaller = a;
            let mut larger = b;
            if i < target {
                smaller += 1;
            }
            if i > target {
                larger += 1;
            }
            (smaller, larger)
        })
    }

    pub fn find_duplicate(nums: Vec<i32>) -> i32 {
        let n = nums.len() - 1;
        let (mut lo, mut hi) = (1, n);

        while lo <= hi {
            let mid = (lo + hi) / 2;
            let (smaller, larger) = Solution::count(&nums, mid as i32);
            if smaller + larger < n as i32 {
                return mid as i32;
            }
            if smaller >= mid as i32 {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        unreachable!();
    }
}

fn main() {
    println!("{:?}", Solution::find_duplicate(vec![3, 1, 2, 1]));
    println!("{:?}", Solution::find_duplicate(vec![1, 2, 3, 3]));
}
