/**
 * @param {string[]} strs
 * @return {string}
 */

var longestCommonPrefix = function(strs) {
    let ans = "",
        cur,
        i = 0,
        j = 0,
        len = strs.length;
    if (strs.length == 0) return "";
    if (strs.length == 1) return strs[0];
    while (true) {
        if (strs[0] == "") {
            return ans;
        }

        cur = strs[0][i];
        for (j = 1; j < len; j++) {
            if (i >= strs[j].length) {
                console.log(ans);
                return ans;
            }
            if (strs[j][i] !== cur) {
                console.log(ans);
                return ans;
            }
        }
        ans += cur;
        i++;
    }
    console.log(ans);
    return ans;
};

longestCommonPrefix(["flower", "flow", "flight"]);
longestCommonPrefix(["dog", "racecar", "car"]);

/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

 */