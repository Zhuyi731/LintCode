/*
 * @lc app=leetcode.cn id=82 lang=javascript
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function (head) {
  if (!head) return head;
  let prev = {
    next: head,
  };
  let cur = head;
  let next = cur.next;
  let prev_bak = prev;
  let currentDeleteVal = "xxx";
  while (cur) {
    if (cur.val === currentDeleteVal) {
      prev.next = cur.next;
      cur = prev.next;
      cur && (next = cur.next);
    } else if (next && cur.val === next.val) {
      currentDeleteVal = cur.val;
      prev.next = next.next;
      cur = prev.next;
      cur && (next = cur.next);
    } else {
      prev = prev.next;
      cur = prev.next;
      cur && (next = cur.next);
    }
  }
  return prev_bak.next;
};
// @lc code=end

const { ListNodeGeneratorByArr } = require("../../js_utils/index");

deleteDuplicates(ListNodeGeneratorByArr([1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([1, 2, 3, 3, 4, 4, 5]));
deleteDuplicates(ListNodeGeneratorByArr([1, 1, 1, 2, 3]));
deleteDuplicates(ListNodeGeneratorByArr([3, 2, 1, 1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([1, 1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([]));
