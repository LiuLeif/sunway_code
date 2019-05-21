// Accepted
// @lc id=985 lang=rust
// problem: sum_of_even_numbers_after_queries
struct Solution {}
impl Solution {
    pub fn sum_even_after_queries(a: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut a = a;
        let mut even = a.iter().filter(|&x| x & 0x1 == 0).sum::<i32>();
        let mut ret = vec![];
        for q in queries {
            let (v, index) = (q[0], q[1] as usize);
            let orig = a[index];
            match (orig & 0x1, (v + orig) & 0x1) {
                (0, 0) => even += v,
                (0, 1) => even -= orig,
                (1, 0) => even += orig + v,
                (1, 1) => {}
                _ => (),
            }
            a[index] += v;
            ret.push(even);
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::sum_even_after_queries(
            vec![1, 2, 3, 4],
            vec![vec![1, 0], vec![-3, 1], vec![-4, 0], vec![2, 3],]
        )
    );
}
