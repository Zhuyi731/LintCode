package main

import "fmt"

func main() {
	actions(
		[]string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"},
		[][]int{{}, {1}, {}, {3}, {}, {7}, {}, {2}, {}, {6}, {}},
		true,
	)
	fmt.Println("/*********************/")
	actions(
		[]string{"SummaryRanges", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals", "addNum", "getIntervals"},
		[][]int{{}, {6}, {}, {6}, {}, {0}, {}, {4}, {}, {8}, {}, {7}, {}, {6}, {}, {4}, {}, {7}, {}, {5}, {}},
		true,
	)
}

func actions(actions []string, param [][]int, act bool) {
	if !act {
		return
	}
	var a SummaryRanges
	for i, v := range actions {
		switch v {
		case "SummaryRanges":
			a = Constructor()
		case "addNum":
			a.AddNum(param[i][0])
		case "getIntervals":
			a.GetIntervals()
		}
	}
}

type SummaryRanges struct {
	Ranges [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (this *SummaryRanges) AddNum(val int) {
	if len(this.Ranges) == 0 {
		this.Ranges = [][]int{{val, val}}
		return
	}
	for i, v := range this.Ranges {
		if v[0] == val+1 {
			this.Ranges[i][0] = val
			// 合并
			this.merge(i, -1) //向前合并
			return
		} else if v[1] == val-1 {
			this.Ranges[i][1] = val
			this.merge(i, 1)
			return
		} else if val >= v[0] && val <= v[1] {
			return
		} else if v[0] > val {
			// 说明应该插入当前位置
			this.Ranges = append(this.Ranges[:i], append([][]int{{val, val}}, this.Ranges[i:]...)...)
			return
		}
	}
	this.Ranges = append(this.Ranges, []int{val, val})
}

func (this *SummaryRanges) merge(i, direction int) {
	if i+direction >= 0 && i+direction < len(this.Ranges) {
		// 检查
		if direction == -1 {
			if this.Ranges[i-1][1] == this.Ranges[i][0]-1 {
				this.Ranges[i-1][1] = this.Ranges[i][1]
				// 然后删除i下标
				this.deleteIndex(i)
			}
		} else {
			if this.Ranges[i][1] == this.Ranges[i+1][0]-1 {
				this.Ranges[i][1] = this.Ranges[i+1][1]
				this.deleteIndex(i + 1)
			}
		}

	}
}

func (this *SummaryRanges) deleteIndex(i int) {
	this.Ranges = append(this.Ranges[:i], this.Ranges[i+1:]...)
}

func (this *SummaryRanges) GetIntervals() [][]int {
	fmt.Println(this.Ranges)
	return this.Ranges
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
