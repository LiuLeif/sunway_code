// 2019-01-02 18:39
// @lc id=868 lang=rust
// problem: binary_gap
struct Solution {}
struct BitIter(i32, i32);
impl Iterator for BitIter {
    type Item = i32;

    fn next(&mut self) -> Option<i32> {
        if self.0 == 0 {
            return None;
        }
        while self.0 % 2 == 0 {
            self.1 += 1;
            self.0 /= 2;
        }
        self.1 += 1;
        self.0 /= 2;
        return Some(self.1 - 1);
    }
}
impl BitIter {
    fn new(n: i32) -> BitIter {
        return BitIter(n, 0);
    }
}
impl Solution {
    pub fn binary_gap(n: i32) -> i32 {
        // 101
        // 0,2,4
        let ret = BitIter::new(n)
            .collect::<Vec<_>>()
            .windows(2)
            .map(|w| w[1] - w[0])
            .max();
        match ret {
            Some(x) => x,
            _ => 0,
        }
    }
}

fn main() {
    println!("{:?}", Solution::binary_gap(5));
}
