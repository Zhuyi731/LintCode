// Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

// For example, given the following triangle

// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

// Note:

// Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.
/**
 * @param {number[][]} triangle
 * @return {number}
 */
var minimumTotal = function (triangle) {
    let dp = [triangle[0][0]]
    for (let i = 1; i < triangle.length; i++) { // i是深度  
        for (let j = triangle[i].length - 1; j >= 0; j--) {
            let cur = triangle[i][j]
            if (j === 0 || j === triangle[i].length -1 ) {
                dp[j] = dp[j===0?j:j-1] + cur
            } else {
                dp[j] = Math.min(dp[j] + cur, dp[j - 1] + cur)
            }
        }
    }
    let ans = dp[0]
    dp.forEach(el => {
        if (el < ans) {
            ans = el
        }
    })
    console.log(ans)
    return ans
};

minimumTotal([
    [2],
    [3, 4],
    [6, 5, 7],
    [4, 1, 8, 3]
])