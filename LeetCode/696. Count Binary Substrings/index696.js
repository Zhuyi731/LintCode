/*
 * @lc app=leetcode.cn id=696 lang=javascript
 *
 * [696] 计数二进制子串
 */

// @lc code=start
/**
 * @param {string} s
 * @return {number}
 */
var countBinarySubstrings = function (s) {
  let ans = 0;
  let i = 0;
  while (i < s.length) {
    let currentCt = [0, 0];
    let curIndex = i;
    let pre = s[i];
    let currentMin;
    while (curIndex <= s.length) {
      const isLast = curIndex === s.length;
      if (s[curIndex] !== pre || isLast) {
        if ((currentCt[0] === 0 || currentCt[1] === 0) && !isLast) {
          // 说明需要切换
          currentCt[s[curIndex]]++;
          pre = s[curIndex];
          curIndex++;
        } else {
          // 说明结束  开始计算
          currentMin = Math.min(currentCt[0], currentCt[1]);
          ans += currentMin;
          break;
        }
      } else {
        currentCt[s[curIndex]]++;
        curIndex++;
      }
    }
    i += currentCt[(~~pre + 1) % 2] || "1";
  }
  console.log(ans);
  return ans;
};
// @lc code=end
countBinarySubstrings("00100"); //2
countBinarySubstrings("00110"); //3
countBinarySubstrings("00110011"); // 6
countBinarySubstrings("10101"); //4
