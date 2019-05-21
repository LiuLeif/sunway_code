// Accepted
// @lc id=1029 lang=rust
// problem: two_city_scheduling
struct Solution {}
impl Solution {
    pub fn two_city_sched_cost(costs: Vec<Vec<i32>>) -> i32 {
        let n = costs.len();
        let mut delta = costs
            .iter()
            .map(|v| v[0] - v[1])
            .enumerate()
            .collect::<Vec<(usize, i32)>>();
        delta.sort_by_key(|v| v.1);
        delta[0..n / 2].iter().map(|v| costs[v.0][0]).sum::<i32>()
            + delta[n / 2..].iter().map(|v| costs[v.0][1]).sum::<i32>()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::two_city_sched_cost(vec![
            vec![10, 20],
            vec![30, 200],
            vec![400, 50],
            vec![30, 20]
        ])
    );
}
