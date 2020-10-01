/**
 * @param {string} leaves
 * @return {number}
 */
function isRed(color) {
    return color === 'r' ? 1 : 0
}
function isYellow(color) {
    return color === 'y' ? 1 : 0
}
var minimumOperations = function (leaves) {
    let dp = [
        [isYellow(leaves[0]),
        isYellow(leaves[0]),
        isYellow(leaves[0])]
    ]
    dp[1] = [
        isYellow(leaves[0]) + isYellow(leaves[1]),
        isYellow(leaves[0]) + isRed(leaves[1]),
        isYellow(leaves[0]) + isRed(leaves[1])
    ]

    dp[2] = [
        dp[1][0] + isYellow(leaves[2]),
        Math.min(dp[1][0] + isRed(leaves[2]), dp[1][1] + isRed(leaves[2])),
        Math.min(dp[1][1] + isYellow(leaves[2]), dp[1][2] + isYellow(leaves[2])),
    ]
    for (let i = 3; i < leaves.length; i++) {
        let color = leaves[i]
        dp[i] = [
            dp[i - 1][0] + isYellow(color),
            Math.min(dp[i - 1][0] + isRed(color), dp[i - 1][1] + isRed(color)),
            Math.min(dp[i - 1][1] + isYellow(color), dp[i - 1][2] + isYellow(color)),
        ]
    }
    return dp[leaves.length - 1][2]
};


console.log(minimumOperations('yry'))
console.log(minimumOperations('rrryyyrryyyrr'))
console.log(minimumOperations('ryr'))