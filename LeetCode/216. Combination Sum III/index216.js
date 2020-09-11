/**
 * @param {number} k
 * @param {number} n
 * @return {number[][]}
 */
var combinationSum3 = function (k, n, cur = []) {
  if (k === 0 || n === 0 || k > n) {
    return []
  }
  let max = 0
  if (cur.length) {
    max = cur[cur.length - 1] + 1
  } else {
    max = 1
  }
  if (k === 1) {
    if (n >= max && n <= 9) {
      return [[...cur, n]]
    } else {
      return false
    }
  }
  let totalResult = []
  for (let i = max; i <= n; i++) {
    let result = combinationSum3(k - 1, n - i, [...cur, i])
    result && (totalResult = totalResult.concat(result))
  }
  return totalResult.length ? totalResult : cur.length ? false : []
}

console.log(combinationSum3(2, 2))
console.log(combinationSum3(3, 2))
console.log(combinationSum3(3, 20))
console.log(combinationSum3(3, 7))
console.log(combinationSum3(3, 9))
