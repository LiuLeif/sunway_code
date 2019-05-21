// Accepted
// @lc id=769 lang=rust
// problem: max_chunks_to_make_sorted
struct Solution {}
impl Solution {
    pub fn max_chunks_to_sorted(arr: Vec<i32>) -> i32 {
        // 1,0,2,3,4
        let mut hi = 0;
        let mut ret = 0;
        for i in 0..arr.len() {
            hi = std::cmp::max(hi, arr[i]);
            if hi == i as i32 {
                ret += 1;
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::max_chunks_to_sorted(vec![1, 0, 2, 4, 3]));
}
