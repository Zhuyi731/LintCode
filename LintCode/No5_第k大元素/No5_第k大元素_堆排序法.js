/**
 * @param n: An integer
 * @param nums: An array
 * @return: the Kth largest element
 * 
 * 快速排序    通过75%数据后TLE
 */
const kthLargestElement = function (n, nums) {
    let len = nums.length,
        heapType = n > len / 2 ? "small" : "big",
        tail = len,
        sorted = 0,
        sortedArr = [];
  
        n = len - n + 1;

    while (sorted <= len) {
        // for(let i = len/2-1;sorted < n && i > 0;i--){
        buildHeap(); //调整堆  使其成为小顶堆
        nums.push(nums[0]);
        nums.splice(0, 1); //将顶部放入尾部
        sorted++;
        tail--
        if (sorted == n ) {
            console.log(nums[len - 1]);
            return nums[len - 1];
        }
        // buildHeap(i); //重新调整剩余的堆
    }

    /**
     * 
     * @param {*root节点} root 
     * @param {*调整的叶子节点} node 
     */
    function buildHeap() {

        for (let i = parseInt(tail / 2) - 1; i >= 0; i--) {
            adjustHeap(i);
        }

    }

    function adjustHeap(node) {
        if (heapType == "small") {
            let cMin = node * 2 + 1;
            if (node * 2 + 2 < tail && nums[node * 2 + 1] > nums[node * 2 + 2]) {
                cMin++;
            }
            if (nums[node] > nums[cMin]) {
                swap(node, cMin);
                let lchild, rchild;
                if (cMin * 2 + 1 < tail) {
                    lchild = cMin * 2 + 1;
                    rchild = cMin * 2 + 2 < tail ? Infinity : nums[cMin * 2 + 2];
                    if (nums[cMin] > Math.min(nums[lchild], nums[rchild])) {
                        adjustHeap(cMin);
                    }
                }

            }
        } else {
            let cMax = node * 2 + 1;
            if (node * 2 + 2 < tail && nums[node * 2 + 1] < nums[node * 2 + 2]) {
                cMax++;
            }
            if (nums[node] > nums[cMax]) {
                swap(node, cMax);
                let lchild, rchild;
                if (cMax * 2 + 1 < tail) {
                    lchild = cMax * 2 + 1;
                    rchild = cMax * 2 + 2 < tail ? Infinity : nums[cMax * 2 + 2];
                    if (nums[cMax] < Math.max(nums[lchild], nums[rchild])) {
                        adjustHeap(cMax);
                    }
                }

            }
        }
    }

    //交换ij
    function swap(i, j) {
        nums[i] = nums[j] ^ nums[i];
        nums[j] = nums[i] ^ nums[j];
        nums[i] = nums[i] ^ nums[j];
    }

}




kthLargestElement(3,
    [9,2,2,4,2]);