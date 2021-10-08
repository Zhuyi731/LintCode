/**
 * @param {number} n
 * @param {number[][]} edges
 * @param {number} start
 * @param {number} end
 * @return {boolean}
 */
 var validPath = function (n, edges, start, end) {
    const hash = { [start]: true }
    const execed = new Set()
    const toExec = [start]
    const net = []
    edges.forEach(el=>{
        !net[el[0]]?net[el[0]]=[el[1]]:net[el[0]].push(el[1])
        !net[el[1]]?net[el[1]]=[el[0]]:net[el[1]].push(el[0])
    })
    while(toExec.length){
        const e = toExec.shift()
        execed.add(e)
            net[e] && net[e].forEach(el=>{
            hash[el] = true 
            if(!execed.has(el) && !toExec.includes(el)){
                toExec.push(el)
            }
        })
    }
    return end in hash
};

console.log(validPath(1,[],0,0))
console.log(validPath(6,[[0,1],[0,2],[3,5],[5,4],[4,3]],0,5))
console.log(validPath(3,[[0,1],[1,2],[2,0]],0,2))
