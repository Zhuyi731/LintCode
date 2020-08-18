/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {ListNode} head
 * @return {TreeNode}
 */

const { ListNodeGeneratorByArr } = require("../../js_utils/index");
var sortedListToBST = function (head) {
  if (!head) {
    return null;
  }
  if (head.next == null) {
    return {
      val: head.val,
      left: null,
      right: null,
    };
  }
  let fastPtr = head;
  let slowPtr = head;
  let bef = slowPtr;
  while (fastPtr) {
    fastPtr = fastPtr.next;
    if (!fastPtr) {
      break;
    }
    fastPtr = fastPtr.next;
    bef = slowPtr;
    slowPtr = slowPtr.next;
  }
  // 此时slowPtr为中点
  bef.next = null;
  return {
    val: slowPtr.val,
    left: sortedListToBST(head),
    right: sortedListToBST(slowPtr.next),
  };
};

console.log(sortedListToBST(ListNodeGeneratorByArr([-10, -3, 0, 5, 9])));
console.log(sortedListToBST(ListNodeGeneratorByArr([])));
console.log(sortedListToBST(ListNodeGeneratorByArr([0])));
console.log(sortedListToBST(ListNodeGeneratorByArr([1, 3])));
