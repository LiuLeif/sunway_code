// Accepted
// @lc id=48 lang=java
// problem: rotate_image
class Solution {
    public void rotate(int[][] matrix) {
        // 1 2 3
        // 4 5 6
        // 7 8 9
        int n = matrix.length;
        for (int i = 0; i < n; i++) {
            for (int j = i; j < n ; j++ ) {
                int hold = matrix[i][j];
                matrix[i][j] = matrix[j][i];
                matrix[j][i] = hold;
            }
        }
        for (int j = 0; j < n / 2; j++) {
            for (int i = 0; i < n; i++) {
                int hold = matrix[i][j];
                matrix[i][j] = matrix[i][n - j - 1];
                matrix[i][n - j - 1] = hold;
            }
        }
    }
}
public class RotateImage {
    public static void main(String[] argv) {
        new Solution().rotate(new int[][] {{1,2,3},{4,5,6},{7,8,9}});

    }
}
