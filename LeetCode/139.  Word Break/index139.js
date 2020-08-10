/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
var wordBreak = function (s, wordDict) {
  let dp = [];
  for (let i = 0; i < s.length; i++) {
    dp[i] = false;
    for (let j = 0; j < wordDict.length; j++) {
      if (dp[i]) break;
      if (i < wordDict[j].length - 1) continue;
      if (
        s.substr(i - wordDict[j].length + 1, wordDict[j].length) === wordDict[j]
      ) {
        if (i - wordDict[j].length === -1) {
          dp[i] = true;
          break;
        }
        dp[i] = dp[i] ? dp[i] : dp[i - wordDict[j].length];
      } else {
        dp[i] = dp[i] ? dp[i] : false;
      }
    }
  }
  console.log(dp[s.length - 1]);
  return dp[s.length - 1];
};

wordBreak("leetcode", ["leet", "code"]);
wordBreak("applepenapple", ["apple", "pen"]);
wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]);
