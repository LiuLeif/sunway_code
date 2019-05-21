// 2018-03-30 11:23
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Tweet struct {
	id   int
	tick int
}

type MinHeap []Tweet

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].tick > h[j].tick
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (ph *MinHeap) Pop() interface{} {
	t := (*ph)[len(*ph)-1]
	*ph = (*ph)[:len(*ph)-1]
	return t
}

func (ph *MinHeap) Push(x interface{}) {
	*ph = append(*ph, x.(Tweet))
}

type Twitter struct {
	posts   map[int][]Tweet
	follows map[int]map[int]bool
	tick    int
}

func Constructor() Twitter {
	ret := Twitter{}
	ret.posts = map[int][]Tweet{}
	ret.follows = map[int]map[int]bool{}
	ret.tick = 0
	return ret
}

func (twitter *Twitter) PostTweet(userId int, tweetId int) {

	if twitter.posts[userId] == nil {
		twitter.posts[userId] = []Tweet{}
	}
	twitter.tick++
	twitter.posts[userId] = append([]Tweet{Tweet{tweetId, twitter.tick}}, twitter.posts[userId]...)
}

func min(nums ...int) int {
	ret := math.MaxInt32
	for _, n := range nums {
		if ret > n {
			ret = n
		}
	}
	return ret
}

func (twitter *Twitter) GetNewsFeed(userId int) []int {
	queue := MinHeap{}

	followees := twitter.follows[userId]
	for i := 0; i < min(10, len(twitter.posts[userId])); i++ {
		heap.Push(&queue, twitter.posts[userId][i])
	}
	for f, _ := range followees {
		for i := 0; i < min(10, len(twitter.posts[f])); i++ {
			heap.Push(&queue, twitter.posts[f][i])
		}
	}
	ret := []int{}
	for i := 0; i < 10 && len(queue) != 0; i++ {
		ret = append(ret, heap.Pop(&queue).(Tweet).id)
	}
	return ret
}

func (twitter *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	if twitter.follows[followerId] == nil {
		twitter.follows[followerId] = map[int]bool{}
	}
	twitter.follows[followerId][followeeId] = true
}

func (twitter *Twitter) Unfollow(followerId int, followeeId int) {
	delete(twitter.follows[followerId], followeeId)
}

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 10)
	twitter.PostTweet(1, 11)
	twitter.PostTweet(1, 12)
	twitter.PostTweet(1, 13)
	twitter.PostTweet(1, 14)
	twitter.PostTweet(1, 15)
	twitter.PostTweet(1, 16)
	twitter.PostTweet(1, 17)
	twitter.PostTweet(1, 18)
	twitter.PostTweet(1, 19)

	// User 1's news feed should return a list with 1 tweet id -> [5].
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 follows user 2.
	twitter.Follow(1, 2)

	// User 2 posts a new tweet (id = 6).
	twitter.PostTweet(2, 20)

	// User 1's news feed should return a list with 2 tweet ids -> [6, 5].
	// Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	fmt.Println(twitter.GetNewsFeed(1))

	// User 1 unfollows user 2.
	twitter.Unfollow(1, 2)

	// User 1's news feed should return a list with 1 tweet id -> [5],
	// since user 1 is no longer following user 2.
	fmt.Println(twitter.GetNewsFeed(1))

}
