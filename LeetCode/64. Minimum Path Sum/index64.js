// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

// Note: You can only move either down or right at any point in time.

// Example:

// Input:
// [
//   [1,3,1],
//   [1,5,1],
//   [4,2,1]
// ]
// Output: 7
// Explanation: Because the path 1→3→1→1→1 minimizes the sum.

/**
 * @param {number[][]} grid
 * @return {number}
 */
var minPathSum = function (grid) {
    let dp = []
    let n = grid.length
    let m = grid[0].length
    let i
    let j
    for (i = 0; i < n; i++) {
        dp[i] = []
        for (j = 0; j < m; j++) {
            if(i === 0 && j===0) {
                dp[i][j] = grid[i][j] 
                continue
            }
            if(i === 0) {
                dp[i][j] = dp[i][j-1] + grid[i][j]
                continue
            }
            if(j === 0 ){
                dp[i][j] = dp[i-1][j] + grid[i][j]
                continue
            }

            if(dp[i-1][j] + grid[i][j] >  dp[i][j-1] + grid[i][j]){
                dp[i][j] = dp[i][j-1] + grid[i][j]
            }else{
                dp[i][j] = dp[i-1][j] + grid[i][j]
            }
        }
    }
    console.log(dp[n-1][m-1])
    return dp[n-1][m-1]
};

minPathSum([
    [1,3,1],
    [1,5,1],
    [4,2,1]
])

minPathSum([
    [1,3]
])
minPathSum([
    [1],
    [1],
    [4]
])
minPathSum([
    [1]
])