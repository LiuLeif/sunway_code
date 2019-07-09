// Accepted
// @lc id=1089 lang=rust
// problem: duplicate_zeros
struct Solution {}
impl Solution {
    pub fn duplicate_zeros(arr: &mut Vec<i32>) {
        for i in (0..arr.len() - 1).rev() {
            if arr[i] == 0 {
                for j in (i + 1..arr.len() - 1).rev() {
                    arr[j + 1] = arr[j];
                }
                arr[i + 1] = 0;
            }
        }
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::duplicate_zeros(&mut vec![1, 0, 2, 3, 0, 4, 5, 0])
    );
}
