package  main 

import(
	"fmt"
)

func main(){
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	var dp = new(int) Slice(len(s))
	dp[0] = 
	for index,_ range s{
		if index > 1 
	}

	for index,_ := range s {
		if isStackValid(stack,s[index]){
			stack = append(stack,s[index])
			curLen++  
		}else{
			if curLen > ans {
				ans = curLen
			}
			stack = stack[:0]
		}
	}
	
	return 0
}
