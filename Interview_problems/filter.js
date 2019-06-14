const fs = require("fs")

let pos1 = fs.readFileSync('./1.txt', 'utf-8')
let pos2 = fs.readFileSync('./3.txt', 'utf-8')


pos1 = pos1.split('\r\n').map(el => el.split(',').sort((a, b) => ~~a > ~~b).join(','))
pos2 = pos2.split('\n').filter(el => /算不起/.test(el)).map(el => el.split(' :')[0].split(' ').sort((a, b) => ~~a > ~~b).join(','))



let ret = []
pos1.forEach(el => {
    if (!pos2.includes(el)) {
        ret.push(el)
    }
});

console.log(ret)