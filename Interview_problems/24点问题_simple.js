// 1-10 4种花色的牌 
// 随机抽取4张 算24点（加减乘除小括号），能得出24的概率是多少？
//  附加题：其中只能用一种算式计算出24的 有多少个？

//采用逆波兰表达式来计算   
//逆波兰表达式在4个数的时候仅5种
// aaaa+++  
// aaa+a++
// aaa++a+
// aa+aa++
// aa+a+a+
let startTime = Date.now()
get24Possibility()
let endTime = Date.now()

console.log(`耗时${endTime - startTime}ms`)

function get24Possibility () {
    let i1, i2, i3, i4, min
    let solutions = {}

    let temp
    for (i1 = 1; i1 < 11; i1++) {
        // console.log(`已计算${i1 - 1}000个`)
        for (i2 = 1; i2 < 11; i2++) {
            for (i3 = 1; i3 < 11; i3++) {
                // console.log(`已计算${i1 - 1}${i2 - 1}${i3 - 1}0个`)
                for (i4 = 1; i4 < 11; i4++) {
                    min = [i1, i2, i3, i4].map(el => {
                        return el == 10 ? 0 : el
                    }).sort().join('')  //取得最小值做为hash   

                    temp = isPossible(i1, i2, i3, i4)

                    if (temp.length > 0) {
                        if (solutions[min]) {
                            solutions[min].ct += temp.length
                            solutions[min].expression = solutions[min].expression.concat(temp)
                        } else {
                            solutions[min] = {
                                ct: temp.length,
                                expression: temp
                            }
                        }
                    } else {
                        if (!solutions[min]) {
                            solutions[min] = {
                                ct: 0,
                                expression: []
                            }
                        }
                    }
                }
            }

        }
    }


    let possibles = [], impossibles = []
    for (let prop in solutions) {
        if (solutions[prop].ct === 0) {
            impossibles.push(prop)
        } else {
            possibles.push(prop)
        }
    }
    console.log('所有组合数' + (possibles.length + impossibles.length))
    console.log('可计算的24点个数' + possibles.length)
    console.log('不可计算的24点个数' + impossibles.length)
    console.log('概率' + (possibles.length * 100 / (possibles.length + impossibles.length)).toFixed(2) + '%')
}

function filterDuplicate (expressions) {
    let len = expressions.length
    let i, j
    for (i = 0; i < len - 1; i++) {
        for (j = 0; j < len; j++) {
            if (expressions[i] === expressions[j]) {
                expressions.splice(j, 1)
                len--
            }
        }
    }
    return expressions
}

function GetReversePolishNotation (expression) {
    //根据逆波兰表达式来计算值
    let numberStack = []
    let num1, num2, curNum, temp
    expression.forEach(el => {
        if (typeof el !== 'number') {//表达式
            num1 = numberStack.pop()
            num2 = numberStack.pop()
            switch (el) {
                case "+": {
                    curNum = num1 + num2
                } break;
                case '-': {
                    curNum = num1 - num2
                } break;
                case '*': {
                    curNum = num1 * num2
                } break;
                case '/': {
                    curNum = num1 / num2
                } break;
            }
            numberStack.push(curNum)
        } else {
            numberStack.push(el)
        }
    })

    return numberStack[0]
}

function isPossible (i1, i2, i3, i4) {
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
                    [i1, i2, _x1, i3, _x2, i4, _x3]
                ]
                allSituation.forEach(el => {
                    // 计算数据stack和operator stack
                    let calculate = GetReversePolishNotation(el)
                    if (Math.abs(calculate - 24) < 0.0002) {
                        ret.push(el)
                    }
                })
            }
        }
    }
    return ret
}