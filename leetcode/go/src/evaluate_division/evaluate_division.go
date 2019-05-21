// 2018-02-12 14:10
package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	ret := []float64{}
	m := map[string]map[string]float64{}
	mem := map[[2]string]float64{}

	for i, equation := range equations {
		f, t := equation[0], equation[1]
		if _, ok := m[f]; !ok {
			m[f] = map[string]float64{}
		}
		m[f][t] = values[i]
		if _, ok := m[t]; !ok {
			m[t] = map[string]float64{}
		}
		m[t][f] = 1.0 / values[i]
	}

	var dfs func(string, string, map[string]bool) (float64, bool)
	dfs = func(f, t string, visited map[string]bool) (float64, bool) {
		if cache, ok := mem[[2]string{f, t}]; ok {
			return cache, true
		}
		if f == t {
			if m[f] != nil {
				return 1, true
			}
			return -1, false
		}
		visited[f] = true
		if children, ok := m[f]; ok {
			for k, v := range children {
				if visited[k] {
					continue
				}
				tmp, ok := dfs(k, t, visited)
				if ok {
					mem[[2]string{f, t}] = v * tmp
					return v * tmp, true
				}
			}
		}
		return -1, false
	}

	for _, query := range queries {
		t, _ := dfs(query[0], query[1], map[string]bool{})
		ret = append(ret, t)
	}
	return ret
}
func main() {
	//      a
	//       \ 2
	//        \
	//         v
	// c<------b
	//     3
	fmt.Println(calcEquation(
		[][]string{
			{"a", "b"},
			{"b", "c"},
		},
		[]float64{2.0, 3.0},
		[][]string{
			{"a", "c"},
			{"b", "a"},
			{"a", "e"},
			{"a", "a"},
			{"x", "x"},
		},
	))
}
