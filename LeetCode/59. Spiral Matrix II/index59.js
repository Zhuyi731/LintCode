/**
 * @param {number} n
 * @return {number[][]}
 */

function changeDirection(currentDirection) {
  if (currentDirection[1] == 1) {
    return [1, 0]
  } else if (currentDirection[1] == -1) {
    return [-1, 0]
  } else if (currentDirection[0] == 1) {
    return [0, -1]
  } else if (currentDirection[0] == -1) {
    return [0, 1]
  }
}

function checkOverFlow(nextCursor, n, ans) {
  if (
    nextCursor[0] >= n ||
    nextCursor[0] < 0 ||
    nextCursor[1] >= n ||
    nextCursor[1] < 0 ||
    ans[nextCursor[0]][nextCursor[1]]
  ) {
    return true
  }
  return false
}

var generateMatrix = function (n) {
  let ans = []
  for (let i = 0; i < n; i++) {
    ans[i] = []
  }
  let cursor = [0, 0]
  let now = 1
  let direction = [0, 1]
  while (1) {
    ans[cursor[0]][cursor[1]] = now
    now++
    let nextCursor = [cursor[0] + direction[0], cursor[1] + direction[1]]
    if (checkOverFlow(nextCursor, n, ans)) {
      direction = changeDirection(direction)
      nextCursor = [cursor[0] + direction[0], cursor[1] + direction[1]]
      if (checkOverFlow(nextCursor, n, ans)) {
        break
      }
    }
    cursor = nextCursor
  }
  console.log(ans)
  return ans
}

generateMatrix(3)
generateMatrix(4)
generateMatrix(5)
