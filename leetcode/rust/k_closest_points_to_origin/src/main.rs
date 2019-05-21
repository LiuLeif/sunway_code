// Accepted
// @lc id=973 lang=rust
// problem: k_closest_points_to_origin
struct Solution {}
impl Solution {
    fn distance(p: &[i32]) -> i32 {
        p[0] * p[0] + p[1] * p[1]
    }

    fn partition(points: &mut [Vec<i32>]) -> usize {
        let mut p = 0;
        let mut f = 1;
        let t = points.len() - 1;
        while f <= t {
            if Solution::distance(&points[f]) >= Solution::distance(&points[p]) {
                f += 1;
            } else {
                points.swap(p + 1, f);
                points.swap(p, p + 1);
                p += 1;
                f += 1;
            }
        }
        p
    }

    fn quick_select(points: &mut [Vec<i32>], k: usize) {
        if points.len() <= 1 {
            return;
        }
        let p = Solution::partition(points);
        if p == k {
            return;
        } else if p > k {
            Solution::quick_select(&mut points[..p], k);
        } else {
            Solution::quick_select(&mut points[p + 1..], k - p - 1);
        }
    }

    pub fn k_closest(points: Vec<Vec<i32>>, k: i32) -> Vec<Vec<i32>> {
        let mut points = points;
        let len = points.len();
        Solution::quick_select(points.as_mut_slice(), k as usize);
        points[..k as usize].to_vec()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::k_closest(
            vec![
                vec![5, -3],
                vec![3, -4],
                vec![-2, 9],
                vec![-8, -8],
                vec![-6, 5],
                vec![6, 1],
                vec![5, 9],
            ],
            6
        )
    );
}
