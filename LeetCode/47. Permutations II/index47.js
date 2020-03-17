// Given a collection of numbers that might contain duplicates, return all possible unique permutations.

// Example:

// Input: [1,1,2]
// Output:
// [
//   [1,1,2],
//   [1,2,1],
//   [2,1,1]
// ]

/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var permuteUnique = function (nums) {
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
                    if(index && nums[index] === nums[index-1] && !(tpUsed[index-1] && tpUsed[index])) return 
                    getPermute(tpCur, tpUsed)
                }
            }
        })
    }
    console.log(...ans)
    console.log('')
    return ans
};


// DEBUG: 

permuteUnique([1, 1, 2])
permuteUnique([1, 3, 3, 6])