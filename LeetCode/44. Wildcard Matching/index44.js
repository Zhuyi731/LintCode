/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
//  ? 匹配一个 * 匹配任意个
var isMatch = function (s, p) {
  let dp = [];

  for (let i = 0; i <= s.length; i++) {
    dp[i] = [];
    if (i == 0) dp[i] = [true];
    for (let j = 1; j <= p.length; j++) {
      if (i === 0) {
        //空字符串做特殊处理
        if (p[j - 1] === "*") {
          // 如果是* 则匹配s的一个，并且可以继续匹配
          dp[i][j] = dp[i][j - 1];
        } else {
          dp[i][j] = false; // 否则不匹配
        }
      } else {
        // 如果s不为空字符串
        if (p[j - 1] === "?" || p[j - 1] === s[i - 1]) {
          // 如果匹配
          dp[i][j] = dp[i - 1][j - 1];
        } else if (p[j - 1] === "*") {
          dp[i][j] = dp[i - 1][j] || dp[i][j - 1];
        } else {
          dp[i][j] = false;
        }
      }
    }
  }
  console.log(dp[s.length][p.length] || false);
  return dp[s.length][p.length] || false;
};
isMatch("aa", "a"); //false
isMatch("acdcb", "ac*?b"); // true
isMatch("aa", "*"); //true
isMatch("cb", "?a"); //false
isMatch("acdcb", "a*c?b"); // false
