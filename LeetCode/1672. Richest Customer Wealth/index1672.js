/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function (accounts) {
  return accounts
    .map((account) =>  account.reduce((a, b) => a + b, 0))
    .reduce((a, b) => (a > b ? a : b), 0);
};
