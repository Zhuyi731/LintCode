/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {
    let l = 0,r = Math.floor((x+1)/2), m= Math.floor((l+r+1)/2)

    while(l<r-1){
        if(m*m>x){
            r = m
            m = Math.floor((l+r)/2) 
        }else if(m*m < x){
            l = m 
            m = Math.floor((l+r)/2) 
        } else {
            l = m 
            break
        }
    }
    if(r*r<=x){
        l=r
    }
    console.log(l,r,m)
    return l
};
mySqrt(4)
mySqrt(8)
// Implement int sqrt(int x).

// Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

// Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

// Example 1:

// Input: 4
// Output: 2
// Example 2:

// Input: 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since 
//              the decimal part is truncated, 2 is returned.