/**
 * @param {number[]} nums
 * @param {number} S
 * @return {number}
 */
var findTargetSumWays = function (nums, S) {
  let sum = [];
  nums.forEach((el, index) => {
    if (index) {
      sum[index] = sum[index - 1] + el;
    } else {
      sum[index] = el;
    }
  });
  dp = [];
  dp[0] = {};
  dp[0][nums[0]] = 1;
  dp[0][-nums[0]] = dp[0][-nums[0]] ? 2 : 1;
  for (let i = 1; i < nums.length; i++) {
    dp[i] = {};
    for (let j = -sum[i]; j <= sum[i]; j++) {
      dp[i][j] = (dp[i - 1][j - nums[i]] || 0) + (dp[i - 1][j + nums[i]] || 0);
    }
  }
  return dp[nums.length - 1][S] || 0;
};

findTargetSumWays([1, 2, 1], 0);
findTargetSumWays([0, 0, 0, 0, 0, 0, 0, 0, 1], 1);
findTargetSumWays([1, 1, 1, 1, 1], 3);
