/*
 * @lc app=leetcode.cn id=25 lang=javascript
 *
 * [25] K 个一组翻转链表
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
var reverseKGroup = function (head, k) {
  let emptyHead = {
    next: head,
  };
  let currentNode = head;
  let origin = emptyHead;

  while (currentNode) {
    currentNode = checkAndReverse(origin, currentNode, k);
    if(currentNode){
        origin = currentNode
        currentNode = currentNode.next
    }
  }

  return emptyHead.next;
  function checkAndReverse(origin, currentNode, k) {
    let ct = 1;
    let bef;
    while (currentNode && currentNode.next && ct < k) {
      bef = currentNode;
      currentNode = currentNode.next;
      currentNode.bef = bef;
      ct++;
    }
    if (ct === k) {
      let next = currentNode.next;
      origin.next = currentNode
      while (currentNode.bef) {
          currentNode.next = currentNode.bef 
          delete currentNode.bef
          currentNode = currentNode.next
      }
      // 然后到顶点了  
      currentNode.next = next 
      return currentNode
    } else {
      // 说明当前路径不足K个  不做翻转
    }
  }
};
// @lc code=end

// const { ListNodeGeneratorByArr } = require("../../js_utils/index");

// // reverseKGroup(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 5);
// // reverseKGroup(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 4);
// reverseKGroup(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 3);
// reverseKGroup(ListNodeGeneratorByArr([1, 2, 3, 4, 5]), 2);
