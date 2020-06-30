/**
 * @param {number} n
 * @return {string[][]}
 */
var solveNQueens = function (n) {
    let result = 0;
  
    dfs(0, n, [], []);
  
    function dfs(curDepth, depth, usedPos, curMatrix) {
      if (curDepth === depth) {
        return result++;
      }
      for (let i = 0; i < depth; i++) {
        let usedPos_bak = usedPos.map((el) => el);
        let curMatrix_bak = curMatrix.map((el) => el);
        if (usedPos_bak[i]) {
          continue;
        }
        if (checkCanAttack(curDepth, depth, i, curMatrix_bak)) {
          continue;
        }
        // 如果可行，则递归
        usedPos_bak[i] = 1;
        let tmp = Array(depth).fill(".");
        tmp[i] = "Q";
        curMatrix_bak.push(tmp);
        dfs(curDepth + 1, depth, usedPos_bak, curMatrix_bak);
      }
    }
    function checkCanAttack(curDepth, depth, i, curMatrix) {
      for (let j = 0; j < curDepth; j++) {
        let minus = curDepth - j;
        if (i - minus >= 0 && curMatrix[j][i - minus] === "Q") {
          return true;
        }
        if (i + minus < depth && curMatrix[j][i + minus] === "Q") {
          return true;
        }
      }
    }
    console.log(result);
    return result;
  };
  
  solveNQueens(4);
  // [[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
  