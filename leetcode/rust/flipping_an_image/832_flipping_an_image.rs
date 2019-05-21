// 2018-12-25 13:56
// @lc id=832 lang=rust
struct Solution {}
impl Solution {
    pub fn flip_and_invert_image(a: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut a = a;
        a.iter_mut().for_each(|v| {
            let l = v.len();
            for i in 0..(l + 1) / 2 {
                if v[i] == v[l - i - 1] {
                    v[i] = (v[i] + 1) % 2;
                    v[l - i - 1] = v[i];
                }
            }
        });
        return a;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::flip_and_invert_image(vec![
            vec![1, 1, 0, 0],
            vec![1, 0, 0, 1],
            vec![0, 1, 1, 1],
            vec![1, 0, 1, 0],
        ])
    );

    println!(
        "{:?}",
        Solution::flip_and_invert_image(vec![vec![1, 1, 0], vec![1, 0, 1], vec![0, 0, 0],])
    );
}
