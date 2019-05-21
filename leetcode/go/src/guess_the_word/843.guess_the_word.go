// 2018-12-05 17:04
package main

import (
	"math"
)

type Master struct {
	secret string
}

func (this *Master) Guess(word string) int {
	la := len(this.secret)
	ret := 0
	for i := 0; i < la; i++ {
		if this.secret[i] == word[i] {
			ret++
		}
	}
	return ret
}

func distance(a, b string) int {
	la := len(a)
	ret := 0
	for i := 0; i < la; i++ {
		if a[i] == b[i] {
			ret++
		}
	}
	return ret
}

func max(nums ...int) int {
	ret := math.MinInt32
	for _, n := range nums {
		if ret < n {
			ret = n
		}
	}
	return ret
}

func chooseCandidate(dataSet map[string]bool) string {
	minHeight := len(dataSet)
	minHeightString := ""

	for s1 := range dataSet {
		height := 0
		hist := map[int]int{}
		for s2 := range dataSet {
			if s1 == s2 {
				continue
			}
			d := distance(s1, s2)
			hist[d]++
			height = max(height, hist[d])
		}
		if height < minHeight {
			minHeight = height
			minHeightString = s1
		}
	}
	return minHeightString
}

func findSecretWord(wordlist []string, master *Master) {
	dataSet := map[string]bool{}
	for _, s := range wordlist {
		dataSet[s] = true
	}
	times := 0
	for len(dataSet) != 0 {
		candidate := chooseCandidate(dataSet)
		times++
		n := master.Guess(candidate)
		if n == len(candidate) {
			// fmt.Println(candidate, times)
			break
		}
		delete(dataSet, candidate)
		for s := range dataSet {
			if distance(candidate, s) != n {
				delete(dataSet, s)
			}
		}
	}
}

func main() {
	findSecretWord([]string{"eykdft", "gjeixr", "eksbjm", "mxqhpk", "tjplhf", "ejgdra", "npkysm", "jsrsid", "cymplm", "vegdgt", "jnhdvb", "jdhlzb", "sgrghh", "jvydne", "laxvnm", "xbcliw", "emnfcw", "pyzdnq", "vzqbuk", "gznrnn", "robxqx", "oadnrt", "kzwyuf", "ahlfab", "zawvdf", "edhumz", "gkgiml", "wqqtla", "csamxn", "bisxbn", "zwxbql", "euzpol", "mckltw", "bbnpsg", "ynqeqw", "uwvqcg", "hegrnc", "rrqhbp", "tpfmlh", "wfgfbe", "tpvftd", "phspjr", "apbhwb", "yjihwh", "zgspss", "pesnwj", "dchpxq", "axduwd", "ropxqf", "gahkbq", "yxudiu", "dsvwry", "ecfkxn", "hmgflc", "fdaowp", "hrixpl", "czkgyp", "mmqfao", "qkkqnz", "lkzaxu", "cngmyn", "nmckcy", "alpcyy", "plcmts", "proitu", "tpzbok", "vixjqn", "suwhab", "dqqkxg", "ynatlx", "wmbjxe", "hynjdf", "xtcavp", "avjjjj", "fmclkd", "ngxcal", "neyvpq", "cwcdhi", "cfanhh", "ruvdsa", "pvzfyx", "hmdmtx", "pepbsy", "tgpnql", "zhuqlj", "tdrsfx", "xxxyle", "zqwazc", "hsukcb", "aqtdvn", "zxbxps", "wziidg", "tsuxvr", "florrj", "rpuorf", "jzckev", "qecnsc", "rrjdyh", "zjtdaw", "dknezk"}, &Master{"cymplm"})
}
