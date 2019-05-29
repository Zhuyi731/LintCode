package main 

import (
	"fmt"
	"strings"
)

func main (){
	fmt.Println(letterCombinations("999"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("7"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("133334"))
	fmt.Println(letterCombinations("3333"))
	fmt.Println(letterCombinations("4396"))
}

var keyMap map[string]string = map[string]string {
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",
}

func letterCombinations(digits string) []string {
	for _,v := range digits{
		if string(v) == "1"{
			return []string{}
		}
	}
	if len(digits) == 0 {
		return []string{}
	}
	return getSlice(digits)
}

func getSlice(input string) []string{
	if len(input) == 1{
		return strings.Split(keyMap[string(input[0])],"")
	}

	
	nextSlice := getSlice(input[1:])
	
	var s0,s1,s2,s3 []string
	s0 =  make([]string, len(nextSlice))
	s1 =  make([]string, len(nextSlice))
	s2 =  make([]string, len(nextSlice))
	if len(keyMap[string(input[0])]) == 4{
		s3 =  make([]string, len(nextSlice))
	}else{
		s3 = []string{}
	}
	for i,_ := range nextSlice {
		s0[i] = string(keyMap[string(input[0])][0]) + nextSlice[i]
		s1[i] = string(keyMap[string(input[0])][1]) + nextSlice[i]
		s2[i] = string(keyMap[string(input[0])][2]) + nextSlice[i]
		if len(keyMap[string(input[0])]) == 4{
			s3[i] =  string(keyMap[string(input[0])][3]) + nextSlice[i]
		}
	}
	ret := append(s0,append(s1,append(s2,s3...)...)...)
	fmt.Println(ret)
	return ret
}

/**
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/