/**
 * @param {number} numCourses
 * @param {number[][]} prerequisites
 * @return {boolean}
 */
var canFinish = function (numCourses, prerequisites) {
  let maps = [];

  prerequisites.forEach((el) => {
    if (!maps[el[0]]) {
      maps[el[0]] = [];
      maps[el[0]][el[1]] = 1;
    } else {
      maps[el[0]][el[1]] = 1;
    }
  });

  // 得到了一个有向图
  // 然后遍历找是否有闭环
  const searched = [];

  function dfs(current, used) {
    searched[current] = true;
    if (current >= maps.length) {
      return true;
    }

    used[current] = true;
    if (!maps[current]) {
      maps[current] = [];
    }
    for (let i = 0; i < maps[current].length; i++) {
      if (!maps[current][i]) continue;
      if (used[i]) {
        return false;
      }
      const usedCopy = used.map((el) => el);
      if (dfs(i, usedCopy) === false) {
        return false;
      }
    }
    return true;
  }

  for (let i = 0; i < numCourses; i++) {
    if (!maps[i]) maps[i] = [];

    for (let j = 0; j < maps[i].length; j++) {
      if (!maps[i][j] || searched[j]) {
        continue;
      }
      let used = [];
      used[i] = 1;
      if (dfs(j, used) === false) {
        console.log(false);
        return false;
      }
    }
  }
  console.log(true);
  return true;
};
canFinish(4, [
  [1, 0],
  [0, 3],
  [0, 2],
  [3, 2],
  [2, 5],
  [4, 5],
  [5, 6],
  [2, 4],
]);
canFinish(4, [
  [0, 1],
  [1, 2],
  [0, 3],
  [3, 0],
]); // false
canFinish(3, [
  [0, 1],
  [0, 2],
  [1, 2],
]); // true
canFinish(3, [
  [1, 0],
  [2, 1],
]); // true
canFinish(2, [[1, 0]]); // true
canFinish(2, [[0, 1]]); // true
canFinish(2, [
  //false
  [1, 0],
  [0, 1],
]);
