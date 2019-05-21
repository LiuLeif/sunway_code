// Accepted
// @lc id=24 lang=java
// problem: swap_nodes_in_pairs

class Solution {
    ListNode[] doSwapPairs(ListNode head) {
        if (head == null) {
            return new ListNode[] {null, null};
        }
        if (head.next == null) {
            return new ListNode[] {head, head};
        }
        ListNode newHead = head.next;
        ListNode nextHead = newHead.next;
        newHead.next = head;
        head.next = null;

        ListNode[] next = doSwapPairs(nextHead);
        head.next = next[0];
        return new ListNode[] {newHead, next[1]};
    }
    public ListNode swapPairs(ListNode head) {
        //  1 2 3 4
        return doSwapPairs(head)[0];
    }
}
public class SwapNodesInPairs {
    public static void main(String[] argv) {
        ListNode head = new ListNode(1);
        head.next = new ListNode(2);
        head.next.next = new ListNode(3);
        head.next.next.next = new ListNode(4);
        ListNode newHead = new Solution().swapPairs(head);
        System.out.printf("%d %d %d %d", newHead.val, newHead.next.val, newHead.next.next.val, newHead.next.next.next.val);
    }
}
