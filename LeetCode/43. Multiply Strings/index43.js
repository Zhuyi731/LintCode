/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var multiply = function (num1, num2) {
    // if (num1.length + num2.length < 31) {
    //     return Number(num1) * Number(num2)
    // }
    num1 = num1.split('').reverse()
    num2 = num2.split('').reverse()
    let multiList = []
    let ans
    for (let i = 0; i < num1.length; i++) {
        for (let j = 0; j < num2.length; j++) {
            if (!multiList[i + j]) multiList[i + j] = 0
            multiList[i + j] += ~~num1[i] * ~~num2[j]
        }
    }

    // let len = multiList.length
    for (let i = 0; i < multiList.length; i++) {
        if (multiList[i] >= 10) {
            if (!multiList[i + 1]) multiList[i + 1] = 0
            multiList[i + 1] += Math.floor(multiList[i] / 10)
            multiList[i] = multiList[i] % 10
        }
    }
    for (let i = multiList.length - 1; i > 0; i--) {
        if (multiList[i] === 0) {
            multiList.splice(i, 1)
        } else {
            break;
        }
    }

    ans = multiList.reverse().join('')
    console.log(ans)
    return ans
};

multiply('0', '9345')
multiply('0', '0')
multiply('9', '9')
multiply('123', '456')
multiply('2', '3')





// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

// Example 1:

// Input: num1 = "2", num2 = "3"
// Output: "6"
// Example 2:

// Input: num1 = "123", num2 = "456"
// Output: "56088"
// Note:

// The length of both num1 and num2 is < 110.
// Both num1 and num2 contain only digits 0-9.
// Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// You must not use any built-in BigInteger library or convert the inputs to integer directly.

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/multiply-strings
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。