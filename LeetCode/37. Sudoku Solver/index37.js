/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */

const print = (pos, board) => {
  console.log("|---.---.---.---.---.---.---.---.---|");
  for (let i = 0; i < 9; i++) {
    let str = "|";
    for (let j = 0; j < 9; j++) {
      if (pos[i][j]) {
        if (pos[i][j].confirm) {
          str += ` ${pos[i][j].val} |`;
        } else {
          str += "   |";
        }
      } else {
        str += ` ${board[i][j]} |`;
      }
    }
    console.log(str);
  }
  console.log("|---.---.---.---.---.---.---.---.---|");
};

var solveSudoku = function (board) {
  let ok = false;
  let result = [];
  const getPossible = (x, y) => {
    // 计算列
    let defaultPossible = {
      1: 1,
      2: 1,
      3: 1,
      4: 1,
      5: 1,
      6: 1,
      7: 1,
      8: 1,
      9: 1,
    };
    for (let i = 0; i < 9; i++) {
      if (board[x][i] !== ".") {
        delete defaultPossible[Number(board[x][i])];
      }
    }
    for (let i = 0; i < 9; i++) {
      if (board[i][y] !== ".") {
        delete defaultPossible[Number(board[i][y])];
      }
    }
    let startX = x - (x % 3);
    let startY = y - (y % 3);
    for (let i = startX; i < startX + 3; i++) {
      for (let j = startY; j < startY + 3; j++) {
        if (board[i][j] !== ".") {
          delete defaultPossible[Number(board[i][j])];
        }
      }
    }
    return defaultPossible;
  };
  const travel = (recordCopy, xx, yy, depth = 0) => {
    xx !== undefined && console.log(xx, yy, recordCopy[xx][yy].val, depth);
    print(recordCopy, board);
    if (ok) {
      return;
    }
    // let recordCopy = record.map((el) => el);
    let currentMin = {
      ct: 100,
      x: -1,
      y: -1,
    };
    for (let i = 0; i < 9; i++) {
      for (let j = 0; j < 9; j++) {
        if (!recordCopy[i][j]) continue;
        if (!recordCopy[i][j].confirm) {
          // 尚未确定
          if (Object.keys(recordCopy[i][j].possible).length < currentMin.ct) {
            currentMin = {
              ct: Object.keys(recordCopy[i][j].possible).length,
              x: i,
              y: j,
            };
          }
        }
      }
    }
    // 最小值
    console.log(`最小值${currentMin.ct} @${currentMin.x}, ${currentMin.y}`);
    if (currentMin.ct === 0) {
      // 说明失败了
      console.log("回溯");
      return;
    } else if (currentMin.ct === 100) {
      // 说明已经找到了
      result = recordCopy;
      ok = true;
      return;
    } else {
      let i = currentMin.x;
      let j = currentMin.y;
      let possibleValues = Object.keys(recordCopy[i][j].possible);
      console.log("全部可能性" + possibleValues.join(","));
      possibleValues.forEach((el, index) => {
        if (ok) return;
        let bak = recordCopy.map((el) => {
          return el.map((a) => {
            if (a) {
              return {
                confirm: a.confirm,
                possible: Object.assign({}, a.possible),
                val: a.val,
              };
            }
          });
        });
        bak[i][j].confirm = true;
        bak[i][j].val = el;
        bak[i][j].possible = {
          [el]: 1,
        };
        deletePos(bak, i, j, el);
        // console.log(recordCopy[2][5].possible);
        travel(bak, i, j, depth + 1);
        console.log(`depth:${depth} index:${index}`);
      });
    }
  };
  const deletePos = (record, x, y, val) => {
    for (let i = 0; i < 9; i++) {
      if (!record[x][i]) continue;
      if (record[x][i].possible[val]) {
        delete record[x][i].possible[val];
      }
    }
    for (let i = 0; i < 9; i++) {
      if (!record[i][y]) continue;
      if (record[i][y].possible[val]) {
        delete record[i][y].possible[val];
      }
    }
    let startX = x - (x % 3);
    let startY = y - (y % 3);
    for (let i = startX; i < startX + 3; i++) {
      for (let j = startY; j < startY + 3; j++) {
        if (!record[i][j]) continue;
        if (record[i][j].possible[val]) {
          delete record[i][j].possible[val];
        }
      }
    }
  };
  let boardRecord = board.map((el, x) => {
    let newEl = el.map((item, y) => {
      if (item === ".") {
        let possible = getPossible(x, y);
        let keys = Object.keys(possible);
        if (keys.length === 1) {
          return {
            possible,
            confirm: true,
            val: keys[0],
          };
        } else {
          return {
            possible,
            confirm: false,
            val: "",
          };
        }
      }
    });
    return newEl;
  });

  travel(boardRecord);

  console.log("result");
  print(result, board);
  // 把result放入board
  result.forEach((row, x) => {
    row.forEach((line, y) => {
      if (line && line.confirm && board[x][y] === ".") {
        board[x][y] = line.val;
      }
    });
  });
  return board;
};
solveSudoku([
  [".", ".", "9", "7", "4", "8", ".", ".", "."],
  ["7", ".", ".", ".", ".", ".", ".", ".", "."],
  [".", "2", ".", "1", ".", "9", ".", ".", "."],
  [".", ".", "7", ".", ".", ".", "2", "4", "."],
  [".", "6", "4", ".", "1", ".", "5", "9", "."],
  [".", "9", "8", ".", ".", ".", "3", ".", "."],
  [".", ".", ".", "8", ".", "3", ".", "2", "."],
  [".", ".", ".", ".", ".", ".", ".", ".", "6"],
  [".", ".", ".", "2", "7", "5", "9", ".", "."],
]);
solveSudoku([
  ["5", "3", ".", ".", "7", ".", ".", ".", "."],
  ["6", ".", ".", "1", "9", "5", ".", ".", "."],
  [".", "9", "8", ".", ".", ".", ".", "6", "."],
  ["8", ".", ".", ".", "6", ".", ".", ".", "3"],
  ["4", ".", ".", "8", ".", "3", ".", ".", "1"],
  ["7", ".", ".", ".", "2", ".", ".", ".", "6"],
  [".", "6", ".", ".", ".", ".", "2", "8", "."],
  [".", ".", ".", "4", "1", "9", ".", ".", "5"],
  [".", ".", ".", ".", "8", ".", ".", "7", "9"],
]);
