// Accepted
// @lc id=223 lang=rust
// problem: rectangle_area
struct Solution {}
impl Solution {
    pub fn compute_area(a: i32, b: i32, c: i32, d: i32, e: i32, f: i32, g: i32, h: i32) -> i32 {
        //       +---------+ cd
        //       |         |     gh              +---------+
        //       |    +----+-----+               |         |       +-------+
        //       |    |    |     |          +----+----+    |       |       |
        //       +----+----+     |          |    |    |    |       |       | +--------+
        //      ab    |          |          |    +----+----+       +-------+ |        |
        //            +----------+          |         |                      |        |
        //           ef                     +---------+                      +--------+
        let (x, y) = (std::cmp::max(a, e), std::cmp::max(b, f));
        let (m, n) = (std::cmp::min(c, g), std::cmp::min(d, h));
        let delta = {
            if m <= x || n <= y {
                0
            } else {
                (m - x) * (n - y)
            }
        };
        (c - a) * (d - b) + (g - e) * (h - f) - delta
    }
}

fn main() {
    println!("{:?}", Solution::compute_area(-3, 0, 3, 4, 0, -1, 9, 2));
}
