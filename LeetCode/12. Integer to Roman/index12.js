/**
 * @param {number} num
 * @return {string}
 */
var intToRoman = function(num) {
    let map = ["I", "V", "X", "L", "C", "D", "M"],
        // "I": "1",
        // "V": "5",
        // "X": "10",
        // "L": "50",
        // "C": "100",
        // "D": "500",
        // "M": "1000"
        ans = "",
        strNum,
        len,
        upper5,
        roman,
        pre;

    while (num > 0) {
        strNum = "" + num;
        upper5 = strNum[0] >= 5;
        len = strNum.length;

        if (upper5) {
            if (strNum[0] == "9") {
                roman = map[len * 2];
                pre = map[len * 2 - 2];
                ans = ans + pre + roman;
                num -= 9 * Math.pow(10, len - 1);
            } else {
                roman = map[len * 2 - 1];
                ans += roman;
                num -= 5 * Math.pow(10, len - 1);
            }
        } else {
            if (strNum[0] == "4") {
                roman = map[len * 2 - 1];
                pre = map[len * 2 - 2];
                ans = ans + pre + roman;
                num -= 4 * Math.pow(10, len - 1);
            } else {
                roman = map[len * 2 - 2];
                ans += rec(~~strNum[0], roman);
                num -= (~~strNum[0]) * Math.pow(10, len - 1);
            }
        }
    }

    console.log(ans);
    return ans;

    function rec(times, str) {
        let ret = "";
        while (times--) {
            ret += str;
        }
        return ret;
    }
};
intToRoman(3); //"III"
intToRoman(4); //"IV"
intToRoman(9); //IX
intToRoman(58); //LVIII
intToRoman(1994); //MCMXCIV

/**
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
 */