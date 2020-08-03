/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function (num1, num2) {
  let result = "";
  let tag = 0;
  let i = num1.length - 1;
  let j = num2.length - 1;

  while (i >= 0 || j >= 0) {
    let x = i >= 0 ? num1.charAt(i) - "0" : 0;
    let y = j >= 0 ? num2.charAt(j) - "0" : 0;
    let temp = x + y + tag;
    tag = Math.floor(temp / 10);
    result = (temp % 10) + result;
    i--;
    j--;
  }
  if (tag === 1) {
    result = 1 + result;
  }
  console.log(result);
  return result;
};

addStrings("9999", "1");
