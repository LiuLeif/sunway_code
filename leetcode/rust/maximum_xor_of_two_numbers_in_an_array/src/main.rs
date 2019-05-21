// Accepted
// @lc id=421 lang=rust
// problem: maximum_xor_of_two_numbers_in_an_array
use std::cell::RefCell;
use std::rc::Rc;
struct Solution {}

type Link = Option<Rc<RefCell<TrieNode>>>;

struct TrieNode {
    left: Link,
    right: Link,
    value: i32,
}

struct TrieTree {
    root: Link,
}

impl TrieNode {
    fn new() -> Link {
        Some(Rc::new(RefCell::new(TrieNode {
            left: None,
            right: None,
            value: -1,
        })))
    }
}

impl TrieTree {
    fn new() -> TrieTree {
        TrieTree {
            root: TrieNode::new(),
        }
    }

    fn add(&self, val: i32) {
        let mut root = self.root.clone().unwrap();
        for i in (0..32).rev() {
            if val & (1 << i) != 0 {
                if root.borrow().left.is_none() {
                    root.borrow_mut().left = TrieNode::new();
                }
                let left = root.borrow().left.clone().unwrap();
                root = left;
            } else {
                if root.borrow().right.is_none() {
                    root.borrow_mut().right = TrieNode::new();
                }
                let right = root.borrow().right.clone().unwrap();
                root = right;
            }
        }
        root.borrow_mut().value = val;
    }

    fn get_max_value(&self, nums: &[i32]) -> i32 {
        let mut ret = 0;
        for num in nums {
            let mut root = self.root.clone().unwrap();
            for i in (0..32).rev() {
                let bit = num & (1 << i);
                let left = root.borrow().left.clone();
                let right = root.borrow().right.clone();
                if bit != 0 {
                    root = right.unwrap_or_else(|| left.unwrap());
                } else {
                    root = left.unwrap_or_else(|| right.unwrap());
                }
            }
            ret = std::cmp::max(ret, num ^ root.borrow().value);
        }
        ret
    }
}

impl Solution {
    pub fn find_maximum_xor(nums: Vec<i32>) -> i32 {
        // 5 ^ 25 = 28
        let tree = &TrieTree::new();
        for i in nums.iter() {
            tree.add(*i);
        }
        return tree.get_max_value(&nums);
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_maximum_xor(vec![
            2604, 6991, 8010, 8260, 5136, 8770, 8568, 3021, 1621, 7303, 9289, 9855, 7267, 7872,
            1163, 2605, 1438, 5683, 7245, 3647, 7977, 2361, 8355, 3215, 1784, 1189, 9334, 4099,
            2467, 9342, 5512, 2823, 1254, 6315, 645, 3638, 5077, 8442, 5957, 501, 8045, 1790, 7357,
            7180, 8614, 2899, 3222, 2994, 228, 5662, 8421, 7060, 669, 9413, 9353, 9697, 9141, 3688,
            580, 8309, 7285, 4091, 9712, 3296, 5322, 7490, 4899, 6083, 9961, 2392, 685, 1862, 5664,
            2667, 6414, 2888, 3624, 5495, 4677, 2932, 9942, 6013, 1235, 6733, 8609, 5504, 9181,
            6886, 7822, 2125, 4800, 1760, 8632, 468, 4927, 3009, 7627, 4167, 4296, 5437
        ])
    );
}
