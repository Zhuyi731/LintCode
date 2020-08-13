/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function (coins, amount) {
    let dp = [0]

    for (let i = 1; i <= amount; i++) {
        dp[i] = 0xffffff
        for (let j = 0; j < coins.length; j++) {
            if (i - coins[j] >= 0) {
                dp[i] = Math.min(dp[i - coins[j]] + 1, dp[i])
            }
        }
    }
    return dp[amount] === 0xffffff? -1 :dp[amount]
};