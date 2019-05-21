// Accepted
// @lc id=977 lang=rust
// problem: squares_of_a_sorted_array
struct Solution {}
impl Solution {
    pub fn sorted_squares(a: Vec<i32>) -> Vec<i32> {
        let mut ret = vec![0; a.len()];
        let (minus, positive): (Vec<i32>, Vec<i32>) = a.into_iter().partition(|&x| x < 0);
        let minus = minus.into_iter().rev().map(|x| x * x).collect::<Vec<i32>>();
        let positive = positive.into_iter().map(|x| x * x).collect::<Vec<i32>>();

        let (mut minus, mut positive) = (minus.as_slice(), positive.as_slice());
        for i in 0..ret.len() {
            let tmp;
            if minus.is_empty() || ((!positive.is_empty()) && positive[0] < minus[0]) {
                tmp = positive[0];
                positive = &positive[1..]
            } else {
                tmp = minus[0];
                minus = &minus[1..]
            }
            ret[i] = tmp;
        }
        ret
    }
}
fn main() {
    println!("{:?}", Solution::sorted_squares(vec![-7, -3, 2, 3, 11]));
}
