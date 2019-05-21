// Accepted
import java.util.HashMap;
import java.util.Arrays;
// @lc id=1 lang=java
// problem: two_sum
class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> counter = new HashMap();
        for (int i = nums.length - 1; i >= 0; i-- ) {
            counter.putIfAbsent(nums[i], i);
        }
        for (int i = 0; i<nums.length; i++) {
            Integer tmp = counter.get(target - nums[i]);
            if (tmp != null && tmp != i) {
                return new int[] {i, tmp};
            }
        }
        return new int[2];
    }
}
public class TwoSum {
    public static void main(String[] argv) {
        System.out.println(Arrays.toString(new Solution().twoSum(new int[] {3, 2, 4}, 7)));
    }
}
