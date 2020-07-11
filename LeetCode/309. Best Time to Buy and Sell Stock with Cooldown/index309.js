/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  // dp[i][0] 表示至第i天持有股票时最大收益
  // dp[i][1] 表示至第i天没有持有股票没有处于冷冻期的最大收益
  // dp[i][2] 表示至第i天没有持有股票且处于冷冻期的最大收益
  let dp = [];
  dp[0] = [-prices[0], 0, 0];
  if (prices.length <= 1) return 0;
  dp[1] = [
    Math.max(-prices[0], -prices[1]),
    0,
    Math.max(prices[1] - prices[0], 0),
  ];

  for (let i = 2; i < prices.length; i++) {
    dp[i] = [];
    dp[i][0] = Math.max(
      dp[i - 1][0],
      dp[i - 2][2] - prices[i],
      dp[i - 1][1] - prices[i]
    );
    dp[i][1] = Math.max(dp[i - 1][0], dp[i - 1][2], dp[i - 1][1]);
    dp[i][2] = dp[i - 1][0] + prices[i];
  }

  let len = prices.length - 1;
  let ans = Math.max(dp[len][0], dp[len][1], dp[len][2]);
  console.log(ans);
  return ans;
};
maxProfit([6, 1, 6, 4, 3, 0, 2]); //7
maxProfit([12, 1, 3, 100, 2]);
maxProfit([1, 2, 3, 0, 2]);
maxProfit([1, 1, 1, 1, 1]);
maxProfit([1, 3, 32, 0, 2]);
maxProfit([12, 1, 3, 0, 2]);

// [6,1,6,4,3,0,2]\n[1, 2, 3, 0, 2]\n[1, 1, 1, 1, 1]\n[1, 3, 32, 0, 2]\n[12, 1, 3, 100, 2]\n[12, 1, 3, 0, 2]
