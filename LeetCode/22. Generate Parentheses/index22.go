package main 

import (
	"fmt"
)

func main(){
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	var ans []string 

	recursive(&ans,[]string{"("},1,n)


	return ans
}

func recursive(ans *[]string,curStack []string,depth int,target int) {
	if depth == 2*target -1{
		curStack = append(curStack,")")
		var temp =""
		for _,v := range curStack {
			temp  = temp + v
		}
		*ans = append(*ans,temp)
		return 
	}

	var curStacka = make([]string , target *2)
	var curStackb = make([]string , target *2)
	
	depth++
	if isLegal("left",curStack,target){
		curStacka = append(curStack,"(")
		recursive(ans,curStacka,depth,target)
	}

	if isLegal("right",curStack,target){
		curStackb = append(curStack,")")
		recursive(ans ,curStackb,depth,target)
	}
}

func isLegal(commaType string,curStack []string,target int ) bool{

	if commaType == "left"{
	var ct = 0
		for _,v := range curStack{
			if v == "("{
				ct ++ 
			}
		}
		if ct >= target{
			return false
		} else {
			return true
		}
	}else{
		var lct =0
		var rct=0
		for _,v := range curStack{
			if v == "("{
				lct ++ 
			}else{
				rct++
			}
		}
		if lct > rct {
			return true
		}else{
			return false
		}
	}
}