/*
 * @lc app=leetcode.cn id=42 lang=javascript
 *
 * [42] 接雨水
 */

// @lc code=start
/**
 * @param {number[]} height
 * @return {number}
 */
var trap = function (height) {
  //
  function findHeight(l, r, arr) {
    let max = arr[l];
    let maxIndex = l;
    for (let i = l; i <= r; i++) {
      if (arr[i] > max) {
        max = arr[i];
        maxIndex = i;
      }
    }
    return [max, maxIndex];
  }

  let sum = 0;
  function findAndSum(l, r) {
    if (l >= r) {
      return;
    }
    if (l === 0) {
      //往左找
      let [maxHeight, maxIndex] = findHeight(l, r, height);
      for (let i = maxIndex + 1; i <= r; i++) {
        sum += maxHeight - height[i];
      }
      findAndSum(0, maxIndex - 1);
    } else {
      let [maxHeight, maxIndex] = findHeight(l, r, height);
      for (let i = l; i < maxIndex; i++) {
        sum += maxHeight - height[i];
      }
      findAndSum(maxIndex + 1, r);
    }
  }

  let [maxHeight, maxIndex] = findHeight(0, height.length - 1, height);
  findAndSum(0, maxIndex - 1);
  findAndSum(maxIndex + 1, height.length - 1);
  console.log(sum);
  return sum;
};
// @lc code=end
// 3 1 6 5 1 0 4 3 6 1 6 6
//     .             .   . .
//     . .           .   . .
//     . .     .     .   . .
// .   . .     .  .  .   . .
// .   . .     .  .  .   . .
// . . . . .   .  .  . . . .
trap([11116, 3, 2, 1, 612, 327, 566, 326, 116, 7676, 16, 26]); //59465
trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]); //6
trap([3, 1, 6, 5, 1, 0, 4, 3, 6, 1, 6, 6]); //24
trap([16, 3, 2, 1, 612, 327, 566, 326, 116, 7676, 16, 26]); // 1165
trap([6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6]); //6
trap([6, 6, 6, 6, 6, 7, 6, 6, 6, 6, 6, 6]); //6
// [3, 1, 6, 5, 1, 0, 4, 3, 6, 1, 6, 6]\n[16, 3, 2, 1, 612, 327, 566, 326, 116, 7676, 16, 26]\n[11116, 3, 2, 1, 612, 327, 566, 326, 116, 7676, 16, 26]
