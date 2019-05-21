// Accepted
// @lc id=149 lang=rust
// problem: max_points_on_a_line
struct Solution {}
use std::collections::HashMap;

// DUMMY
#[derive(Debug, PartialEq, Eq)]
pub struct Point {
    pub x: i32,
    pub y: i32,
}

impl Point {
    #[inline]
    pub fn new(x: i32, y: i32) -> Self {
        Point { x, y }
    }
}
// DUMMY

impl Solution {
    fn gcd(mut a: i32, mut b: i32) -> i32 {
        while a != 0 {
            let tmp = a;
            a = b % a;
            b = tmp;
        }
        b
    }

    fn get_slop(p: &Point, q: &Point) -> (i32, i32) {
        if p.x == q.x {
            return (1, 0);
        }
        if p.y == q.y {
            return (0, 1);
        }
        let (mut delta_x, mut delta_y) = (p.x - q.x, p.y - q.y);
        let is_negative = delta_x * delta_y < 0;
        delta_x = delta_x.abs();
        delta_y = delta_y.abs();
        let gcd = Solution::gcd(delta_x, delta_y);
        delta_x /= gcd;
        delta_y /= gcd;
        if is_negative {
            delta_x = -delta_x
        }
        (delta_x, delta_y)
    }

    pub fn max_points(points: Vec<Point>) -> i32 {
        if points.is_empty() {
            return 0;
        }
        let mut ret = 0;
        for (i, p) in points.iter().enumerate() {
            let mut counter = HashMap::<(i32, i32), i32>::new();
            let mut dup = 0;
            for (j, q) in points.iter().enumerate() {
                if i == j {
                    continue;
                }
                if p == q {
                    dup += 1;
                    continue;
                }
                let slop = Solution::get_slop(&p, &q);
                *(counter.entry(slop).or_insert(0)) += 1;
            }
            ret = std::cmp::max(ret, dup + counter.values().max().unwrap_or(&0));
        }
        ret + 1
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::max_points(vec![
            Point::new(84, 250),
            Point::new(0, 0),
            Point::new(1, 0),
            Point::new(0, -70),
            Point::new(0, -70),
            Point::new(1, -1),
            Point::new(21, 10),
            Point::new(42, 90),
            Point::new(-42, -230),
        ])
    );
}
