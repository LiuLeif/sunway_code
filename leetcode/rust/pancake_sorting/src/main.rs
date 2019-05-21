// Accepted
// 2019-01-07 13:37
// @lc id=969 lang=rust
// problem: pancake_sorting
struct Solution {}
impl Solution {
    pub fn pancake_sort(a: Vec<i32>) -> Vec<i32> {
        // 3,2,4,1,
        let mut ret = vec![];
        let mut a = a;
        for i in (1..a.len()).rev() {
            // find max index of [0..i]
            let argmax = a[..i + 1].iter().enumerate().max_by_key(|t| t.1).unwrap().0;
            if argmax == i {
                continue;
            }
            if argmax != 0 {
                ret.push(argmax as i32 + 1);
            }
            ret.push(i as i32 + 1);

            let mut tmp = a[argmax + 1..i + 1].to_vec();
            tmp.reverse();
            tmp.extend(&a[..argmax]);
            a = tmp;
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::pancake_sort(vec![2, 3, 1, 4]));
    println!("{:?}", Solution::pancake_sort(vec![1, 3, 2]));
}
