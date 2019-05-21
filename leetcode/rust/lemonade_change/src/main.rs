// Accepted
// @lc id=860 lang=rust
// problem: lemonade_change
struct Solution {}
impl Solution {
    pub fn lemonade_change(bills: Vec<i32>) -> bool {
        let mut count = (0, 0, 0);
        for bill in bills {
            let change = match bill {
                5 => {
                    count.0 += 1;
                    0
                }
                10 => {
                    count.1 += 1;
                    1
                }
                20 => {
                    count.2 += 1;
                    let mut change = 3;
                    if count.1 >= 1 {
                        count.1 -= 1;
                        change -= 2;
                    }
                    change
                }
                _ => 0,
            };
            if count.0 < change {
                return false;
            }
            count.0 -= change;
        }
        true
    }
}

fn main() {
    println!("{:?}", Solution::lemonade_change(vec![5, 5, 5, 10, 20]));
}
