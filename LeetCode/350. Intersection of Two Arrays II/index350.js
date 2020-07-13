/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var intersect = function (nums1, nums2) {
  let result = [];
  let tmp = {};

  nums1.forEach((el) => {
    tmp[el] = tmp[el] ? tmp[el] + 1 : 1;
  });
  nums1.forEach((el) => {
    if (tmp[el]) {
      tmp[el]--;
      result.push(el);
    }
  });
  return result;
};
