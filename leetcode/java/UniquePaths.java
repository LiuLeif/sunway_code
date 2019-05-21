// Accepted
// @lc id=62 lang=java
// problem: unique_paths
class Solution {
    public int uniquePaths(int m, int n) {
        int [][] dp = new int [m + 1][n + 1];
        for (int i = 1; i < m + 1; i++) {
            for (int j = 1; j < n + 1; j++) {
                if (i == 1 && j == 1) {
                    dp[i][j] = 1;
                } else {
                    dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
                }
            }
        }
        return dp[m][n];
    }
}
public class UniquePaths {
    public static void main(String[] argv) {
        System.out.println(new Solution().uniquePaths(7, 3));
    }
}
