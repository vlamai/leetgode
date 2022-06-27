package _7_Insert_Interval

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	s := newInterval[0]
	e := newInterval[1]
	res := make([][]int, 0)
	for _, interval := range intervals {
		cs := interval[0]
		ce := interval[1]
		if s >= cs && s <= ce {
			fmt.Printf("start %v\n", interval)
			fmt.Println(cs, s)
		} else if e <= ce && e >= cs {
			fmt.Printf("stop  %v\n", interval)
			fmt.Println(ce, e)
		} else if cs >= e {
			fmt.Println(cs, e)
			fmt.Printf("bigger%v\n", interval)
			res = append(res, interval)
		} else {
			res = append(res, interval)
		}
	}
	return res
}
