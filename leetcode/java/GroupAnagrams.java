// Accepted
// @lc id=49 lang=java
// problem: group_anagrams
import java.util.*;
class Solution {
    public List<List<String>> groupAnagrams(String[] strs) {
        HashMap<String, List<String >> counter = new HashMap();
        for (String s : strs) {
            byte [] tmp = s.getBytes();
            Arrays.sort(tmp);
            String key = new String(tmp);
            counter.putIfAbsent(key, new ArrayList<String>());
            counter.get(key).add(s);
        }
        return new ArrayList(counter.values());
    }
}
public class GroupAnagrams {
    public static void main(String[] argv) {
        System.out.println(new Solution().groupAnagrams(new String[] {"eat", "tea", "tan", "ate", "nat", "bat"}));
    }
}
