/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    if(!s) return s;
    let i, j, ans = s[0],
        len = s.length,
        isPalindrome = []; //isPalindrome[i][j]表示 s[i]~s[j]是否为palindrome  

    //初始化
    for (i = 0; i < len; i++) {
        isPalindrome[i] = [];
        isPalindrome[i][i] = true;
        if (i != len - 1) {
            if (s[i] == s[i + 1]) {
                isPalindrome[i][i + 1] = true;
                ans = s.substring(i, i + 2);
            } else {
                isPalindrome[i][i + 1] = false;
            }
        }
    }

    //遍历查找
    //i 
    for (j = 2; j < len; j++) {
        for (i = 0; i < len - j; i++) {
            if (isPalindrome[i + 1][i + j - 1] && s[i] == s[i + j]) {
                isPalindrome[i][i + j] = true;
                ans = s.substring(i, i + j + 1);
            } else {
                isPalindrome[i][i + j] = false;
            }
        }
    }
    console.log(ans);
    return ans;
};
longestPalindrome("aaaaab");
/**
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
 */