// 2018-12-28 15:07
// @lc id=867 lang=rust
struct Solution {}
impl Solution {
    pub fn transpose(a: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut ret = vec![vec![0; a.len()]; a[0].len()];
        for i in 0..a.len() {
            for j in 0..a[0].len() {
                ret[j][i] = a[i][j];
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::transpose(vec![vec![1, 2, 3]]));
}
