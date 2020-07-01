/*
 * @lc app=leetcode.cn id=61 lang=javascript
 *
 * [61] 旋转链表
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
 * @param {number} k
 * @return {ListNode}
 */
var rotateRight = function (head, k) {
  let len = 1;
  let head_bak = head;
  while (head_bak && head_bak.next) {
    len++;
    head_bak = head_bak.next;
  }
  let tailNode = head_bak;
  let headNode = { next: head };
  k = k % len;
  if (len <= 1 || k <= 0) return head;

  k = len - k;
  // 找到第k个节点

  let kNode;
  let ct = 1;
  head_bak = head;
  while (ct < k) {
    head_bak = head_bak.next;
    ct++;
  }
  // 需要在这里做阶段
  let curNode = head_bak;
  let nextNode = head_bak.next;

  curNode.next = null;
  tailNode.next = headNode.next;
  headNode.next = nextNode;

  return headNode.next;
};
const { ListNodeGeneratorByArr } = require("../../js_utils/index");

rotateRight(ListNodeGeneratorByArr([1]), 1);
rotateRight(ListNodeGeneratorByArr([]), 5);
rotateRight(ListNodeGeneratorByArr([0, 1, 2]), 4);
rotateRight(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 2);
rotateRight(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 3);

// @lc code=end
