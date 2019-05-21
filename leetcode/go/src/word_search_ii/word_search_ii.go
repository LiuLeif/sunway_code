// 2017-11-22 18:51
package main

import "fmt"

func dfs(ret *[]string, trie *TrieNode, visited [][]bool, board [][]byte, i, j int) {
	trie = trie.Next[board[i][j]-'a']
	if trie == nil {
		return
	}

	if trie.Word != "" {
		*ret = append(*ret, trie.Word)
		trie.Word = ""
	}

	visited[i][j] = true
	if i > 0 && visited[i-1][j] == false {
		dfs(ret, trie, visited, board, i-1, j)
	}
	if i < len(board)-1 && visited[i+1][j] == false {
		dfs(ret, trie, visited, board, i+1, j)
	}
	if j > 0 && visited[i][j-1] == false {
		dfs(ret, trie, visited, board, i, j-1)
	}
	if j < len(board[0])-1 && visited[i][j+1] == false {
		dfs(ret, trie, visited, board, i, j+1)
	}
	visited[i][j] = false
}

func findWords(board [][]byte, words []string) []string {
	ret := []string{}
	trie := buildTrie(words)
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			dfs(&ret, trie, visited, board, x, y)
		}
	}

	return ret
}

type TrieNode struct {
	Next [26]*TrieNode
	Word string
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{}
	for _, word := range words {
		p := root
		for i := 0; i < len(word); i++ {
			index := word[i] - 'a'
			if p.Next[index] == nil {
				p.Next[index] = &TrieNode{}
			}
			p = p.Next[index]
		}
		p.Word = word
	}
	return root
}

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words := []string{"ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"}
	fmt.Println(findWords(board, words))
}
