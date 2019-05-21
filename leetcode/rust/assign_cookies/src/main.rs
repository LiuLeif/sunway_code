// Accepted
// @lc id=455 lang=rust
// problem: assign_cookies
struct Solution {}
impl Solution {
    pub fn find_content_children(g: Vec<i32>, s: Vec<i32>) -> i32 {
        // 7,8,9,10
        // 5,6,7,8
        //
        let mut g = g;
        let mut s = s;
        g.sort();
        s.sort();
        let (mut a, mut b, mut ret) = (0, 0, 0);
        while a < g.len() && b < s.len() {
            if s[b] >= g[a] {
                ret += 1;
                a += 1;
            }
            b += 1;
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_content_children(vec![10, 9, 8, 7], vec![5, 6, 7, 8])
    );
}
