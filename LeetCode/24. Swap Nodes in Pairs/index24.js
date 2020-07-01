/*
 * @lc app=leetcode.cn id=24 lang=javascript
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

const { ListNodeGeneratorByArr } = require("../../js_utils/index");

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var swapPairs = function (head) {
    if(!head || !head.next) return head
    let curNode = head;
    let nextNode = curNode.next;
  
    head = nextNode;
    let tail = nextNode.next;
    nextNode.next = curNode;
    curNode.next = tail;
    let bef = curNode;
  
    while (curNode && curNode.next && curNode.next.next) {
      curNode = curNode.next;
      nextNode = curNode.next;
      tail = nextNode.next;
      bef.next = nextNode;
      nextNode.next = curNode;
      curNode.next = tail;
      bef = curNode;
    }
    return head;
};

swapPairs(ListNodeGeneratorByArr([1, 2, 3, 4]));
// @lc code=end

// @after-stub-for-debug-begin
module.exports = swapPairs;
// @after-stub-for-debug-end