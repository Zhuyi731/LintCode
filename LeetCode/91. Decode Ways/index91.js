// A message containing letters from A-Z is being encoded to numbers using the following mapping:

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.

// Example 1:

// Input: "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
// Example 2:

// Input: "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

/**
 * @param {string} s
 * @return {number}
 */
var numDecodings = function (s) {
    let dp = [s[0] == 0 ? 0 : 1]
    let len = s.length
    for (let i = 1; i < len; i++) {
        let temp = ~~(s[i-1]+s[i])
        if (temp >= 10 && temp < 27 ) {
            if (s[i] == 0) {
                // 如果等于0  则只能与前面的结合
                dp[i] = i > 1 ? dp[i - 2] : 1
            } else {
                dp[i] = dp[i - 1] + (i > 1 ? dp[i - 2] : 1)
            }
        } else {
            if (s[i] == 0 && (s[i - 1] > 2 || s[i - 1] == 0 || i == 0)) {
                console.log(0)
                return 0
            }
            dp[i] = dp[i - 1]
        }
    }
    console.log(dp[len - 1])
    return dp[len - 1]
};
// "102"
// "000"
// "13723784"
// "18294790"
// "20438921"
// "1212323421"
numDecodings("110") //1
numDecodings("102") //32
numDecodings("17") //2
numDecodings("100") //32
numDecodings("1212323421") //32
numDecodings("20438921") //2
numDecodings("18294790") //0 
numDecodings("13723784")
numDecodings("000")
numDecodings("102")
numDecodings("12")
numDecodings("226")