/**
 * @param {number[][]} graph
 * @return {boolean}
 */
var isBipartite = function (graph) {
    const colors = [];
  
    const depRender = (index, color, depth) => {
      let curRelation = graph[index];
      for (let j = 0; j < curRelation.length; j++) {
        let optColor = (color + 1) % 2;
        if (colors[curRelation[j]] === color) {
          // 说明冲突了
          debugger;
          return false;
        } else if (colors[curRelation[j]] === optColor) {
          // 说明渲染过了
          continue;
        } else {
          colors[curRelation[j]] = optColor;
          if (!depRender(curRelation[j], colors[curRelation[j]], depth + 1)) {
            return false;
          }
        }
      }
      return true;
    };
  
    for (let i = 0; i < graph.length; i++) {
      if (colors[i] === undefined) {
        // 深搜
        colors[i] = 0;
        if (depRender(i, 0, 0) === false) {
          console.log(false);
          return false;
        }
      }
    }
  
    console.log(true);
    return true;
  };
  // [ [1, 3], [0, 2], [1, 3], [0, 2]]\n[[3], [2, 4], [1], [0, 4], [1, 3]]\n[ [1, 2, 3], [0, 2], [0, 1, 3], [0, 2]]
  isBipartite([
    [1, 3],
    [0, 2],
    [1, 3],
    [0, 2],
  ]);
  
  isBipartite([
    [3, 4, 6], //0
    [3, 6], //1
    [3, 6], //2
    [0, 1, 2, 5], //3
    [0, 7, 8], //4
    [3], //5
    [0, 1, 2, 7], //6
    [4, 6], //7
    [4], // 8
    [], // 9
  ]);
  
  isBipartite([[3], [2, 4], [1], [0, 4], [1, 3]]);
  
  
  
  isBipartite([
    [1, 2, 3],
    [0, 2],
    [0, 1, 3],
    [0, 2],
  ]);
  