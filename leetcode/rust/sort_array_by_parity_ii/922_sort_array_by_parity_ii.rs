// 2018-12-26 15:33
// @lc id=922 lang=rust
struct Solution {}
impl Solution {
    pub fn sort_array_by_parity_ii(a: Vec<i32>) -> Vec<i32> {
        // 2 4 6 8 1 3 5 7
        let mut a = a;
        a.sort_by_key(|x| x % 2);
        let mut ret = vec![];
        for (x, y) in a.iter().take(a.len() / 2).zip(a.iter().skip(a.len() / 2)) {
            ret.push(*x);
            ret.push(*y);
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::sort_array_by_parity_ii(vec![4, 2, 5, 7]));
}
