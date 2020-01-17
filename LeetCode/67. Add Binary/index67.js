/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function (a, b) {

    if (a.length < b.length) {
        let temp
        temp = a
        a = b
        b = temp
    }

    a = a.split('')
    b = b.split('')
    let shouldPlusOne = 0
    let ans = ''
    let cur
    for (let i = 0; i < a.length; i++) {
        if (i < b.length) {
            cur = shouldPlusOne + ~~a[a.length - 1 - i] + ~~b[b.length - 1 - i]
        } else {
            cur = shouldPlusOne + ~~a[a.length - 1 - i]
        }
        if (cur > 1) {
            shouldPlusOne = 1
            ans = (cur === 2 ? '0' : '1') + ans
        } else {
            shouldPlusOne = 0
            ans = cur + ans
        }
    }
    shouldPlusOne && (ans = '1' + ans)
    // 还要去除ans多余的0
    let index = ans.split('').findIndex(el => el === '1')
    ans = ans.slice(index)
    console.log(ans)
    return ans
};

addBinary('1010', '1011')
addBinary('11', '1')
addBinary('1111', '1111')
addBinary('0', '0100')
addBinary('11', '1')
addBinary('1111', '1')