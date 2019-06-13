const fs = require("fs")

possible = fs.readFileSync('./interview_problems/possible.txt', 'utf-8')
impossible = fs.readFileSync('./interview_problems/impossible.txt', 'utf-8')

possible = possible.split("\r\n").filter(el => /组合为/.test(el))
impossible = impossible.split("\r\n")

possible = possible.map(el => {
    match = /组合为:(.*)/.exec(el)
    tt = match[1].split(" ").map(el => ~~el)
    tt = tt.sort((e1,e2)=>e1>e2)
    return match[1].split(" ").map(el => ~~el).sort((e1,e2)=>e1>e2).map(el => {
        if (el == "10") {
            return 0
        } else {
            return el
        }
    }).join("")
})

impossible = impossible.map(el => {
    return el.split(" ").map(el => ~~el).sort().map(el => {
        if (el == "10") {
            return 0
        } else {
            return el
        }
    }).join("")
})

//然后去重
var len = possible.length
for (var i = 0; i < len - 1; i++) {
    for (var j = i + 1; j < len; j++) {
        if (possible[i] == possible[j]) {
            possible.splice(j, 1)
            len--
        }
    }
}

len = impossible.length
for (var i = 0; i < len - 1; i++) {
    for (var j = i + 1; j < len; j++) {
        if (impossible[i] == impossible[j]) {
            impossible.splice(j, 1)
            len--
        }
    }
}
console.log(possible.length)
console.log(impossible.length)
console.log(impossible.length / possible.length)