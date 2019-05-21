// Accepted
// @lc id=207 lang=rust
// problem: course_schedule
// tags: topo_sort
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let mut ingress = vec![0; num_courses as usize];
        let mut conn = vec![vec![]; num_courses as usize];
        for v in prerequisites {
            conn[v[0] as usize].push(v[1]);
            ingress[v[1] as usize] += 1;
        }
        let mut queue = VecDeque::new();
        for i in 0..num_courses {
            if ingress[i as usize] == 0 {
                queue.push_back(i);
            }
        }
        while !queue.is_empty() {
            let top = queue.pop_front().unwrap() as usize;
            for x in conn[top].clone() {
                ingress[x as usize] -= 1;
                if ingress[x as usize] == 0 {
                    queue.push_back(x);
                }
            }
        }
        ingress.into_iter().all(|x| x == 0)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::can_finish(2, vec![vec![1, 0], vec![0, 1]])
    );
}
