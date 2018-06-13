/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 * 
 * 快速排序    通过75%数据后TLE
 */
const kthLargestElement = function (n, nums) {
    let newArr = [],
        type = 1,
        i,
        len = nums.length,
        ct = 0;

    if (n < len / 2) {
        type = -1;
        n = len - n + 1;
    }

    newArr[0] = nums[0];
    for (i = 1; i < len; i++) {
           insert(i);
    }

    function insert(index) {
        if(index<ct){
            //遍历newArr 插入排序
            for(let j = 0;j < ct;j++){
                if(type == 1 && nums[i] > newArr[j])
            }
        }
    }
}





kthLargestElement(3, [9, 2, 2, 4, 2]);