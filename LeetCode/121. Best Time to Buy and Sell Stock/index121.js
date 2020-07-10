/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function (prices) {
  let profit = 0;
  let currentMin = prices[0];
  prices.forEach((price) => {
    let curProfit = price - currentMin;
    profit = curProfit > profit ? curProfit : profit;

    if (price < currentMin) {
      currentMin = price;
    }
  });
  return profit
};

maxProfit([7, 1, 5, 3, 6, 4]); // 5
maxProfit([7, 6, 4, 3, 1]); // 5
