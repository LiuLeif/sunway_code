// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1122 lang=rust
 *
 * [1122] Relative Sort Array
 */

// @lc code=start
use std::collections::HashMap;
impl Solution {
    pub fn relative_sort_array(arr1: Vec<i32>, arr2: Vec<i32>) -> Vec<i32> {
        let mut ret = arr1;
        let mut ord = HashMap::<i32, i32>::new();
        let n = arr2.len() as i32;

        arr2.into_iter().enumerate().for_each(|(i, v)| {
            ord.insert(v, i as i32);
        });
        ret.sort_by(|a, b| {
            let (x, y) = (ord.get(&a).unwrap_or(&n), ord.get(&b).unwrap_or(&n));
            if (x == y) {
                a.cmp(b)
            } else {
                x.cmp(y)
            }
        });
        ret
    }
}
// @lc code=end

fn main() {
    println!(
        "{:?}",
        Solution::relative_sort_array(
            vec![2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19],
            vec![2, 1, 4, 3, 9, 6]
        )
    );
}
