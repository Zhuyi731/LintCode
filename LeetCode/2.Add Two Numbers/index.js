/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
function ListNode(val) {
  this.val = val;
  this.next = null;
}
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
  let ans = new ListNode(null),
    ret = ans,
    flag = 0,
    nextNode,
    tail;

  while ((l1 && l1.val != null) || (l2 && l2.val != null)) {
    l1 == null && (l1 = {});
    l2 == null && (l2 = {});
    ans.val += (~~l1.val + ~~l2.val);
    flag = ans.val >= 10 ? 1 : 0;
    ans.val = ans.val % 10;
    ans.next = new ListNode(flag);
    tail = ans;
    ans = ans.next;
    l1 = l1.next;
    l2 = l2.next;
  }
  if (ans.val == 0) {
    tail.next = null;
  }
  return ret;
};

function creatList(arr) {
  let head = new ListNode(),
    ret = head,
    tail;
  arr.forEach(el => {
    head.val = el;
    head.next = new ListNode();
    tail = head;
    head = head.next;
  });
  tail.next = null;
  return ret;
}
let listA = new ListNode(2);
listA.next = new ListNode(4);
listA.next.next = new ListNode(3);
let listB = new ListNode(5);
listB.next = new ListNode(6);
listA.next.next = new ListNode(3);

// let listA = new ListNode(2);
// listA.next = new ListNode(4);
// listA.next.next = new ListNode(3);
// let listB = new ListNode(5);
// listB.next = new ListNode(6);
// listB.next.next = new ListNode(4);
pf(addTwoNumbers(creatList([2,4,6,1,3]), creatList([5,6,4])));
pf(addTwoNumbers(creatList([0]), creatList([0])));

function pf(ans) {
  console.log(ans.val);
  ans = ans.next;
  while (ans) {
    console.log(`-->${ans.val}`);
    ans = ans.next;
  }
}

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.