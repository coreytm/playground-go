package daily

import "sort"

// This problem was asked by Snapchat.
// Given an array of time intervals (start, end) for classroom lectures (possibly overlapping), find the minimum number of rooms required.
// For example, given [(30, 75), (0, 50), (60, 150)], you should return 2.

type (
	timeType   int
	singleTime struct {
		tt timeType
		t  int64
	}
	sortByTime []singleTime
)

const (
	startTime = timeType(0)
	endTime   = timeType(1)
)

func (s sortByTime) Len() int {
	return len(s)
}

func (s sortByTime) Less(i, j int) bool {
	return s[i].t < s[j].t
}

func (s sortByTime) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// minRooms accepts a list of time intervals and returns the min # of rooms necessary (i.e. # overlaps).
// runtime: O(nlog(n)) because sorting uses a version of quick sort
// space:   O(n) because ues of sortByTime which is another array of size 2n
func minRooms(in [][]int64) int {
	ts := sortByTime{}
	for i := 0; i < len(in); i++ {
		st := singleTime{
			tt: startTime,
			t:  in[i][0],
		}
		et := singleTime{
			tt: endTime,
			t:  in[i][1],
		}

		ts = append(ts, st, et)
	}

	sort.Sort(ts)

	max := 0
	cnt := 0

	for _, t := range ts {
		if t.tt == startTime {
			cnt++
			if cnt > max {
				max = cnt
			}
			continue
		}
		cnt--
	}

	return max
}
