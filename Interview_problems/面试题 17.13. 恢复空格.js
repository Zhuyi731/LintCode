// 哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

// 注意：本题相对原题稍作改动，只需返回未识别的字符数

//

// 示例：

// 输入：
// dictionary = ["looked","just","like","her","brother"]
// sentence = "jesslookedjustliketimherbrother"
// 输出： 7
// 解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
// 提示：

// 0 <= len(sentence) <= 1000
// dictionary中总字符数不超过 150000。
// 你可以认为dictionary和sentence中只包含小写字母。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/re-space-lcci
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

/**
 * @param {string[]} dictionary
 * @param {string} sentence
 * @return {number}
 */
var respace = function (dictionary, sentence) {
    if(!sentence.length){
        return 0
    }
  let dp = [1];
  let lenList = dictionary.map((el) => el.length);
  let firstStr = sentence[0];
  if (dictionary.includes(firstStr)) {
    dp[0] = 0;
  }

  for (let i = 1; i < sentence.length; i++) {
    let min = 999999;
    for (let j = 0; j < lenList.length; j++) {
      if (i - lenList[j] >= -1) {
        let subStr = sentence.substring(i - lenList[j] + 1, i + 1);
        if (dictionary[j] === subStr) {
          let prevMin = i - lenList[j] >= 0 ? dp[i - lenList[j]] : 0;
          min = min > prevMin ? prevMin : min;
        }
      }
    }
    dp[i] = min > dp[i - 1] + 1 ? dp[i - 1] + 1 : min;
  }
  return dp[sentence.length - 1];
};
respace(["potimzz"], "potimzzpotimzz");
respace(
  ["e", "looked", "just", "like", "her", "brother"],
  "jesslookedjustliketimherbrother"
); // 6
respace(
  ["j", "looked", "just", "like", "her", "brother"],
  "jesslookedjustliketimherbrother"
); // 6
respace(["e"], "jesslookedjustliketimherbrother"); // 26
respace(
  ["looked", "just", "like", "her", "brother"],
  "jesslookedjustliketimherbrother"
); //7
