/**
 * @param {number} x
 * @param {number} n
 * @return {number}
 */
var myPow = function (x, n) {
  if (n == 0) {
    return 1;
  }

  let originX = x;
  let act = "up";
  if (n < 0) {
    act = "down";
    n = -n;
  }

  let absn = Math.abs(n);
  let currentPlus = x;
  let result = 1;
  while (absn) {
    let temp = absn & 1;
    absn = Math.floor(absn / 2);
    if (temp) {
      result =   Number((result * currentPlus));
    }
    currentPlus = Number((currentPlus * currentPlus));
  }

  if (act === "down") {
    result = 1 / result 
  }

  result = Number(result.toFixed(5))
  console.log(result);
  return result;
};
myPow(0.25519 ,-6) // 3620.91299
myPow(2.0, 10); // 1024.00000
myPow(34.00515, -3); // 0.00003
myPow(8.88023, 3); //700.28148
myPow(2.1, 3); //  9.26100
myPow(2.0, -2); // 0.25000
