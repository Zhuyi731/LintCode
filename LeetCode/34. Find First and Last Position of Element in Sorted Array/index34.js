/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var searchRange = function(nums, target) {
    // 首先 二分法查找出最左边的解  
    let resultL= -1,resultR = -1 
    let l=0,r=nums.length,mid = Math.floor((l+r)/2) 
    while(1){
       
        if(target > nums[mid]){
            l = mid 
            mid =  Math.floor((l+r)/2)
        }else if(target > nums[mid]){
            r=mid 
            mid =  Math.floor((l+r)/2)
        }else{
            if(mid=0 || nums[mid]!=target){
                //说明找到最左边的了  
                resultL = mid  
                resultR = mid 
                break;
            }else{
                r=mid 
                mid =  Math.floor((l+r)/2)
            }
        }
    }

    // 然后找最右边的   
    return [l,r]
};

console.log(searchRange([5,7,7,8,8,10],8))
console.log(searchRange([5,7,7,8,8,10],6))

// Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

// Your algorithm's runtime complexity must be in the order of O(log n).

// If the target is not found in the array, return [-1, -1].

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]