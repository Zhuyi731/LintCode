/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var inorderTraversal = function (root) {
  if (!root) return []
  let left = []
  let right = []
  if (root.left) {
    left = inorderTraversal(root.left)
  }
  if (root.right) {
    right = inorderTraversal(root.right)
  }

  if (!root.left && !root.right) {
    return [root.val]
  }

  return left.concat([root.val]).concat(right)
}

function TreeNode(val) {
  this.val = val
  this.left = this.right = null
}

function parseTreeNode(param) {
  if (!Array.isArray(param)) {
    onParameterError()
  }

  let root = null
  const fifo = []
  let i = 0
  while (i < param.length) {
    if (i === 0) {
      root = new TreeNode(param[i])
      i += 1
      fifo.push(root)
      continue
    }
    const parent = fifo.shift()
    if (param[i] != null) {
      const left = new TreeNode(param[i])
      parent.left = left
      fifo.push(left)
    }
    if (i + 1 < param.length && param[i + 1] != null) {
      const right = new TreeNode(param[i + 1])
      parent.right = right
      fifo.push(right)
    }
    i = i + 2
  }
  return root
}

console.log(inorderTraversal(parseTreeNode([1, 2, 3, 4, 5, null, 7])))
console.log(inorderTraversal(parseTreeNode([1, 2, 3, 4, 5, 6, 7])))
console.log(
  inorderTraversal(
    parseTreeNode([1, 3, null, 2, 4, 6, null, 1, null, 3, 5, 6, 87, 9, 1])
  )
)
