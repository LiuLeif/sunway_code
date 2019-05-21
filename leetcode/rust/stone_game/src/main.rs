// 2019-01-02 22:44
// @lc id=877 lang=rust
// problem: stone_game
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn stone_game(piles: Vec<i32>) -> bool {
        // 5,3,4,5
        fn play(
            f: usize,
            t: usize,
            piles: &[i32],
            cache: &mut HashMap<(usize, usize), (i32, i32)>,
        ) -> (i32, i32) {
            if f == t {
                return (piles[f], 0);
            }
            if let Some(c) = cache.get(&(f, t)) {
                return *c;
            }
            let first = play(f + 1, t, piles, cache);
            let last = play(f, t - 1, piles, cache);

            let ret;
            if first.1 + piles[f] > last.1 + piles[t] {
                ret = (first.1 + piles[f], first.0);
            } else {
                ret = (last.1 + piles[t], last.0);
            }
            cache.insert((f, t), ret);
            return ret;
        }
        let ret = play(
            0,
            piles.len() - 1,
            &piles,
            &mut HashMap::<(usize, usize), (i32, i32)>::new(),
        );
        return ret.0 > ret.1;
    }
}

fn main() {
    println!("{:?}", Solution::stone_game(vec![5, 4, 5, 2]));
    println!("{:?}", Solution::stone_game(vec![5, 7, 4, 5, 2]));
}
