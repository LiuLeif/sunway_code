// 2018-12-03 15:20
package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
type Stack []*NestedInteger

func (this *Stack) Push(x *NestedInteger) {
	*this = append(*this, x)
}

func (this *Stack) Pop() *NestedInteger {
	ret := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return ret
}

func (this Stack) Empty() bool {
	return len(this) == 0
}

func (this Stack) Peek() *NestedInteger {
	return this[len(this)-1]
}

func deserialize(s string) *NestedInteger {
	// [[]]
	stack := Stack{}
	stack.Push(&NestedInteger{})
	hasDigit := false
	digit := 0
	minus := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			stack.Push(&NestedInteger{})
		case ']':
			if hasDigit {
				tmp := &NestedInteger{}
				if minus {
					digit = -digit
				}
				tmp.SetInteger(digit)
				digit = 0
				minus = false
				hasDigit = false
				stack.Peek().Add(*tmp)
			}
			tmp := stack.Pop()
			stack.Peek().Add(*tmp)
		case ',':
			if hasDigit {
				tmp := &NestedInteger{}
				if minus {
					digit = -digit
				}
				tmp.SetInteger(digit)
				digit = 0
				minus = false
				hasDigit = false
				stack.Peek().Add(*tmp)
			}
		case '-':
			minus = true
		default:
			hasDigit = true
			digit = digit*10 + int(s[i]-'0')
		}
	}
	if hasDigit {
		tmp := &NestedInteger{}
		if minus {
			digit = -digit
		}
		tmp.SetInteger(digit)
		return tmp
	}
	return stack.Peek().GetList()[0]
}

// type NestedInteger struct {
// }

// func (n NestedInteger) IsInteger() bool {
// 	return true
// }

// func (n NestedInteger) GetInteger() int {
// 	return 0
// }

// func (n *NestedInteger) SetInteger(value int) {}

// func (n *NestedInteger) Add(elem NestedInteger) {}

// func (n NestedInteger) GetList() []*NestedInteger {
// 	return nil
// }

func main() {

}
