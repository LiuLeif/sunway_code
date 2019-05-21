// 2017-11-12 22:54
package UnionFindSet

type UnionFindSet struct {
	parent map[int]int
	depth  map[int]int
}

func NewUnionFindSet(nums []int) UnionFindSet {
	ret := UnionFindSet{}
	ret.parent = make(map[int]int)
	ret.depth = make(map[int]int)
	for _, v := range nums {
		ret.parent[v] = v
		ret.depth[v] = 1
	}

	return ret
}

func (this *UnionFindSet) Union(a, b int) {
	x := this.Find(a)
	y := this.Find(b)
	if this.depth[x] > this.depth[y] {
		this.parent[y] = x
	} else if this.depth[x] < this.depth[y] {
		this.parent[x] = y
	} else {
		this.parent[x] = y
		this.depth[y]++
	}
}

func (this *UnionFindSet) Find(a int) int {
	if this.parent[a] == a {
		return a
	}
	return this.Find(this.parent[a])
}
