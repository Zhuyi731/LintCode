 
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
// 节点生成器  将 1->3->2->12这样的字符串，转为ListNode

function ListNodeGenerator(nodeStr) {
    let nodes = nodeStr.split('->')
    let head = {
        val: nodes[0],
        next:null 
    } 
    let headCopy = head 
    for(let i=0;i<nodes.length;i++){
        
    }
    nodes.forEach((node, index)=>{
        headCopy.val = node 
        headCopy.next = 
    })
    head = nodes[0]

}

module.exports = {
    ListNodeGenerator
}