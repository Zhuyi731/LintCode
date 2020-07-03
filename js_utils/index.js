/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
// 节点生成器  将 1->3->2->12这样的字符串，转为ListNode

function ListNodeGenerator(nodeStr) {
  let nodes = nodeStr.split("->");
  let head = {
    val: nodes[0],
    next: null,
  };
  let headCopy = head;
  for (let i = 0; i < nodes.length; i++) {}
  nodes.forEach((node, index) => {
    headCopy.val = node;
    headCopy.next = null;
  });
  head = nodes[0];
}

function ListNodeGeneratorByArr(nodes) {
  if (!nodes.length) return null;
  let head = {
    val: nodes[0],
    next: null,
  };
  let curPtr = head;
  for (let i = 1; i < nodes.length; i++) {
    let node = nodes[i];
    let curNode = {
      val: node,
      next: null,
    };
    curPtr.next = curNode;
    curPtr = curPtr.next;
  }
  return head;
}

function ListNodeToArr(node) {
  let arr = [];
  while (node) {
    arr.push(node.value);
    node = node.next;
  }
  return arr;
}

module.exports = {
  ListNodeGenerator,
  ListNodeGeneratorByArr,
  ListNodeToArr,
};
