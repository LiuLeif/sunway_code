// Accepted
// @lc id=303 lang=rust
// problem: range_sum_query_immutable
struct NumArray {
    sum: Vec<i32>,
}

impl NumArray {
    fn new(nums: Vec<i32>) -> Self {
        let mut sum = vec![0; nums.len() + 1];
        for i in 1..nums.len() + 1 {
            sum[i] = sum[i - 1] + nums[i - 1];
        }
        NumArray { sum }
    }

    fn sum_range(&self, i: i32, j: i32) -> i32 {
        self.sum[j as usize + 1] - self.sum[i as usize]
    }
}

fn main() {
    let array = NumArray::new(vec![-2, 0, 3, -5, 2, -1]);
    println!("{:?}", array.sum_range(0, 1)); // 1
    println!("{:?}", array.sum_range(2, 5)); // -1
    println!("{:?}", array.sum_range(0, 5)); // -3
}
