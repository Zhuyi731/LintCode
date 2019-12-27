// 63. Unique Paths II
// Medium

// 1178

// 192

// Add to List

// Share
// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

// The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

// Now consider if some obstacles are added to the grids. How many unique paths would there be?



// An obstacle and empty space is marked as 1 and 0 respectively in the grid.

// Note: m and n will be at most 100.

// Example 1:

// Input:
// [
//   [0,0,0],
//   [0,1,0],
//   [0,0,0]
// ]
// Output: 2
// Explanation:
// There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right

/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */
var uniquePathsWithObstacles = function (obstacleGrid) {
    let n = obstacleGrid.length
    let m = obstacleGrid[0].length

    let i, j
    let ans = []
    for (i = 0; i < n; i++) {
        ans[i] = []
        for (j = 0; j < m; j++) {
            if (!i && !j) {
                ans[i][j] = obstacleGrid[i][j] ? 0 : 1
                continue
            }
            ans[i][j] = 0
            if (obstacleGrid[i][j]) continue
            if (i && obstacleGrid[i - 1][j] !== 1) {//有障碍物
                ans[i][j] += ans[i - 1][j]
            }
            if (j && obstacleGrid[i][j - 1] !== 1) {
                ans[i][j] += ans[i][j - 1] 
            }
        }
    }
    return ans[i - 1][j - 1]
};
console.log(
    uniquePathsWithObstacles(
        [
            [0, 1]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [1]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 0, 0],
            [0, 1, 0],
            [0, 0, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 0, 0],
            [0, 1, 0],
            [0, 1, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 0, 0],
            [0, 1, 0],
            [0, 0, 0],
            [0, 1, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 0, 0],
            [0, 1, 0],
            [0, 0, 0],
            [1, 0, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 1, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0, 0, 0]
        ]
    )
)
console.log(
    uniquePathsWithObstacles(
        [
            [0], [0], [0]
        ]
    )
)
