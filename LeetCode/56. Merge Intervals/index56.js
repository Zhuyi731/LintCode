/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function (intervals) {
    intervals.sort((a,b)=>{
        if(a[0] > b[0]){
            return 1
        }else{
            return -1
        }
    })
    for (let i = 1; i < intervals.length; i++) {
        if (intervals[i - 1][1] >= intervals[i][0]) {
            intervals[i - 1][1] = Math.max(intervals[i][1], intervals[i - 1][1])
            intervals.splice(i, 1)
            i--
        }
    }
    return intervals
};