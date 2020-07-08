/*
 * @lc app=leetcode.cn id=88 lang=javascript
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
  let index1 = m - 1;
  let index2 = n - 1;
  let curIndex = m + n - 1;
  while (index2 >= 0 && index1 >= 0) {
    if (nums1[index1] < nums2[index2]) {
      nums1[curIndex] = nums2[index2];
      index2--;
    } else {
      nums1[curIndex] = nums1[index1];
      index1--;
    }
    curIndex--;
  }

  if (index2 >= 0) {
    while (index2 >= 0) {
      nums1[index2] = nums2[index2--];
    }
  }

  console.log(nums1);
};
// @lc code=end

merge([1], 1, [1], 1);
merge([1], 1, [0], 0);
merge([0], 0, [1], 1);
merge([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3);
