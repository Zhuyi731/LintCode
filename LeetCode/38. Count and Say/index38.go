package main

import(
	"fmt"
	"bytes"
	"strconv"
)

func main(){
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
	fmt.Println(countAndSay(5))
	fmt.Println(countAndSay(6))
}

func countAndSay(n int) string {
	var ans []string = make([]string,n+1)
	ans[0] = "1"
	ans[1] = "11"
	if n<2{
		return ans[n-1]
	}
	var temp bytes.Buffer
	var ct = 0
	var strlen int
	for i:=1;i<n;i++{
		strlen = len(ans[i])
		ct = 1
		temp =  bytes.Buffer{}
		for j:=1;j<strlen;j++{	
			if ans[i][j]==ans[i][j-1]{
				ct ++ 
			}else{
				temp.WriteString(strconv.Itoa(ct))
				temp.WriteString(string(ans[i][j-1]))
				ct = 1 
			}
		}
		if ans[i][strlen-1] == ans[i][strlen-2]{
			temp.WriteString(strconv.Itoa(ct))
				temp.WriteString(string(ans[i][strlen-1]))
		}else{
			temp.WriteString("1")
			temp.WriteString(string(ans[i][strlen-1]))
		}
		ans[i+1] = temp.String()		
	}
	return ans[n-1]
}


// The count-and-say sequence is the sequence of integers with the first five terms as following:

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221
// 1 is read off as "one 1" or 11.
// 11 is read off as "two 1s" or 21.
// 21 is read off as "one 2, then one 1" or 1211.

// Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

// Note: Each term of the sequence of integers will be represented as a string.