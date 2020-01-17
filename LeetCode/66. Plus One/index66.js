/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
    let len = digits.length
    for (let i = len - 1; i >= -1; i--) {
        if (i === -1) {
            digits.unshift(1)
            break
        }
        digits[i] = digits[i] + 1
        if (digits[i] === 10) {
            digits[i] = 0
        } else {
            break;
        }
    }
    console.log(digits)
    return digits
};

plusOne([1, 9, 9])
plusOne([1, 2, 3])
plusOne([9, 9, 9])