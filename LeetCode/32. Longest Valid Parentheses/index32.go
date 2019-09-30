package  main 

import(
	"fmt"
)

func main(){
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	var stack []byte 
	var ans int = 0  
	var curLen int = 0
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

func isStackValid(stack []byte,cur byte) bool{
	var l,r int 
	l=0
	r=0

	var LL string = "("
	var byteL byte = LL[0]

	for _,v:=range stack{
		if v == byteL{
			l ++
		}else{
			r++
		}
	}
	if l == r  {
		if cur == byteL{
			return true
		}else{
			return false
		}
	}
	return l < r
}