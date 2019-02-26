/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function(s, numRows) {
    let len = s.length,
        interval = 2 * numRows - 2,
        depth, cur,
        ans = "";
    for (depth = 0; depth < numRows; depth++) {
        let ct = 0;
        cur = interval * ct + depth;
        while (cur < len) {
            ans += s[cur];
            let center = cur + 2 * (numRows - depth - 1);
            if (!(depth == 0 || depth == numRows - 1) && center < len) {
                ans += s[center];
            }
            ct++;
            cur = interval * ct + depth;
        }
    }
    return ans;
    console.log(ans);
};
convert("PAYPALISHIRING", 4);

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/