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
  let next = head.next;
  let prev_bak = prev;
  while (cur && next) {
    if (next.val === cur.val) {
      cur.next = next.next;
      next = cur.next;
    } else {
      cur = cur.next;
      cur && (next = cur.next);
    }
  }
  return prev_bak.next;
};

const { ListNodeGeneratorByArr } = require("../../js_utils/index");

deleteDuplicates(ListNodeGeneratorByArr([1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([1, 2, 3, 3, 4, 4, 5]));
deleteDuplicates(ListNodeGeneratorByArr([1, 1, 1, 2, 3]));
deleteDuplicates(ListNodeGeneratorByArr([3, 2, 1, 1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([1, 1, 1]));
deleteDuplicates(ListNodeGeneratorByArr([]));

// Given a sorted linked list, delete all duplicates such that each element appear only once.

// Example 1:

// Input: 1->1->2
// Output: 1->2
// Example 2:

// Input: 1->1->2->3->3
// Output: 1->2->3
