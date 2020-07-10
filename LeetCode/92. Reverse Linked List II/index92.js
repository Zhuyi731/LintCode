/*
 * @lc app=leetcode.cn id=92 lang=javascript
 *
 * [92] 反转链表 II
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
 * @param {number} m
 * @param {number} n
 * @return {ListNode}
 */
var reverseBetween = function (head, m, n) {
  let currentCt = 0;
  let bef = {
    next: head,
  };
  let _head = bef;
  let statBak;
  let reversList = null;
  let cur 
  while (head) {
    currentCt++;
    if (currentCt === m) {
      statBak = head;
    }
    if (currentCt >= m && currentCt <= n) {
       cur = head;
      bef.next = head.next;
      cur.next = reversList;
      reversList = cur;
    } else{
        bef = bef.next;
    }
    head = bef.next

    if (currentCt === n) {
      // 结束反转
      let next = bef.next 
      bef.next = reversList;
      statBak.next = next
      break;
    }
  }
  return _head.next
};
// @lc code=end
const { ListNodeGeneratorByArr } = require("../../js_utils/index");

reverseBetween(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 1, 5);
reverseBetween(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 1, 2);
reverseBetween(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 2, 4);

