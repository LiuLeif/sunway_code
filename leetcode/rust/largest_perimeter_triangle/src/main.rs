// Accepted
// @lc id=976 lang=rust
// problem: largest_perimeter_triangle
struct Solution {}
impl Solution {
    pub fn largest_perimeter(a: Vec<i32>) -> i32 {
        let mut a = a;
        a.sort_by_key(|i| -i);
        a.windows(3)
            .skip_while(|w| w[0] >= w[1] + w[2])
            .next()
            .map_or(0, |w| w[0] + w[1] + w[2])
    }
}

fn main() {
    // 2,3,3,4
    println!("{:?}", Solution::largest_perimeter(vec![3, 2, 3, 4]));
}
