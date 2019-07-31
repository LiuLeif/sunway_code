// Accepted
// @lc id=1103 lang=rust
// problem: distribute_candies_to_people
struct Solution {}
impl Solution {
    pub fn distribute_candies(candies: i32, num_people: i32) -> Vec<i32> {
        let mut ret = vec![0; num_people as usize];
        let mut candies = candies;
        let mut current = 1;
        while candies >= current {
            ret[((current - 1) % num_people) as usize] += current;
            candies -= current;
            current += 1;
        }
        ret[((current - 1) % num_people) as usize] += candies;
        ret
    }
}

fn main() {
    println!("{:?}", Solution::distribute_candies(10, 10));
}
