package main

import "fmt"

func main() {
	//[false,false,true,true,false,true,true,false]
	fmt.Println(camelMatch([]string{"uAxaqlzahfialcezsLfj", "cAqlzyahaslccezssLfj", "AqlezahjarflcezshLfj", "AqlzofahaplcejuzsLfj", "tAqlzahavslcezsLwzfj", "AqlzahalcerrzsLpfonj", "AqlzahalceaczdsosLfj", "eAqlzbxahalcezelsLfj"}, "AqlzahalcezsLfj")) //[true,true,true,true,true]
	fmt.Println(camelMatch([]string{"aksvbjLiknuTzqon", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"}, "ksvjLiknTzqn"))                                                                            //[true,true,true,true,true]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FT"))                                                                                                                                          //[false,false,false,false,false]
	fmt.Println(camelMatch([]string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}, "CooP"))                                                                                                                                                   //[false,false,true]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBaT"))                                                                                                                                       //[false,true,false,false,false]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB"))                                                                                                                                          //[true,false,true,true,false]
	fmt.Println(camelMatch([]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FoBa"))                                                                                                                                        //[true,false,true,false,false]
}

func isUpper(r rune) bool {
	return int(r) >= 65 && int(r) <= 90
}

func camelMatch(queries []string, pattern string) []bool {
	spliters := []string{}
	curIndex := -1
	for _, r := range pattern {
		if int(r) >= 65 && int(r) <= 90 {
			curIndex++
			spliters = append(spliters, string(r))
		} else {
			if curIndex == -1 {
				curIndex = 0
				spliters = append(spliters, string(r))
			} else {
				spliters[curIndex] += string(r)
			}
		}
	}

	ans := make([]bool, len(queries))

	for i, v := range queries {
		currentMatchIndex := -1
		mapIndex := 0
		ans[i] = true
		for _, r := range v {
			if isUpper(r) {
				if currentMatchIndex >= 0 && mapIndex < len(spliters[currentMatchIndex]) { // 上一个没有匹配到
					ans[i] = false
					break
				}
				currentMatchIndex++
				mapIndex = 0
				if currentMatchIndex >= len(spliters) || r != rune(spliters[currentMatchIndex][mapIndex]) {
					ans[i] = false
					break
				}
				mapIndex++
			} else { // 小写
				if currentMatchIndex == -1 {
					if isUpper(rune(spliters[0][0])) {
						continue
					}
					currentMatchIndex = 0
				}
				if mapIndex < len(spliters[currentMatchIndex]) && r == rune(spliters[currentMatchIndex][mapIndex]) {
					mapIndex++
				}
			}
		}
		if currentMatchIndex < len(spliters)-1 {
			ans[i] = false
		}
	}

	return ans
}
