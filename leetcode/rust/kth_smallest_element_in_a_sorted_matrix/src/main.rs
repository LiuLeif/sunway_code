// Accepted
// @lc id=378 lang=rust
// problem: kth_smallest_element_in_a_sorted_matrix
use std::collections::binary_heap::*;

struct Solution {}
impl Solution {
    pub fn kth_smallest(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
        // tuple: (-val,x,y)
        let mut pq = BinaryHeap::new();
        let (m, n) = (matrix[0].len(), matrix[0].len());
        for i in 0..n {
            pq.push((-matrix[0][i], 0, i));
        }
        for _ in 0..k - 1 {
            let (_, x, y) = pq.pop().unwrap();
            if x < n - 1 {
                pq.push((-matrix[x + 1][y], x + 1, y));
            }
        }
        return -pq.pop().unwrap().0;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::kth_smallest(vec![vec![1, 5, 9], vec![10, 11, 13], vec![12, 13, 15]], 8)
    );
}
