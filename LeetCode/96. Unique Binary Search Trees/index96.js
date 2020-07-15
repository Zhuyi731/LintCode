/**
 * @param {number} n
 * @return {number}
 */
var numTrees = function (n) {
  let dp = [1, 1];
  dp = [1, 1];
  for (let i = 2; i <= n; i++) {
    dp[i] = 0;
    for (let j = 1; j <= i; j++) {
      dp[i] += dp[j - 1] * dp[i - j];
    }
  }
  console.log(dp[n]);
  return dp[n];
};

numTrees(3);
numTrees(4);
numTrees(5);
numTrees(6);
numTrees(7);
