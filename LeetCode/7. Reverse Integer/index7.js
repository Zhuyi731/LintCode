/**
 * @param {number} x
 * @return {number}
 */
var reverse = function(x) {
    let MIN = 0 - Math.pow(2, 31),
        max = Math.pow(2, 31) - 1,
        ans;

    if (x == 0) return 0;

    let prefix = x < 0 ? "-" : "",
        stringNum = Number.prototype.toString.call(Math.abs(x));

    stringNum = stringNum.split("").reverse().join("");
    let firstNotZero = 0;
    while (firstNotZero < stringNum.length) {
        if (stringNum[firstNotZero] === "0") {
            firstNotZero++;
        } else {
            break;
        }
    }

    stringNum = stringNum.substr(firstNotZero);
    ans = prefix + stringNum;

    if (ans < MIN || ans > max) {
        return 0
    }

    return ans;
};

reverse(123);
reverse(-123);
reverse(120);
reverse(0);
reverse(-120200);
/**
 * 
 * Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

 */