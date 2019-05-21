// 2018-12-26 11:41
// @lc id=617 lang=rust
mod BinaryTree;
type TreeNode = BinaryTree::TreeNode;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn merge_trees(
        t1: Option<Rc<RefCell<TreeNode>>>,
        t2: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if t1 == None && t2 == None {
            return None;
        }
        let t1_val = {
            if let Some(ref x) = t1 {
                x.borrow().val
            } else {
                0
            }
        };
        let t2_val = {
            if let Some(ref x) = t2 {
                x.borrow().val
            } else {
                0
            }
        };
        let root = Rc::new(RefCell::new(TreeNode::new(t1_val + t2_val)));
        let (t1_left, t1_right) = {
            if let Some(ref x) = t1 {
                (x.borrow().left.clone(), x.borrow().right.clone())
            } else {
                (None, None)
            }
        };

        let (t2_left, t2_right) = {
            if let Some(ref x) = t2 {
                (x.borrow().left.clone(), x.borrow().right.clone())
            } else {
                (None, None)
            }
        };

        root.borrow_mut().left = Solution::merge_trees(t1_left, t2_left);
        root.borrow_mut().right = Solution::merge_trees(t1_right, t2_right);
        return Some(root);
    }
}

fn main() {
    let root1 = BinaryTree::new(vec![Some(1), Some(3), Some(2), Some(5)]);
    let root2 = BinaryTree::new(vec![
        Some(2),
        Some(1),
        Some(3),
        None,
        Some(4),
        None,
        Some(7),
    ]);
    println!("{:?}", Solution::merge_trees(root1, root2));
}
