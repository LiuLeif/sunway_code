// Accepted
// @lc id=71 lang=rust
// problem: simplify_path
struct Solution {}
impl Solution {
    pub fn simplify_path(path: String) -> String {
        let mut stack = vec![];
        for p in path.split("/") {
            match p {
                "." | "" => (),
                ".." => {
                    stack.pop();
                }
                _ => {
                    stack.push(p);
                }
            }
        }
        "/".to_owned() + &stack.join("/")
    }
}

fn main() {
    println!("{:?}", Solution::simplify_path("/../c/".to_owned()));
}
