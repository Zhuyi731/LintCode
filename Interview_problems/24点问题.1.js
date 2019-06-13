// 1-10 4种花色的牌 
// 随机抽取4张 算24点（加减乘除小括号），能得出24的概率是多少？
//  附加题：其中只能用一种算式计算出24的 有多少个？


// 令 四张牌分别为  A1 A2 A3 A4  3个算符为x1,x2,x3则算数方法有如下几种   
// A1 x1 A2 x2 A3 x3 A4  //不带括号的情况  
// (A1xA2)xA3xA4  A1x(A2xA3)xA4  A1xA2x(A3xA4)  (A1xA2xA3)xA4     //一个括号的情况
// (A1xA2)x(A3xA4) ((A1xA2)xA3)xA4 (A1x(A2xA3))xA4  A1x((A2xA3)xA4) A1x(A2x(A3xA4)) //两个的情况
const fs = require('fs')

let allSolutions = []

get24Possibility()
function get24Possibility () {
    let i1, i2, i3, i4
    let possibleTypes = []
    let impossibleTypes = []
    let l1 = 0, l2 = 0
    let temp
    for (i1 = 1; i1 < 11; i1++) {
        for (i2 = 1; i2 < 11; i2++) {
            for (i3 = 1; i3 < 11; i3++) {
                console.log(`已计算${i1 - 1}${i2 - 1}${i3 - 1}0个`)
                for (i4 = 1; i4 < 11; i4++) {
                    temp = isPossible(i1, i2, i3, i4)
                    if (temp.length > 0) {
                        possibleTypes.push(temp)
                    } else {
                        impossibleTypes.push([i1, i2, i3, i4])
                    }
                }
            }
            fs.appendFileSync("./possible.txt", possibleTypes.map(el => {
                return "组合为:" + el[0].type.join(" ") + "\r\n" + "解法如下:\r\n" + el.map(el => el.str).join("\r\n") + "\r\n"
            }).join("\r\n"), "utf-8")
            fs.appendFileSync("./impossible.txt", impossibleTypes.map(el => el.join(" ")).join("\r\n") + "\r\n", "utf-8")
            l1 += possibleTypes.length
            l2 += impossibleTypes.length
            possibleTypes = []
            impossibleTypes = []
        }
    }
    console.log(l1)
    console.log(l2)
}

function isPossible (i1, i2, i3, i4) {
    let x1, x2, x3
    let operatorMap = ["+", "-", "*", "/"]
    let result
    let ret = []
    // 1 2 3 4 + - * /
    for (x1 = 0; x1 < 4; x1++) {
        for (x2 = 0; x2 < 4; x2++) {
            for (x3 = 0; x3 < 4; x3++) {
                let all = getFullPermutation([i1, i2, i3, i4])
                all.forEach(el => {
                    i1 = el[0]
                    i2 = el[1]
                    i3 = el[2]
                    i4 = el[3]

                    result = eval(`${i1}${operatorMap[x1]}${i2}${operatorMap[x2]}${i3}${operatorMap[x3]}${i4}`)
                    if (result === 24) {
                        ret.push({
                            type: [i1, i2, i3, i4],
                            str: `${i1}${operatorMap[x1]}${i2}${operatorMap[x2]}${i3}${operatorMap[x3]}${i4}`
                        })
                    }
                    //处理括号情况
                    let bracketsTypes = [
                        `(${i1}${operatorMap[x1]}${i2})${operatorMap[x2]}${i3}${operatorMap[x3]}${i4}`,//(A1xA2)xA3xA4
                        `${i1}${operatorMap[x1]}(${i2}${operatorMap[x2]}${i3})${operatorMap[x3]}${i4}`,//A1x(A2xA3)xA4
                        `${i1}${operatorMap[x1]}${i2}${operatorMap[x2]}(${i3}${operatorMap[x3]}${i4})`,// A1xA2x(A3xA4) 
                        `(${i1}${operatorMap[x1]}${i2}${operatorMap[x2]}${i3})${operatorMap[x3]}${i4}`,//(A1xA2xA3)xA4
                        `${i1}${operatorMap[x1]}(${i2}${operatorMap[x2]}${i3}${operatorMap[x3]}${i4})`,//(A1xA2xA3)xA4
                        `(${i1}${operatorMap[x1]}${i2})${operatorMap[x2]}(${i3}${operatorMap[x3]}${i4})`,//(A1xA2)x(A3xA4)   
                        `((${i1}${operatorMap[x1]}${i2})${operatorMap[x2]}${i3})${operatorMap[x3]}${i4}`,//((A1xA2)xA3)xA4
                        `(${i1}${operatorMap[x1]}(${i2}${operatorMap[x2]}${i3}))${operatorMap[x3]}${i4}`,//(A1x(A2xA3))xA4
                        `${i1}${operatorMap[x1]}((${i2}${operatorMap[x2]}${i3})${operatorMap[x3]}${i4})`,// A1x((A2xA3)xA4)
                        `${i1}${operatorMap[x1]}(${i2}${operatorMap[x2]}(${i3}${operatorMap[x3]}${i4}))`// A1x(A2x(A3xA4)) 
                    ]
                    //括号的情况
                    // let shouldRunBrackets = [
                    //     [0, 1].includes(x1) && [2, 3].includes(x2),
                    //     [2, 3].includes(x3) && !([2, 3].includes(x1) && [2, 3].includes(x2)),
                    //     ([0, 1].includes(x2) && [0, 1].includes(x3) && [2, 3].includes(x2)),
                    //     // [[0,1].includes()]
                    // ]

                    bracketsTypes.forEach((el, index) => {
                        // if (!shouldRunBrackets[index]) return
                        result = eval(el)
                        if (result === 24) {
                            ret.push({
                                type: [i1, i2, i3, i4],
                                str: el
                            })
                        }
                    })
                })
            }
        }
    }
    return ret
}


function getFullPermutation (original) {
    let i1, i2, i3
    let all = []
    for (i1 = 0; i1 < 4; i1++) {
        for (i2 = 0; i2 < 4; i2++) {
            if (i2 === i1) continue
            for (i3 = 0; i3 < 4; i3++) {
                if (i3 === i2 || i3 === i1) continue
                i4 = [0, 1, 2, 3].filter(el => { return (el != i1 && el != i2 && el != i3) })[0]
                all.push([original[i1], original[i2], original[i3], original[i4]])
            }
        }
    }
    strAll = all.map(el => el.join(""))
    for (let i = 0; i < all.length; i++) {
        let tpStrAll = strAll.filter(el => el != all[i].join(""))
        if (tpStrAll.length < all.length - 1) {
            all.splice(i, 1)
            strAll.splice(i, 1)
            i--
        }
    }
    return all
}