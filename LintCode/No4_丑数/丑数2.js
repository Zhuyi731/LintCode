const nthUglyNumber = function (n) {
    let arr235 = [1,0,0],
        i;

    let uglyArr = [1, 2, 3, 4, 5];
    for (i = 5; i < n; i++) {
        let cc = getMin(uglyArr[arr235[0] + 1] * 2, uglyArr[arr235[1] + 1] * 3, uglyArr[arr235[2] + 1] * 5);
        gain(cc,i);
    }

    function gain(cc,index){
        uglyArr[index] = cc.ugly;
        cc.num.forEach(function(element) {
            arr235[element]++;
        }, this);
    }

    function getMin(a1, a2, a3) {
        let ret = {};
        if (a1 <= a2 && a1 <= a3) {
            if (a1 == a2) {
                ret.num = [0,1];
                ret.ugly = a1;
            } else if (a1 == a3) {
                ret.num = [0,2];
                ret.ugly = a1;
            } else {
                ret.num = [0];
                ret.ugly = a1;
            }
        }
        if (a2 <= a1 && a2 <= a3){
            if (a2 == a1) {
                ret.num = [0,1];
                ret.ugly = a2;
            } else if (a2 == a3) {
                ret.num = [1,2];
                ret.ugly = a2;
            } else {
                ret.num = [1];
                ret.ugly = a2;
            }
        }
        if(a3<= a1 && a3 <= a2){
            if (a3 == a1) {
                ret.num = [0,2];
                ret.ugly = a3;
            } else if (a3 == a2) {
                ret.num = [1,2];
                ret.ugly = a3;
            } else {
                ret.num = [2];
                ret.ugly = a3;
            }
        }
        if(a1 == a2 && a2 == a3){
           ret = {
                num:[0,1,2],
                ugly:a1
            }
        }
        return ret;


    }
    return uglyArr[n-1];

}
console.log(nthUglyNumber(10));