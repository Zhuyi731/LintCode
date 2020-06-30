/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function (l1, l2) {
    if(!l1 || !l2) return l2;

    if(l1.val<l2.val){
        l1.next = mergeTwoLists(l1.next, l2);
        return l1;
    } else {
        l2.next = mergeTwoLists(l1,l2.next);
        return l2;
    }
};

const { ListNodeGeneratorByArr } = require("../../js_utils/index");

function test(a, b) {
  a = ListNodeGeneratorByArr(a);
  b = ListNodeGeneratorByArr(b);
  mergeTwoLists(a, b);
}

test([-10,-9,-9,0,1,8],
    [-1,9])
test([2], [1]);
test([], [0]);
