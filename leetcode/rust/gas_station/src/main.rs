// Accepted
// @lc id=134 lang=rust
// problem: gas_station
struct Solution {}
impl Solution {
    pub fn can_complete_circuit(gas: Vec<i32>, cost: Vec<i32>) -> i32 {
        let n = gas.len();
        let mut start = 0;
        'restart: while start < n {
            let mut tank = 0;
            for i in 0..n {
                let index = (start + i) % n;
                tank += gas[index] - cost[index];
                if tank < 0 {
                    start += i + 1;
                    continue 'restart;
                }
            }
            return start as i32;
        }
        -1
    }
}
// -1 -1 1
fn main() {
    println!(
        "{:?}",
        Solution::can_complete_circuit(vec![2, 3, 4], vec![3, 4, 3])
    );
}
