/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function (s, p) {
  let dp = [[true]]; // dp[i][j] 表示 s的前i个字符串是否与p的前j个字符串匹配
  for (let i = 0; i <= s.length; i++) {
    dp[i] = [];

    for (let j = 0; j <= p.length; j++) {
      if (i == 0) {
        if (j === 0) {
          dp[i][j] = true;
          continue;
        }

        if (p[j - 1] === "*") {
          // 如果为通配
          dp[i][j] = dp[i][j - 2];
        } else {
          dp[i][j] = false;
        }
      } else {
        if (j == 0) {
          dp[i][j] = false;
          continue;
        }
        if (p[j - 1] === "*") {
          // 如果为通配
          if (p[j - 2] === "." || s[i - 1] === p[j - 2]) {
            dp[i][j] = dp[i - 1][j] || (j >= 2 && dp[i][j - 2]); // 匹配一个 或者不匹配
          } else {
            dp[i][j] = (j >= 2 && dp[i][j - 2]) || false; // 不匹配  或者false
          }
        } else if (p[j - 1] === "." || s[i - 1] === p[j - 1]) {
          dp[i][j] = dp[i - 1][j - 1];
        } else {
          dp[i][j] = false;
        }
      }
    }
  }
  return !!dp[s.length][p.length];
};
console.log(isMatch("aaa", "ab*a*c*a")); //true
console.log(isMatch("a", "ab*")); //true
console.log(isMatch("aa", "a*")); //true
console.log(isMatch("abcbcsdcbc", ".*cbc.*cba")); //false
console.log(isMatch("aab", "c*a*b*")); //true
console.log(isMatch("ab", ".*c")); //false
console.log(isMatch("aa", "a")); //false
console.log(isMatch("ab", ".*")); //true
console.log(isMatch("mississippi", "mis*is*p*.")); //false
console.log(isMatch("bbbba", ".*a*a")); // true

/**
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/
