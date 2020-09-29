#  Accepted
#
# @lc app=leetcode id=58 lang=python
#
# [58] Length of Last Word
#

# @lc code=start
class Solution(object):
    def lengthOfLastWord(self, s):
        """
        :type s: str
        :rtype: int
        """
        s = [x for x in s.split(" ") if x != ""]
        if len(s) == 0:
            return 0
        return len(s[-1])


# @lc code=end


if __name__ == "__main__":
    print(Solution().lengthOfLastWord("   "))
