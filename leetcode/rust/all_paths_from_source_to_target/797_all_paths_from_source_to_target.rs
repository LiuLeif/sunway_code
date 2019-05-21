// 2018-12-28 17:11
// @lc id=797 lang=rust
struct Solution {}
impl Solution {
    fn dfs(
        record: &mut Vec<Vec<i32>>,
        curr: &mut Vec<i32>,
        f: i32,
        t: i32,
        graph: &Vec<Vec<i32>>,
    ) -> () {
        if f == t {
            record.push(curr.clone());
            return;
        }
        graph[f as usize].iter().for_each(|n| {
            curr.push(*n);
            Solution::dfs(record, curr, *n, t, graph);
            curr.pop();
        })
    }

    pub fn all_paths_source_target(graph: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        // 0--->1
        // |    |
        // v    v
        // 2--->3
        Solution::dfs(&mut ret, &mut vec![0], 0, graph.len() as i32 - 1, &graph);
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::all_paths_source_target(vec![vec![1, 2,], vec![3], vec![3], vec![],])
    );
}
