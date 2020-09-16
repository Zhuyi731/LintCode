/**
 * @param {string[]} strs
 * @param {number} m
 * @param {number} n
 * @return {number}
 */

function ct01(str) {
  let ct0 = 0
  let ct1 = 0
  for (let i = 0; i < str.length; i++) {
    if (str[i] === "0") {
      ct0++
    } else {
      ct1++
    }
  }
  return [ct0, ct1]
}

var findMaxForm = function (strs, m, n) {
  let dp = []
  for (let i = 0; i <= n; i++) {
    dp[i] = []
    for (let j = 0; j <= m; j++) {
      dp[i][j] = []
      for (let k = 0; k < strs.length; k++) {
        let [ct0, ct1] = ct01(strs[k])
        if (i >= ct1 && j >= ct0) {
          dp[i][j][k] = k
            ? Math.max(dp[i][j][k - 1], dp[i - ct1][j - ct0][k-1] + 1)
            : 1
        } else {
          dp[i][j][k] = k ? dp[i][j][k - 1] : 0
        }
      }
    }
  }
  console.log(dp[n][m][strs.length - 1])
  return dp[n][m][strs.length - 1]
}

findMaxForm(["10", "0001", "111001", "1", "0"], 5, 3)

findMaxForm(["10", "1", "0"], 1, 1)
