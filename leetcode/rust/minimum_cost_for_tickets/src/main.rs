// Accepted
// @lc id=983 lang=rust
// problem: minimum_cost_for_tickets
struct Solution {}
impl Solution {
    fn first_larger(days: &[i32], target: i32) -> usize {
        let (mut lo, mut hi) = (0, days.len() as i32 - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if days[mid as usize] > target {
                hi = mid - 1
            } else {
                lo = mid + 1;
            }
        }
        lo as usize
    }

    fn dfs(cache: &mut Vec<i32>, days: &[i32], costs: &[i32], start: usize) -> i32 {
        if start == days.len() {
            return 0;
        }
        if cache[start] != 0 {
            return cache[start];
        }
        let ret = vec![
            costs[0] + Solution::dfs(cache, days, costs, start + 1),
            costs[1] + Solution::dfs(
                cache,
                days,
                costs,
                Solution::first_larger(days, days[start] + 6),
            ),
            costs[2] + Solution::dfs(
                cache,
                days,
                costs,
                Solution::first_larger(days, days[start] + 29),
            ),
        ]
        .into_iter()
        .min()
        .unwrap();
        cache[start] = ret;
        ret
    }

    pub fn mincost_tickets(days: Vec<i32>, costs: Vec<i32>) -> i32 {
        let mut cache = vec![0; days.len()];
        Solution::dfs(&mut cache, days.as_slice(), costs.as_slice(), 0)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::mincost_tickets(vec![1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31], vec![2, 7, 15])
    );
}
