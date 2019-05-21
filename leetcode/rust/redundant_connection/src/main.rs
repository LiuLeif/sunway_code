// Accepted
// @lc id=684 lang=rust
// problem: redundant_connection
use std::collections::HashMap;
struct Solution {}
struct UFS {
    parent: HashMap<i32, i32>,
}
impl UFS {
    fn find(&mut self, mut x: i32) -> i32 {
        self.parent.entry(x).or_insert(x);
        while self.parent[&x] != x {
            x = self.parent[&x];
        }
        return x;
    }

    fn union(&mut self, x: i32, y: i32) {
        let (px, py) = (self.find(x), self.find(y));
        if px == py {
            return;
        }
        self.parent.insert(px, py);
    }
    fn new() -> UFS {
        UFS {
            parent: HashMap::new(),
        }
    }
}
impl Solution {
    pub fn find_redundant_connection(edges: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ufs = UFS::new();
        for edge in edges {
            let (a, b) = (edge[0], edge[1]);
            if ufs.find(a) == ufs.find(b) {
                return edge;
            }
            ufs.union(a, b);
        }
        return vec![];
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_redundant_connection(vec![
            vec![1, 2],
            vec![2, 3],
            vec![3, 4],
            vec![1, 4],
            vec![1, 5],
        ])
    );
    println!(
        "{:?}",
        Solution::find_redundant_connection(vec![vec![1, 2], vec![1, 3], vec![2, 3],])
    );
}
