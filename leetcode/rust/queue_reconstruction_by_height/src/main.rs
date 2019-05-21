// Accepted
// 2019-01-04 12:14
// @lc id=406 lang=rust
// problem: queue_reconstruction_by_height
struct Solution {}
impl Solution {
    pub fn reconstruct_queue(people: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut people = people;
        let len = people.len() as i32;
        people.sort_by_key(|v| -(v[0] * len - v[1]));
        let mut ret = vec![];
        for v in people {
            ret.insert(v[1] as usize, v);
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::reconstruct_queue(vec![
            vec![7, 0],
            vec![4, 4],
            vec![7, 1],
            vec![5, 0],
            vec![6, 1],
            vec![5, 2],
        ])
    );
}
