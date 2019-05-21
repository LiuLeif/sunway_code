// Accepted
// @lc id=119 lang=rust
// problem: pascal_s_triangle_ii
struct Solution {}
impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let row_index = row_index as usize;
        let mut row_a = vec![0; row_index + 1];
        let mut row_b = vec![0; row_index + 1];
        row_a[0] = 1;
        for i in 1..=row_index {
            for j in 0..=i {
                let a = {
                    if j == 0 {
                        0
                    } else {
                        row_a[j - 1]
                    }
                };
                let b = {
                    if j == i {
                        0
                    } else {
                        row_a[j]
                    }
                };
                row_b[j] = a + b;
            }
            std::mem::swap(&mut row_a, &mut row_b);
        }
        row_a
    }
}

fn main() {
    println!("{:?}", Solution::get_row(3));
}
