/**
 * @param {number[][]} nums
 * @return {number[]}
 */
var smallestRange = function (nums) {
  // 首先利用插入排序  将所有数字排序
  let mergedArr = nums[0].map((el) => ({
    val: el,
    index: 0,
  }));
  for (let i = 1; i < nums.length; i++) {
    let index = 0;
    let curArr = nums[i];
    for (let j = 0; j < curArr.length; j++) {
      if (index === mergedArr.length) {
        mergedArr.push({
          val: curArr[j],
          index: i,
        });
        index++;
        continue;
      }
      if (curArr[j] < mergedArr[index].val) {
        mergedArr.splice(index, 0, {
          val: curArr[j],
          index: i,
        });
        index++;
      } else {
        index++;
        j--;
      }
    }
  }

//   console.log(mergedArr);

  // 然后遍历找最小区间
  const defaultArr = new Array(nums.length).fill(0);
  // 用来标志是否已经找全
  let minArea = 0x3ffff;
  let minResult = [];
  let currentResult = [];
  for (let i = 0; i <= mergedArr.length - nums.length; i++) {
    let curIndex = i;
    let curFillResult = defaultArr.map((el) => el);
    currentResult[0] = mergedArr[curIndex].val;
    if(i && mergedArr[i].val === mergedArr[i-1].val) continue
    while (curIndex < mergedArr.length) {
      if (mergedArr[curIndex].index - currentResult[0] > minArea) {
        break;
      }
      curFillResult[mergedArr[curIndex].index] = 1;
      if (curFillResult.findIndex((el) => !el) === -1) {
        // 说明结束了
        currentResult[1] = mergedArr[curIndex].val;
        // 比较大小
        if (currentResult[1] - currentResult[0] < minArea) {
          minArea = currentResult[1] - currentResult[0];
          minResult = [currentResult[0], currentResult[1]];
        }
        // console.log(i);
        // console.log(currentResult);
        break;
      }
      curIndex++;
    }
  }

//   console.log(minResult);
  return minResult;
};
const test = require("./testcase.json");
smallestRange(test);
smallestRange([[10]]);

smallestRange([[10], [11]]);

smallestRange([
  [4, 10, 15, 24, 26],
  [0, 9, 12, 20],
  [5, 18, 22, 30],
]);
