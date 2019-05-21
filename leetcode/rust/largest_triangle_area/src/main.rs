// Accepted
// 2019-01-07 15:23
// @lc id=812 lang=rust
// problem: largest_triangle_area
struct Solution {}
impl Solution {
    fn area(x: &Vec<i32>, y: &Vec<i32>) -> f64 {
        // x1,y1
        // x2,y2
        // x3,y3
        let a = x
            .into_iter()
            .zip(y.into_iter().cycle().skip(1))
            .fold(0.0, |sum, (x, y)| sum + (x * y) as f64);
        let b = y
            .into_iter()
            .zip(x.into_iter().cycle().skip(1))
            .fold(0.0, |sum, (x, y)| sum + (x * y) as f64);
        return (a - b).abs() / 2.0;
    }

    fn max(a: f64, b: f64) -> (f64) {
        if a > b {
            return a;
        }
        return b;
    }

    pub fn largest_triangle_area(points: Vec<Vec<i32>>) -> f64 {
        let n = points.len();
        let mut max_area = 0.0;
        for i in 0..n {
            for j in 0..n {
                for k in 0..n {
                    if i == j || j == k || i == k {
                        continue;
                    }
                    max_area = Solution::max(
                        max_area,
                        Solution::area(
                            &vec![points[i][0], points[j][0], points[k][0]],
                            &vec![points[i][1], points[j][1], points[k][1]],
                        ),
                    );
                }
            }
        }
        return max_area;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::largest_triangle_area(vec![
            vec![0, 0],
            vec![0, 1],
            vec![1, 0],
            vec![0, 2],
            vec![2, 0],
        ])
    );
}
