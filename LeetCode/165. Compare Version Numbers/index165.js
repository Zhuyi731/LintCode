
/**
 * @param {string} version1
 * @param {string} version2
 * @return {number}
 */
 var compareVersion = function (version1, version2) {
    let r1 = r2 = 0
    let l1 = l2 = 0
    let n1 = n2 = 0
    while (l1 < version1.length || l2 < version2.length) {
        n1 = 0
        n2 = 0
        while (l1 < version1.length && r1 < version1.length && version1[r1] !== ".") {
            r1++
        }
        n1 = ~~version1.slice(l1, r1)
        l1 = r1 + 1
        r1++
        while (l2 < version2.length && r2 < version2.length && version2[r2] !== ".") {
            r2++
        }
        n2 = ~~version2.slice(l2, r2)
        l2 = r2 + 1
        r2++
        if (n1 !== n2) {
            return n1 > n2 ? 1 : -1
        }
    }

    return 0
};

console.log(compareVersion("0.1", "1.1"))       // -1
console.log(compareVersion("1.0.1", "1"))       // 1
console.log(compareVersion("7.5.2.4", "7.5.3")) //-1
console.log(compareVersion(".1", "0.001"))      //0
console.log(compareVersion("1.01", "1.001"))    //0
console.log(compareVersion("1.0", "1.0.0"))     //0
