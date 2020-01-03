// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:

// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Follow up:

// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function (nums) {
    let dp = [nums[0]]
    let ans = nums[0]
    // dp[i] 以i结尾的最大和
    for (let i = 1; i < nums.length; i++) {
        dp[i] = dp[i - 1] > 0 ? (nums[i] + dp[i - 1]) : nums[i]
        ans = ans > dp[i] ? ans : dp[i]
    }
    console.log(ans)
    return ans
};

maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4])

var maxSubArray2 = function (nums) {
    let curMax = nums[0]
    let sum = nums[0]
    // dp[i] 以i结尾的最大和
    for (let i = 1; i < nums.length; i++) {
        if(sum < 0){
            sum = nums[i]
        }else{
            sum = sum + nums[i] 
        }
        curMax = curMax > sum ? curMax : sum
    }
    console.log(curMax)
    return curMax
};

maxSubArray2([-2, 1, -3, 4, -1, 2, 1, -5, 4])
