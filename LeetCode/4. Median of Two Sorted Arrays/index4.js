/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function(nums1, nums2) {
    if (nums1[0] > nums2[0]) {
        let temp;
        temp = nums1;
        nums1 = nums2;
        nums2 = temp;
    }
    let len1 = nums1.length,
        len2 = nums2.length,
        totalLen = len1 + len2,
        mid = Math.floor((totalLen - 1) / 2),
        ct = 0,
        ans = [],
        len1Index = 0,
        currentIndex = 0,
        len2Index = 0;

    isOk(0, 1);


    while (ct < totalLen) {
        if (nums1[len1Index] > nums2[len2Index]) {
            len2Index++;
        } else if (nums1[len1Index] == nums2[len2Index]) {
            len1Index++;
            len2Index++;
        } else {

        }
    }

    function isOk(cur, type) {
        if (ansType == 0) {
            if (cur == mid || cur == mid + 1) {
                ans.push(type == 1 ? nums1[len1Index] : nums2[len2Index]);
            }
        } else {
            if (cur == mid) {
                ans.push(type == 1 ? nums1[len1Index] : nums2[len2Index])
            }
        }

    }
};