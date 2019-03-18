/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function(s, p) {
    let sIndex = 0,
        pIndex = 0,
        ans = false;

    for (let i = 0; i < s.length; i++) {
        if (p[pIndex + 1] == "*") {
            while ((s[i] == p[pIndex] || p[pIndex] == ".") && i < s.length) {
                if (isMatch(s.substring(i + 1, s.length), p.substring(pIndex + 2, p.length))) {
                    return true;
                }
                i++;
            }
            i--;
            pIndex += 2;
        } else {
            if (s[i] == p[pIndex] || p[pIndex] == ".") {
                pIndex++;
            } else {
                return false;
            }
        }
    }

    if (pIndex == p.length) {
        return true;
    } else {
        if (p.substring(pIndex, p.length).replace(/([a-z]\*)/g, "") == "") {
            return true;
        }
        return false;
    }
};
console.log(isMatch("a", "ab*")); //true
console.log(isMatch("aaa", "ab*a*c*a")); //true
console.log(isMatch("abcbcsdcbc", ".*cbc.*cba")); //false
console.log(isMatch("ab", ".*c"));
console.log(isMatch("aa", "a")); //false
console.log(isMatch("aa", "a*")); //true
console.log(isMatch("ab", ".*")); //true
console.log(isMatch("aab", "c*a*b*")); //true
console.log(isMatch("mississippi", "mis*is*p*.")); //false

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