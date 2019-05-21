// Accepted
// @lc id=748 lang=rust
// problem: shortest_completing_word
struct Solution {}
impl Solution {
    fn count(s: &str) -> Vec<i32> {
        let mut ret = vec![0; 26];
        s.to_lowercase()
            .chars()
            .filter(|&c| c.is_ascii_alphabetic())
            .for_each(|c| {
                let index = (c as u8 - b'a') as usize;
                ret[index] += 1;
            });
        ret
    }

    pub fn shortest_completing_word(license_plate: String, words: Vec<String>) -> String {
        let mut words = words;
        words.sort_by_key(|w| w.len());
        let counter_a = Solution::count(license_plate.as_str());
        for s in words {
            if Solution::count(s.as_str())
                .iter()
                .enumerate()
                .all(|(i, &v)| counter_a[i] <= v)
            {
                return s;
            }
        }

        return "".to_owned();
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::shortest_completing_word(
            "1s3 PSt".to_owned(),
            vec![
                "step".to_owned(),
                "steps".to_owned(),
                "stripe".to_owned(),
                "stepple".to_owned(),
            ]
        )
    );
}
