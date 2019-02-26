/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
//a normal solution
var twoSum = function (nums, target) {
  let result,
    pair;

  nums.some((val, index) => {
    result = nums.indexOf(target - val);
    result !== -1 && result !== index && (pair = [index, result]);
    return result !== -1 && result !== index;
  });

  return pair;
};
console.log(twoSum([3,2,4], 6));