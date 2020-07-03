/*
 * @lc app=leetcode.cn id=378 lang=javascript
 *
 * [378] 有序矩阵中第K小的元素
 */

/**
 * @param {number[][]} matrix
 * @param {number} k
 * @return {number}
 */
var kthSmallest = function (matrix, k) {
  let len = matrix.length;

  let l = matrix[0][0];
  let r = matrix[len - 1][len - 1];
  let mid = l + ((r - l) >> 1);

  while (l < r) {
    let x = len - 1;
    let y = 0;
    let ptr;
    let record = new Array(len - 1).fill(len);
    while (x >= 0 && y < len) {
      ptr = matrix[x][y];
      if (ptr > mid) {
        record[x] = y;
        x--;
      } else {
        y++;
      }
    }

    record = record.map((el) => {
      console.log(el);
      return el === undefined ? len - 1 : el;
    });
    let sum = record.reduce((prev, cur) => {
      return prev + cur;
    }, 0);

    if (sum < k) {
      l = mid + 1;
      mid = l + ((r - l) >> 1);
    } else {
      r = mid;
      mid = l + ((r - l) >> 1);
    }
  }
  return mid;
};

kthSmallest(
  [
    [1, 5, 9],
    [10, 11, 13],
    [12, 13, 15],
  ],
  8
);
