/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
var combinationSum = function (candidates, target) {
    let ans = []
      candidates = candidates.sort()
      dfs([],0)
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
                 i>=index && dfs(cur.concat([el]),i)
              })
          }
      }
  
  };

combinationSum([2,5,2,1,2],5)
combinationSum([10,1,2,7,6,1,5],8)
console.log(combinationSum([2, 3, 6, 7], 7))
console.log(combinationSum([2, 3, 5], 8))

// Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

// The same repeated number may be chosen from candidates unlimited number of times.

// Note:

// All numbers (including target) will be positive integers.
// The solution set must not contain duplicate combinations.
// Example 1:

// Input: candidates = [2,3,6,7], target = 7,
// A solution set is:
// [
//   [7],
//   [2,2,3]
// ]
// Example 2:

// Input: candidates = [2,3,5], target = 8,
// A solution set is:
// [
//   [2,2,2,2],
//   [2,3,3],
//   [3,5]
// ]