// Accepted
// @lc id=932 lang=rust
// problem: beautiful_array
struct Solution {}
impl Solution {
    fn transform(from: Vec<i32>) -> Vec<i32> {
        if from.len() <= 2 {
            return from;
        }
        let mut ret = vec![];
        let even = from.iter().step_by(2).cloned().collect::<Vec<i32>>();
        let odd = from
            .iter()
            .skip(1)
            .step_by(2)
            .cloned()
            .collect::<Vec<i32>>();
        ret.extend(Solution::transform(even));
        ret.extend(Solution::transform(odd));
        return ret;
    }

    pub fn beautiful_array(n: i32) -> Vec<i32> {
        // 1,2,3,4,5
        // 1 3 5 | 2 4 6
        return Solution::transform((1..n + 1).collect::<Vec<i32>>());
    }
}

fn main() {
    println!("{:?}", Solution::beautiful_array(10));
}
