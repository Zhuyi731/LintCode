/**
 * @param {number[]} nums
 * @return {boolean}
 */
var judgePoint24 = function (nums) {
  all = getFullPermutation(nums)
  for (let i = 0; i < all.length; i++) {
    const result = isPossible(...all[i])
    if (result) {
      console.log(true)
      return true
    }
  }
  //   const result = isPossible(...nums)
  console.log(false)
  return false
}

function getFullPermutation(original) {
  let i1, i2, i3
  let all = []
  for (i1 = 0; i1 < 4; i1++) {
    for (i2 = 0; i2 < 4; i2++) {
      if (i2 === i1) continue
      for (i3 = 0; i3 < 4; i3++) {
        if (i3 === i2 || i3 === i1) continue
        i4 = [0, 1, 2, 3].filter((el) => {
          return el != i1 && el != i2 && el != i3
        })[0]
        all.push([original[i1], original[i2], original[i3], original[i4]])
      }
    }
  }
  strAll = all.map((el) => el.join(""))
  for (let i = 0; i < all.length; i++) {
    let tpStrAll = strAll.filter((el) => el != all[i].join(""))
    if (tpStrAll.length < all.length - 1) {
      all.splice(i, 1)
      strAll.splice(i, 1)
      i--
    }
  }
  return all
}

function GetReversePolishNotation(expression) {
  //根据逆波兰表达式来计算值
  // let result
  let numberStack = []
  let num1, num2, curNum, temp
  expression.forEach((el) => {
    if (typeof el !== "number") {
      //表达式
      num2 = numberStack.pop()
      num1 = numberStack.pop()
      switch (el) {
        case "+":
          {
            curNum = num1 + num2
          }
          break
        case "-":
          {
            curNum = num1 - num2
          }
          break
        case "*":
          {
            curNum = num1 * num2
          }
          break
        case "/":
          {
            curNum = num1 / num2
          }
          break
      }
      numberStack.push(curNum)
    } else {
      numberStack.push(el)
    }
  })

  return numberStack[0]
}

function isPossible(i1, i2, i3, i4) {
  let x1, x2, x3
  let _x1, _x2, _x3
  let operatorMap = ["+", "-", "*", "/"]
  let ret = []
  for (x1 = 0; x1 < 4; x1++) {
    for (x2 = 0; x2 < 4; x2++) {
      for (x3 = 0; x3 < 4; x3++) {
        _x1 = operatorMap[x1]
        _x2 = operatorMap[x2]
        _x3 = operatorMap[x3]
        let allSituation = [
          [i1, i2, i3, i4, _x1, _x2, _x3],
          [i1, i2, i3, _x1, i4, _x2, _x3],
          [i1, i2, i3, _x1, _x2, i4, _x3],
          [i1, i2, _x1, i3, i4, _x2, _x3],
          [i1, i2, _x1, i3, _x2, i4, _x3],
        ]
        for (let i = 0; i < allSituation.length; i++) {
          // 计算数据stack和operator stack
          let calculate = GetReversePolishNotation(allSituation[i])
          //   console.log(calculate[0])
          if (Math.abs(calculate - 24) < 0.0002) {
            //设置一个阈值  避免浮点数计算不精确的问题
            return true
          }
        }
      }
    }
  }
  return false
}

// GetReversePolishNotation([8, 4, "*", 7, 1, "-", "-"])
judgePoint24([4, 1, 8, 7])
judgePoint24([1, 2, 1, 2])
judgePoint24([1, 5, 5, 5])
