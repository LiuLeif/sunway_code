// Accepted
import java.util.ArrayList;
import java.util.List;
// @lc id=46 lang=java
// problem: permutations
class Solution {
    private void backtrack(List<List<Integer >> ret, ArrayList<Integer> curr, int [] nums, boolean [] visited) {
        if (curr.size() == nums.length) {
            ret.add(new ArrayList(curr));
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            if (visited[i]) {
                continue;
            }
            visited[i] = true;
            curr.add(nums[i]);
            backtrack(ret, curr, nums, visited);
            curr.remove(curr.size() - 1);
            visited[i] = false;
        }
    }
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer >> ret = new ArrayList();
        backtrack(ret, new ArrayList<Integer>(), nums, new boolean[nums.length]);
        return ret;
    }
}
public class Permutations {
    public static void main(String[] argv) {
        System.out.println(new Solution().permute(new int[] {1, 2, 3}));
    }
}
