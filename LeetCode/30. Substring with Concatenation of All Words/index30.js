/*
 * @lc app=leetcode.cn id=30 lang=javascript
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string[]} words
 * @return {number[]}
 */
var findSubstring = function (s, words) {
  function cmp(index, len, hashMap) {
    let word = s.substr(index, len);
    if (hashMap[word]) {
      hashMap[word]--;
      if (hashMap[word] === 0) {
        delete hashMap[word];
      }
      return true;
    } else {
      return false;
    }
  }

  if (!words.length || !s.length) return [];
  let len = words[0].length;
  let result = [];
  let hashMap = {};
  words.forEach((word) => {
    if (hashMap[word]) {
      hashMap[word] += 1;
    } else {
      hashMap[word] = 1;
    }
  });

  for (let i = 0; i < len; i++) {
    // 每个节点都计算一次
    let ct;
    let ct2 = i;
    while (ct2 < s.length) {
      // 滑动窗口
      ct = ct2;
      let hashMapCopy = Object.assign({}, hashMap);
      while (ct < ct2 + len * words.length) {
        if (!cmp(ct, len, hashMapCopy)) {
          break;
        }
        ct = ct + len;
      }
      if (Object.keys(hashMapCopy).length === 0) {
        // 说明成功了
        // hashMapCopy = Object.assign({}, hashMap);
        result.push(ct - len * words.length);
      }
      ct2 += len;
    }
  }
  console.log(result);
  return result;
};
// @lc code=end

findSubstring("ababbaababa", ["ab", "ba", "ab"]);
findSubstring("wordgoodgoodgoodbestword", ["word", "good", "best", "word"]);
findSubstring("barfoothefoobarman", ["foo", "bar"]);
