/*
 * @lc app=leetcode.cn id=108 lang=javascript
 *
 * [108] 将有序数组转换为二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {number[]} nums
 * @return {TreeNode}
 */
var sortedArrayToBST = function (nums) {
  if (nums.length === 0) {
    return null;
  }
  if (nums.length === 1) {
    return {
      val: nums[0],
      left: null,
      right: null,
    };
  }

  let mid = nums.length >> 1;
  let l = 0;
  let r = nums.length;
  let node = {
    val: nums[mid],
    left: null,
    right: null,
  };
  node.left = sortedArrayToBST(nums.slice(0, mid));
  node.right = sortedArrayToBST(nums.slice(mid + 1));
  return node;
};

const { ListNodeToArr } = require("../../js_utils/index");
console.log(ListNodeToArr(sortedArrayToBST([-10, -3, 0, 5, 9])));
// @lc code=end
