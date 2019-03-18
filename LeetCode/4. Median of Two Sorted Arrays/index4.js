/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
//采用插入排序，将较短的数组，插入较长的数组，取中间两位即可

var findMedianSortedArrays = function(nums1, nums2) {
    if (nums1.length > nums2.length) {
        let temp;
        temp = nums1;
        nums1 = nums2;
        nums2 = temp;
    }

    let len1 = nums1.length,
        len2 = nums2.length,
        totalLen = len1 + len2,
        mid = Math.floor(totalLen / 2),
        ans = [],
        num1Index = 0,
        i = 0;

    if (nums1[0] > nums2[0]) {
        nums1.splice(0, 0, nums2[0]);
        num1Index++;
        i++;
        len1++;
    }

    for (; i < len2; i++) {
        inject();
        len1 = nums1.length;
        if (num1Index > mid) {
            break;
        }
    }

    if (totalLen % 2 == 0) { //偶数位
        ans = nums1[mid] + nums1[mid - 1];
        ans /= 2;
    } else {
        ans = nums1[mid];
    }
    console.log(ans);
    return ans;

    function inject() {
        let current = nums2[i];
        while (current > nums1[num1Index] && num1Index < len1) {
            num1Index++;
        }

        nums1.splice(num1Index, 0, current);
    }

};
nums1 = [1, 3]
nums2 = [2]
findMedianSortedArrays(nums1, nums2);