// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

// You are given a target value to search. If found in the array return its index, otherwise return -1.

// You may assume no duplicate exists in the array.

// Your algorithm's runtime complexity must be in the order of O(log n).

// Example 1:

// Input: nums = [4,5,6,7,0,1,2], target = 0
// Output: 4
// Example 2:

// Input: nums = [4,5,6,7,0,1,2], target = 3
// Output: -1
/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function (nums, target) {
	if (nums.length === 0) return -1
	if (nums.length === 1) return nums[0] === target ? 0 : -1
	// 首先找到分割点  
	let len = nums.length
	let mid = Math.floor(len / 2)
	let left = 0
	let right = len
	while (true) {
		if (nums[mid] > nums[mid + 1]) {
			break;
		}
		if (nums[0] < nums[len - 1]) {
			mid = len - 1
			break
		}
		// if (mid === 1 || mid === len - 2) {
		// 	if (nums[0] > nums[1]) {
		// 		mid = 0
		// 		break;
		// 	}
		// 	if (nums[len - 2] > nums[len - 1]) {
		// 		mid = len - 2
		// 		break
		// 	}
		// 	mid = len
		// 	break
		// }

		if (nums[mid] > nums[0]) {
			//说明在右边
			left = mid
			mid = Math.floor((left + right) / 2)
		} else {
			//在左边 
			right = mid
			mid = Math.floor((right + left) / 2)
		}
	}
	// 然后查看当前数在哪个区间  
	if (target > nums[0]) {
		//说明在左边区间或者不存在 
		let temp = nums.slice(0, mid + 1)
		return binarySearch(temp, target)
	} else if (target < nums[0]) {
		//说明在右边区间或者不存在
		let temp = nums.slice(mid + 1, len)
		let ans = binarySearch(temp, target)
		return ans === -1 ? -1 : mid + ans + 1
	} else {
		return 0
	}

	function binarySearch(arr, x) {
		let l = 0
		let r = arr.length
		let mid = Math.floor((l + r) / 2)
		while (l < r) {
			if (l === r - 1) {
				if (x === arr[l]) {
					return l
				} else if (x === arr[r]) {
					return r
				} else {
					return -1
				}
			}
			if (x < arr[mid]) {
				// 左边  
				r = mid
				mid = Math.floor((l + r) / 2)
			} else if (x > arr[mid]) {
				l = mid
				mid = Math.floor((l + r) / 2)
			} else {
				return mid
			}
		}
		return -1
	}
};

console.log(search([3, 4, 5, 6, 7, 8, 1, 2], 2), 7)
console.log(search([7, 8, 1, 2, 3, 4, 5, 6], 2), 3)
console.log(search([3, 5, 1], 5), 1)
console.log(search([1, 3], 3), 1)
console.log(search([1, 3], 1), 0)
console.log(search([3, 1], 1), 1)
console.log(search([3, 1], 3), 0)
console.log(search([3, 1], 2), -1)
console.log(search([1, 3], 0), -1)
console.log(search([4, 5, 6, 7, 0, 1, 2], 2), 6)
console.log(search([4, 5, 6, 7, 0, 1, 2], 4), 0)
console.log(search([4, 5, 6, 7, 0, 1, 2], 5), 1)
console.log(search([4, 5, 6, 7, 0, 1, 2], 0), 4)
console.log(search([4, 5, 6, 7, 0, 1, 2], 3), -1)