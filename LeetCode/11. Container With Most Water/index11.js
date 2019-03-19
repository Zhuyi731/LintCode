/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function(height) {
    let ans = 0,
        i, j,
        len = height.length;

    for (i = 0; i < len - 1; i++) {
        for (j = i + 1; j < len; j++) {
            let contain = Math.min(height[i], height[j]) * (j - i);
            ans = (ans > contain) ? ans : contain;
        }
    }
    console.log(ans);
    return ans;
};
maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]);

/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
 */