/**
 * @param {number} m
 * @param {number} n
 * @return {number}
 */
// var uniquePaths = function(m, n) {
//     let i, j
//     let ans = [[1]]
//     for(i=0;i<m;i++){
//         ans[i] = new Array()
//         for(j=0;j<n;j++){
//            if(i || j){
//                ans[i][j] = 0
//                ans[i][j] += j?ans[i][j-1]:0
//                ans[i][j] += i?ans[i-1][j]:0
//            } else{
//                ans[i][j]=1
//            }
//         }
//     }
//     console.log(ans[i-1][j-1])
//     return ans[i-1][j-1]
// };

var uniquePaths = function (n, m) {
    if (n === 1 || m === 1) {
        return 1
    }
    let total = n + m - 2
    let ans = 1
    let i
    for (i = 1; i < Math.min(m, n); i++) {
        ans *= (total - i + 1)
        ans /= i
    }
    return ans
}

console.log(uniquePaths(3, 2))
console.log(uniquePaths(7, 3))
console.log(uniquePaths(1, 3))
console.log(uniquePaths(2, 1))


// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

// How many possible unique paths are there?


// Above is a 7 x 3 grid. How many possible unique paths are there?

// Note: m and n will be at most 100.

// Example 1:

// Input: m = 3, n = 2
// Output: 3
// Explanation:
// From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Right -> Down
// 2. Right -> Down -> Right
// 3. Down -> Right -> Right
// Example 2:

// Input: m = 7, n = 3
// Output: 28