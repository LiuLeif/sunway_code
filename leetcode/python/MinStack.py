#  Accepted
#
# @lc app=leetcode id=155 lang=python
#
# [155] Min Stack
#

# @lc code=start
class MinStack(object):
    def __init__(self):
        """
        initialize your data structure here.
        """
        self.data = []
        self.min_data = []

    def push(self, x):
        """
        :type x: int
        :rtype: None
        """
        self.data.append(x)
        if len(self.min_data) == 0:
            self.min_data.append(x)
        else:
            self.min_data.append(min(self.min_data[-1], x))

    def pop(self):
        """
        :rtype: None
        """
        self.data = self.data[:-1]
        self.min_data = self.min_data[:-1]

    def top(self):
        """
        :rtype: int
        """
        return self.data[-1]

    def getMin(self):
        """
        :rtype: int
        """
        return self.min_data[-1]


# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()
# @lc code=end


if __name__ == "__main__":
    stack = MinStack()
    stack.push(1)
    stack.pop()
    stack.push(1)
    stack.push(2)
    print(stack.top())
    stack.pop()
    print(stack.getMin())
