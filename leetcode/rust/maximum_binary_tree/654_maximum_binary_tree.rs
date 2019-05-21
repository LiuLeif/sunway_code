// 2018-12-25 15:38
// @lc id=654 lang=rust
mod BinaryTree;
type TreeNode = BinaryTree::TreeNode;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    fn do_construct(nums: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
        if nums.len() == 0 {
            return None;
        }
        let largestIndex = {
            let mut ret = 0;
            for i in 0..nums.len() {
                if nums[i] > nums[ret] {
                    ret = i
                }
            }
            ret
        };
        let root = Rc::new(RefCell::new(TreeNode::new(nums[largestIndex])));
        if nums.len() == 1 {
            return Some(root);
        }
        root.borrow_mut().left = Solution::do_construct(&nums[..largestIndex]);
        root.borrow_mut().right = Solution::do_construct(&nums[largestIndex + 1..]);
        return Some(root);
    }

    pub fn construct_maximum_binary_tree(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        return Solution::do_construct(&nums);
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::construct_maximum_binary_tree(vec![3, 2, 1, 6, 0, 5])
    );
}
