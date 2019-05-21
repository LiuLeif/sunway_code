// Accepted
// @lc id=997 lang=rust
// problem: find_the_town_judge
struct Solution {}
impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        // 1 2 3 1-2 2-3
        let mut a = vec![0; n as usize];
        let mut b = vec![0; n as usize];
        for v in trust {
            a[v[0] as usize - 1] += 1;
            b[v[1] as usize - 1] += 1;
        }
        a.into_iter()
            .zip(b.into_iter())
            .position(|(x, y)| x == 0 && y == n - 1)
            .map_or(-1, |i| i as i32 + 1)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_judge(
            4,
            vec![vec![1, 3], vec![1, 4], vec![2, 3], vec![2, 4], vec![4, 3],]
        )
    );
    println!("{:?}", Solution::find_judge(2, vec![vec![1, 2]]));
}
