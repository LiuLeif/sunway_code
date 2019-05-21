// 2018-01-03 14:30
package main

import (
	"fmt"
	"sort"
)

func helper(m map[string][]string, visited map[[2]string]int, start string, tickets int, record []string) []string {
	if tickets == 0 {
		return record
	}
	for i := 0; i < len(m[start]); i++ {
		next := m[start][i]
		if visited[[2]string{start, next}] == 0 {
			continue
		}
		visited[[2]string{start, next}]--
		record = append(record, next)
		ret := helper(m, visited, next, tickets-1, record)
		if ret != nil {
			return ret
		}
		record = record[:len(record)-1]
		visited[[2]string{start, next}]++
	}
	return nil
}

func findItinerary(tickets [][]string) []string {
	// start := "JFK"
	m := map[string][]string{}
	visited := map[[2]string]int{}
	for i := 0; i < len(tickets); i++ {
		if m[tickets[i][0]] == nil {
			m[tickets[i][0]] = []string{}
		}
		visited[[2]string{tickets[i][0], tickets[i][1]}]++
		m[tickets[i][0]] = append(m[tickets[i][0]], tickets[i][1])
	}
	for _, v := range m {
		sort.Strings(v)
	}
	return helper(m, visited, "JFK", len(tickets), []string{"JFK"})
}
func main() {
	fmt.Println(findItinerary([][]string{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"LHR", "SFO"},
	}))
	fmt.Println(findItinerary([][]string{
		{"JFK", "SFO"},
		{"JFK", "ATL"},
		{"SFO", "ATL"},
		{"ATL", "JFK"},
		{"ATL", "SFO"},
	}))
	fmt.Println(findItinerary([][]string{
		{"JFK", "KUL"},
		{"JFK", "NRT"},
		{"NRT", "JFK"},
	}))
	fmt.Println(findItinerary([][]string{
		{"EZE", "AXA"},
		{"TIA", "ANU"},
		{"ANU", "JFK"},
		{"JFK", "ANU"},
		{"ANU", "EZE"},
		{"TIA", "ANU"},
		{"AXA", "TIA"},
		{"TIA", "JFK"},
		{"ANU", "TIA"},
		{"JFK", "TIA"},
	}))
}
