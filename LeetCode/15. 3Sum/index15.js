/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    nums.sort((a, b) => a > b ? 1 : -1);

    let i, j, k,
        len = nums.length,
        ans = [];

    for (i = 0; i < len - 2; i++) {
        if (nums[i] + nums[i + 1] > 0) {
            break;
        }
        if (i && nums[i] == nums[i - 1]) {
            continue;
        }
        for (j = i + 1; j < len - 1; j++) {
            if (nums[i] + nums[j] + nums[j + 1] > 0) {
                break;
            }
            if (nums[j] == nums[j - 1] && j != i + 1) {
                continue;
            }
            let res = binarySearch(j, len, 0 - nums[i] - nums[j]);
            if (res != null) {
                ans.push([nums[i], nums[j], res]);
            }
            // for (k = j + 1; k < len; k++) {
            //     if (nums[i] + nums[j] + nums[k] > 0) {
            //         break;
            //     } else if (nums[i] + nums[j] + nums[k] == 0) {
            //         if (!ans.find(el => {
            //                 return el[0] == nums[i] && el[1] == nums[j];
            //             })) {
            //             ans.push([nums[i], nums[j], nums[k]]);
            //         }
            //         break;
            //     }
            // }
        }
    }

    function binarySearch(l, r, target) {
        let mid = Math.floor((l + r) / 2);
        if (l == r) {
            return null;
        }
        if (l == r - 1) {
            if (target == nums[l]) {
                return nums[l];
            } else if (target == nums[r]) {
                return nums[r];
            } else {
                return null;
            }
        }

        if (nums[mid] > target) {
            r = mid;
            return binarySearch(l, r, target);
        } else if (nums[mid] < target) {
            l = mid;
            return binarySearch(l, r, target);
        } else {
            return nums[mid];
        }
    }

    return ans;

};

threeSum([0, 0, 0]);
threeSum([-1, 0, 1, 2, -1, -4]);
threeSum([-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6]);

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */