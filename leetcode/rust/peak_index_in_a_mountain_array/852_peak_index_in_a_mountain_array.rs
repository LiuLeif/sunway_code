// 2018-12-28 14:17
// @lc id=852 lang=rust
struct Solution {}
impl Solution {
    pub fn peak_index_in_mountain_array(a: Vec<i32>) -> i32 {
        let (mut lo, mut hi) = (0, a.len() - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if a[mid + 1] > a[mid] {
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return lo as i32;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::peak_index_in_mountain_array(vec![0, 2, 1, 0])
    );
    println!(
        "{:?}",
        Solution::peak_index_in_mountain_array(vec![0, 3, 4, 5, 0])
    );
}
