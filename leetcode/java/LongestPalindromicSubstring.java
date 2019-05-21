// Accepted
import java.util.Arrays;
// @lc id=5 lang=java
// problem: longest_palindromic_substring
class Solution {
    public String longestPalindrome(String s) {
        int n = s.length();
        if (n == 0) {
            return s;
        }
        boolean [][] isPalindrome = new boolean[n][n];
        for (int i = 0; i < n; i++) {
            isPalindrome[i][i] = true;
        }
        int maxLen = 0;
        int maxStart = 0;
        for (int step = 1; step <n; step++) {
            for (int i = 0; i <n - step; i++) {
                isPalindrome[i][i + step] = (s.charAt(i) == s.charAt(i + step));
                if (step != 1) {
                    isPalindrome[i][i + step] &= isPalindrome[i + 1][i + step - 1];
                }
                if (isPalindrome[i][i + step] && step > maxLen) {
                    maxLen = step;
                    maxStart = i;
                }
            }
        }
        return s.substring(maxStart, maxStart + maxLen + 1);
    }
}
public class LongestPalindromicSubstring {
    public static void main(String[] argv) {
        System.out.println(new Solution().longestPalindrome("abab"));
    }
}
