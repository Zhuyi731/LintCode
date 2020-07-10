/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let profit = 0;
  for (let i = 1; i < prices.length; i++) {
    if (prices[i] - prices[i - 1] > 0) {
      profit += prices[i] - prices[i - 1];
    }
  }
  console.log(profit);
  return profit;
};

maxProfit([7, 1, 5, 3, 6, 4]); // 7
maxProfit([1, 2, 3, 4, 5]); // 4
maxProfit([7, 6, 4, 3, 1]); // 0
