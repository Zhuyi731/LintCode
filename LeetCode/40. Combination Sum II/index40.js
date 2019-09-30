/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum2 = function (candidates, target) {
    let ans = []
    candidates = candidates.sort((a,b)=>a-b)
    dfs([],-1)
    ans = Array.from(new Set(ans.map(el=>el.join(','))))
    ans = ans.map(el=>el.split(','))
    console.log(ans.join(''))
    return ans
    function sum(cur) {
        return cur.reduce((prev, cur) => {
            return prev + cur
        },0)
    }

    function dfs(cur,index) {
        let total = sum(cur)
        if (total > target) {
        } else if (total === target) {
            ans.push(cur)
        } else {
            candidates.forEach((el,i)=>{
               if(index===-1 && i>=1 && candidates[i-1]===candidates[i]){
               }else{
                i>index && dfs(cur.concat([el]),i)
               }
            })
        }
    }

};

combinationSum([2,5,2,1,2],5)
combinationSum([10,1,2,7,6,1,5],8)
console.log(combinationSum([2, 3, 6, 7], 7))
console.log(combinationSum([2, 3, 5], 8))

// Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// Each number in candidates may only be used once in the combination.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [10,1,2,7,6,1,5], target = 8,
// A solution set is:
// [
//   [1, 7],
//   [1, 2, 5],
//   [2, 6],
//   [1, 1, 6]
// ]
// Example 2:

// Input: candidates = [2,5,2,1,2], target = 5,
// A solution set is:
// [
//   [1,2,2],
//   [5]
// ]