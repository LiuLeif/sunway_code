// Accepted
import java.util.Arrays;
// @lc id=26 lang=java
// problem: remove_duplicates_from_sorted_array
class Solution {
    public int removeDuplicates(int[] nums) {
        if (nums.length == 0) {
            return 0;
        }
        int lo = 0;
        for (int i = 0; i<nums.length; i++) {
            if (nums[i] != nums[lo]) {
                nums[ ++lo] = nums[i];
            }
        }
        return lo + 1;
    }
}
public class RemoveDuplicatesFromSortedArray {
    public static void main(String[] argv) {
        int [] nums = {1, 2, 3, 3, 3, 4};
        new Solution().removeDuplicates(nums);
        System.out.println(Arrays.toString(nums));
    }
}
