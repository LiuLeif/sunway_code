// Accepted
// @lc id=102 lang=java
// problem: binary_tree_level_order_traversal
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer >> ret = new ArrayList();
        if (root == null) {
            return ret;
        }
        Deque<TreeNode> queue = new ArrayDeque();
        queue.addLast(root);
        while (!queue.isEmpty()) {
            List<Integer> row = new ArrayList();
            int len = queue.size();
            for (int i = 0; i<len; i++) {
                TreeNode top = queue.pollFirst();
                row.add(top.val);
                if (top.left != null) {
                    queue.addLast(top.left);
                }
                if (top.right != null) {
                    queue.addLast(top.right);
                }
            }
            ret.add(row);
        }
        return ret;
    }
}
