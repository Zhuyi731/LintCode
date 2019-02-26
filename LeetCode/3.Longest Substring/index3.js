/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    if(s.length ==0 )return 0;
    var len = s.length,
        left = 0,
        right = 1,
        cur = 1,
        max = 1,
        index,
        containingStrs = s.substr(0, 1);

    while (right < len) {
        index = containingStrs.indexOf(s[right]);
        if (index > -1) {
            left = left + index + 1;
            cur = right - left + 1;
            containingStrs = containingStrs.substr(index + 1, containingStrs.length) + s[right];
        } else {
            containingStrs += s[right];
            cur++;
            max = cur > max ? cur : max;
        }
        right++;
    }
    console.log(max);
    return max
};
//TEST:
lengthOfLongestSubstring("ba");
/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3 
Explanation: The answer is "abc", with the length of 3. 
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3. 
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 */