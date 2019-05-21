// 2018-12-27 10:12
// @lc id=728 lang=rust
struct Solution {}
struct NumIter(i32);

impl iterator for NumIter {
    type Item = i32;
    fn next(&mut self) -> Option<i32> {
        if self.0 == 0 {
            return None;
        }
        let ret = self.0 % 10;
        self.0 /= 10;
        Some(ret)
    }
}
impl Solution {
    pub fn self_dividing_numbers(left: i32, right: i32) -> Vec<i32> {
        fn is_self_dividing(x: i32) -> bool {
            NumIter(x).all(|y| y != 0 && x % y == 0)
        }

        return (left..right + 1).filter(|&x| is_self_dividing(x)).collect();
    }
}

fn main() {
    println!("{:?}", Solution::self_dividing_numbers(1, 22));
}
