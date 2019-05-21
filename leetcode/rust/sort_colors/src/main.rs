// Accepted
// @lc id=75 lang=rust
// problem: sort_colors
struct Solution {}
impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let (mut zero, mut one, mut two) = (None, None, None);
        for i in 0..nums.len() {
            match nums[i] {
                0 => {
                    let mut index = i;
                    if let Some(two_index) = two {
                        nums.swap(two_index, index);
                        index = two_index;
                        two = Some(two_index + 1);
                    }
                    if let Some(one_index) = one {
                        nums.swap(one_index, index);
                        index = one_index;
                        one = Some(one_index + 1);
                    }
                    if zero.is_none() {
                        zero = Some(index);
                    }
                }
                1 => {
                    let mut index = i;
                    if let Some(two_index) = two {
                        nums.swap(two_index, index);
                        index = two_index;
                        two = Some(two_index + 1);
                    }
                    if one.is_none() {
                        one = Some(index);
                    }
                }
                2 => {
                    let index = i;
                    if two.is_none() {
                        two = Some(index);
                    }
                }
                _ => (),
            }
        }
    }
}

fn main() {
    let mut v = vec![2, 0, 2, 1, 1, 0, 2];
    Solution::sort_colors(&mut v);
    println!("{:?}", v);
}
