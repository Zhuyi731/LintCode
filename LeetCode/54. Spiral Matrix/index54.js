/**
 * @param {number[][]} matrix
 * @return {number[]}
 */
var spiralOrder = function (matrix) {
  if (matrix.length == 0) {
    return []
  }
  var curIndex = [0, 0]
  var result = []
  var direction = 0
  const directions = [
    [0, 1],
    [1, 0],
    [0, -1],
    [-1, 0],
  ]
  var lenx = matrix.length
  var leny = matrix[0].length
  while (result.length !== lenx * leny) {
    result.push(matrix[curIndex[0]][curIndex[1]])
    matrix[curIndex[0]][curIndex[1]] = "x"
    let prex = curIndex[0] + directions[direction][0]
    let prey = curIndex[1] + directions[direction][1]
    if (
      prex >= lenx ||
      prex < 0 ||
      prey >= leny ||
      prey < 0 ||
      matrix[prex][prey] === "x"
    ) {
      direction = (direction + 1) % 4
      // 检查修改direction后的路径
    }
    curIndex[0] += directions[direction][0]
    curIndex[1] += directions[direction][1]
  }
  return result
}

console.log(
  spiralOrder([
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
  ])
) //[1,2,3,6,9,8,7,4,5]
console.log(
  spiralOrder([
    [1, 2, 3, 4],
    [5, 6, 7, 8],
    [9, 10, 11, 12],
  ])
) //[1,2,3,4,8,12,11,10,9,5,6,7]
