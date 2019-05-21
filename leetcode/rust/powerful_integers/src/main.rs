// Accepted
// @lc id=970 lang=rust
// problem: powerful_integers
struct Solution {}
struct MultiIter {
    val: i32,
    multi: i32,
    bound: i32,
}

impl MultiIter {
    fn new(multi: i32, bound: i32) -> Self {
        MultiIter {
            val: 1,
            multi,
            bound,
        }
    }
}

impl Iterator for MultiIter {
    type Item = i32;
    fn next(&mut self) -> Option<i32> {
        if self.val > self.bound {
            return None;
        }
        let ret = self.val;
        if self.multi == 1 {
            self.bound = 0;
        }
        self.val *= self.multi;
        Some(ret)
    }
}
impl Solution {
    pub fn powerful_integers(x: i32, y: i32, bound: i32) -> Vec<i32> {
        let mut mark = vec![false; bound as usize + 1];
        for a in MultiIter::new(x, bound) {
            for b in MultiIter::new(y, bound).take_while(|x| a + x <= bound) {
                mark[(a + b) as usize] = true;
            }
        }
        mark.into_iter()
            .enumerate()
            .filter(|&(index, v)| v)
            .map(|(index, v)| index as i32)
            .collect()
    }
}

fn main() {
    println!("{:?}", Solution::powerful_integers(3, 5, 15));
}
