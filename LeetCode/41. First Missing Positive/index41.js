/**
 * @param {number[]} nums
 * @return {number}
 */
var firstMissingPositive = function (nums) {
    let sorted = []
    let len = nums.length
    for(var i=0;i<=len;i++){
        let el = nums[i]
        if (el > 0 && el <= len) {
            sorted[el] = 1
        }
    }
    let ans = len+1
    for(var i=1;i<=len;i++){
        if(sorted[i] !== 1){
            ans = i 
            break
        }
    }
    console.log(ans)
    return ans
};
firstMissingPositive([1])
firstMissingPositive([1, 2, 0])
firstMissingPositive([3, 4, -1, 1])
firstMissingPositive([7, 8, 9, 11, 12])

// Given an unsorted integer array, find the smallest missing positive integer.

// Example 1:

// Input: [1,2,0]
// Output: 3
// Example 2:

// Input: [3,4,-1,1]
// Output: 2
// Example 3:

// Input: [7,8,9,11,12]
// Output: 1
// Note:

// Your algorithm should run in O(n) time and uses constant extra space.