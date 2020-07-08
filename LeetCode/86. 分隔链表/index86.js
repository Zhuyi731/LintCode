/*
 * @lc app=leetcode.cn id=86 lang=javascript
 *
 * [86] 分隔链表
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
 * @param {number} x
 * @return {ListNode}
 */
var partition = function (head, x) {
  let stored = {
    next: null,
  };

  let bef = {
    next: head,
  };
  let storedbak = stored;
  let headBak = bef;
  while (head) {
    if (head.val >= x) {
      stored.next = head;
      stored = stored.next;
      bef.next = head.next;
      head = head.next;
      stored.next = null
    } else {
      bef = head;
      head = head.next;
    }
  }
  bef.next = storedbak.next;
  return headBak.next;
};
// @lc code=end

const { ListNodeGeneratorByArr } = require("../../js_utils/index");

partition(ListNodeGeneratorByArr([4, 2, 3, 4, 5]), 3);
partition(ListNodeGeneratorByArr([1, 4, 3, 2, 5, 2]), 3);
partition(ListNodeGeneratorByArr([4, 2, 3, 4, 5]), 2);
