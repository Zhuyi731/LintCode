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

/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function (lists) {
    let allArr = [] 
    lists.forEach(list=>{
        allArr.concat(ListNodeToArr(list))
    })
    allArr = allArr.sort((a,b)=>a-b)
    let result = ListNodeGeneratorByArr(allArr)
};
