/**
 * @param {number[][]} edges
 * @return {number[]}
 */
class Union {
  constructor(len) {
    let parent = new Array(len).fill(0)
    this.parent = parent.map((el, index) => index)
  }

  find(x) {
    while (x !== this.parent[x]) {
      x = this.parent[x]
    }
    return x
  }

  union(x, y) {
    let rootX = this.find(x)
    let rootY = this.find(y)
    if (rootX == rootY) {
      return false
    }
    this.parent[rootY] = rootX
    return true
  }
}

function judgeCircle(edges, removedEdge) {
  const union = new Union(edges.length + 1)
  for (let i = 0; i < edges.length; i++) {
    if (i === removedEdge) {
      continue
    }
    if (!union.union(edges[i][0], edges[i][1])) {
      return false
    }
  }
  return true
}

var findRedundantDirectedConnection = function (edges) {
  // 首先计算 入度

  let inDegree = []
  for (let i = 0; i < edges.length; i++) {
    inDegree[edges[i][1]] = inDegree[edges[i][1]]
      ? inDegree[edges[i][1]] + 1
      : 1
  }

  // 然后遍历度数
  for (let i = edges.length - 1; i >= 0; i--) {
    if (inDegree[edges[i][1]] === 2) {
      // 如果度数为2  尝试删除其中的一条边
      if (judgeCircle(edges, i)) {
        return edges[i]
      }
    }
  }

  for (let i = edges.length - 1; i >= 0; i--) {
    // 没有度数为2的 尝试删除其中的一条边 查看是否成环
    if (judgeCircle(edges, i)) {
      return edges[i]
    }
  }
  return []
}

console.log(
  findRedundantDirectedConnection([
    [1, 2],
    [1, 3],
    [2, 3],
  ])
)
console.log(
  findRedundantDirectedConnection([
    [1, 2],
    [2, 3],
    [3, 4],
    [4, 1],
    [1, 5],
  ])
)
console.log(
  findRedundantDirectedConnection([
    [1, 2],
    [2, 3],
    [3, 1],
    [4, 2],
  ])
)

console.log(
  findRedundantDirectedConnection([
    [1, 2],
    [2, 3],
    [3, 1],
  ])
)
