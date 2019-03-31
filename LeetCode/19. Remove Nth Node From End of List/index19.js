/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
function Node(val) {
    this.val = val;
    this.next = null;
}

function creatNodeList(array) {
    let head = new Node(array[0]),
        cur = head;

    array.forEach((el, index) => {
        if (index) {
            cur.next = new Node(el);
            cur = cur.next;
        }
    });

    return head;
}
var removeNthFromEnd = function (head, n) {
    let current = head;
    let lastNNodes = new Array(n + 1);

    while (current.next != null) {
        lastNNodes.shift();
        lastNNodes.push(current);
        current = current.next
    }
    lastNNodes.shift();
    lastNNodes.push(current);

    if (lastNNodes[0] == null) { //说明是删除的第一个节点
        head = lastNNodes[2];
    } else {
        lastNNodes[0].next = lastNNodes[1].next;
    }
    if (!head) {
        return {
            val: null,
            next: null
        }
    }
    return head;
};
removeNthFromEnd(creatNodeList([1]), 1);
removeNthFromEnd(creatNodeList([1, 2, 3, 4, 5]), 2);
//
// Given a linked list, remove the n-th node from the end of list and return its head.

// Example:

// Given linked list: 1->2->3->4->5, and n = 2.

// After removing the second node from the end, the linked list becomes 1->2->3->5.
// Note:

// Given n will always be valid.

// Follow up:

// Could you do this in one pass?