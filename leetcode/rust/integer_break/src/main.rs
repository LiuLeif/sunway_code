// Accepted
// @lc id=343 lang=rust
// problem: integer_break
struct Solution {}
impl Solution {
    fn dfs(cache: &mut [i32], n: usize) -> i32 {
        if n <= 4 {
            return n as i32;
        }
        if cache[n] != 0 {
            return cache[n as usize];
        }
        let ret = (2..n)
            .map(|x| x as i32 * Solution::dfs(cache, n - x))
            .max()
            .unwrap();
        cache[n] = ret;
        ret
    }

    pub fn integer_break(n: i32) -> i32 {
        match n {
            2 => return 1,
            3 => return 2,
            _ => (),
        }
        // 10=3+3+4
        return Solution::dfs(&mut vec![0; n as usize + 1], n as usize);
    }
}

fn main() {
    println!("{:?}", Solution::integer_break(10));
}
