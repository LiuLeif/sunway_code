// 2019-01-03 15:34
// @lc id=496 lang=rust
// problem: next_greater_element_i
use std::collections::HashMap;
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    pub fn next_greater_element(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        let n = nums2.len();
        let mut index = HashMap::new();
        let mut next_larger = vec![-1; n];
        let mut stack = VecDeque::new();
        for i in 0..n {
            index.insert(nums2[i], i);
            while !stack.is_empty() && nums2[*stack.back().unwrap()] < nums2[i] {
                let index: usize = stack.pop_back().unwrap();
                next_larger[index] = nums2[i];
            }
            stack.push_back(i)
        }
        nums1
            .into_iter()
            .map(|i| next_larger[*index.get(&i).unwrap()])
            .collect()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::next_greater_element(vec![4, 1, 2], vec![1, 3, 4, 2],)
    );
    println!(
        "{:?}",
        Solution::next_greater_element(vec![2, 4], vec![1, 2, 3, 4],)
    );
}
