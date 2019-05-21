// Accepted
// @lc id=883 lang=rust
// problem: projection_area_of_3d_shapes
struct Solution {}
impl Solution {
    pub fn projection_area(grid: Vec<Vec<i32>>) -> i32 {
        let N = grid.len();
        let mut row_max = vec![0; N];
        let mut col_max = vec![0; N];
        let mut ret = 0;
        for i in 0..N {
            for j in 0..N {
                if grid[i][j] != 0 {
                    ret += 1;
                }
                row_max[i] = std::cmp::max(row_max[i], grid[i][j]);
                col_max[j] = std::cmp::max(col_max[j], grid[i][j]);
            }
        }
        ret += row_max.iter().sum::<i32>();
        ret += col_max.iter().sum::<i32>();
        ret
    }
}

fn main() {
    println!("{:?}", Solution::projection_area(vec![vec![2],]));
}
