// Accepted
// @lc id=210 lang=rust
// problem: course_schedule_ii
// tags: topo_sort
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let mut ingress = vec![0; num_courses as usize];
        let mut conn = vec![vec![]; num_courses as usize];
        let mut ret = vec![];
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
            ret.push(top as i32);
            for x in conn[top].clone() {
                ingress[x as usize] -= 1;
                if ingress[x as usize] == 0 {
                    queue.push_back(x);
                }
            }
        }
        if ingress.into_iter().all(|x| x == 0) {
            ret.reverse();
            ret
        } else {
            vec![]
        }
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_order(4, vec![vec![1, 0], vec![2, 0], vec![3, 1], vec![3, 2]])
    );
}
