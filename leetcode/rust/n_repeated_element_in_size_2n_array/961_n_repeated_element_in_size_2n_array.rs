// 2018-12-24 15:54
// @lc id=961 lang=rust
struct Solution {}
impl Solution {
    pub fn repeated_n_times(a: Vec<i32>) -> i32 {
        for i in 0..a.len() {
            let prev_one = {
                let mut index = i as i32 - 1;
                if index < 0 {
                    index += a.len() as i32
                }
                a[index as usize]
            };
            let prev_two = {
                let mut index = i as i32 - 2;
                if index < 0 {
                    index += a.len() as i32
                }
                a[index as usize]
            };
            if a[i] == prev_one || a[i] == prev_two {
                return a[i];
            }
        }
        return 0;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::repeated_n_times(vec![5, 1, 5, 2, 5, 3, 5, 4])
    );
}
