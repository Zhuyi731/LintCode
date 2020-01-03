// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

// Example 1:

// Input: "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()"
// Example 2:

// Input: ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()"
/**
 * @param {string} s
 * @return {number}
 */
var longestValidParentheses = function (s) {
    let dp = [0] //dp[i] 表示以i结尾的最大长度  
    let ans = 0
    for (let i = 1; i < s.length; i++) {
        if (s[i] === ')') {
            if (s[i - 1] === '(') {
                dp[i] = (i > 1 ? dp[i - 2] : 0) + 2
            } else if (i - dp[i - 1] > 0 && s[i - dp[i - 1] - 1] === '(') { // s[i-1] === ')'  namei-1  
                dp[i] = (i - dp[i - 1] > 2 ? dp[i - dp[i - 1] - 2] : 0) + dp[i - 1] + 2
            } else {
                dp[i] = 0
            }
        } else {
            dp[i] = 0
        }
        ans = ans > dp[i] ? ans : dp[i]
    }
    console.log(ans)
};

longestValidParentheses("(")
longestValidParentheses(")")
longestValidParentheses(")(")
longestValidParentheses("()")
longestValidParentheses("(()())")
longestValidParentheses(")()())")
longestValidParentheses("(()")