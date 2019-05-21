// Accepted
import java.util.HashMap;
// @lc id=3 lang=java
// problem: longest_substring_without_repeating_characters
class Solution {
    public int lengthOfLongestSubstring(String s) {
        // abcabcbb
        HashMap<Character, Integer> counter = new HashMap();
        int lo = 0;
        int ret = 0;
        for (int i = 0; i<s.length(); i++) {
            Integer prev = counter.get(s.charAt(i));
            if (prev != null) {
                lo = Math.max(lo, prev + 1);
            }
            counter.put(s.charAt(i), i);
            ret = Math.max(ret, i - lo + 1);
        }
        return ret;
    }
}
public class LongestSubstringWithoutRepeatingCharacters {
    public static void main(String[] argv) {
        System.out.println(new Solution().lengthOfLongestSubstring("a"));
    }
}
