/**
 * 描述
 * 计算数字k在0到n中的出现的次数，k可能是0~9的一个值
 * 
 * 样例
 * 例如n=12，k=1，在 [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]，我们发现1出现了5次 (1, 10, 11, 12)
 */


/**
 * @param k: An integer
 * @param n: An integer
 * @return: An integer denote the count of digit k in 1..n
 */
const digitCounts = function (k, n) {
    let len,
        i,
        countArr = [0],
        ct = 0;
    n = n.toString();
    len = n.length;

    //计算arr数组
    countCountArr(len);

    for (i = 0; i < len; i++) {
        ct = ct + getCt(n,n[i], len - i);
    }
    return ct;


    function getCt(str,number, index) {
        if (number > k) {
            return countArr[index];
        } else if (number < k) {
            return countArr[index - 1];
        } else {
            return parseInt(str.slice(len - k)) + 1;
        }
    }

    function countCountArr(len) {
        for (i = 0; i < len; i++) {
            countArr[i + 1] = Math.pow(10, i);
        }
    }


}

console.log(digitCounts(1,12));
