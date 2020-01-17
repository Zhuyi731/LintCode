// Given a collection of distinct integers, return all possible permutations.

// Example:

// Input: [1,2,3]
// Output:
// [
//   [1,2,3],
//   [1,3,2],
//   [2,1,3],
//   [2,3,1],
//   [3,1,2],
//   [3,2,1]
// ]

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/permutations
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permute = function (nums) {
    let ans = []

    getPermute([], nums.map(el => false))
    function getPermute(cur, curUsed) {
        let ct = 0
        curUsed.forEach(el => {
            if (!el) {
                ct++
            }
        })
        let isEnd = ct === 1

        curUsed.forEach((el, index) => {
            if (!el) {
                let tpCur = cur.map(el => el)
                let tpUsed = curUsed.map(el => el)
                tpUsed[index] = true
                tpCur.push(nums[index])
                if (isEnd) {
                    ans.push(tpCur)
                } else {
                    getPermute(tpCur, tpUsed)
                }
            }
        })


    }
    console.log(...ans)
    console.log('')
    return ans
};

permute([1, 2])
permute([1, 2, 3])
permute([1, 1, 3])