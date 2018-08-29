package snapchat

import (
	"fmt"
	"sort"
)

type Video struct {
	id       string
	watchCnt int
}

type User struct {
	id      string
	name    string
	friends []*User
	videos  []*Video
}

type sortByWatchCnt []*Video

func (s sortByWatchCnt) Less(a, b int) bool {
	return s[a].watchCnt > s[b].watchCnt
}

func (s sortByWatchCnt) Len() int {
	return len(s)
}

func (s sortByWatchCnt) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

// getMostViewedVideos returns the videos which the user's friends have watched the most of.
// assume d = max depth of 10
func getMostViewedVideos(user *User, d int) ([]*Video, error) {
	if user == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	if d > 10 || d < 1 {
		return nil, fmt.Errorf("dpeth must be between 1 and 10")
	}

	stack := []*User{}
	cur := user
	stack = append(stack, cur)

	videoCounts := make(map[string]int)
	visitedUsers := make(map[string]bool)
	visitedUsers[cur.id] = true

	depth := 1
	for len(stack) > 0 || depth < d {
		// pop the stack
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// append users friends to the stack
		stack = append(stack, cur.friends...)

		// skip the user if already visited
		if visitedUsers[cur.id] {
			continue
		}

		// add friends watched videos to counter
		for _, video := range cur.videos {
			if _, ok := videoCounts[video.id]; ok {
				videoCounts[video.id]++
			} else {
				videoCounts[video.id] = 1
			}
		}

		// set the user to visited
		visitedUsers[cur.id] = true
		depth++
	}

	videoList := sortByWatchCnt{}
	for id, cnt := range videoCounts {
		videoList = append(videoList, &Video{id: id, watchCnt: cnt})
	}

	sort.Sort(videoList)

	return []*Video(videoList), nil
}
