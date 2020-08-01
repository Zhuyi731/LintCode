/**
 * @param {number} N
 * @return {boolean}
 */
var divisorGame = function (N) {
  let dp = [false, false, true, false];
  for (let i = 4; i <= N; i++) {
    for (let j = 1; j < i / 2; j++) {
      if (i % j === 0 && dp[i - j] === false) {
        dp[i] = true;
        break;
      }
    }
    if (!dp[i]) {
      dp[i] = false;
    }
  }

  console.log(dp[N]);
  return dp[N];
};

divisorGame(3);
