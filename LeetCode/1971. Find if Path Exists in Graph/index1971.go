package main

import "fmt"

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
	fmt.Println(validPath(1, [][]int{}, 0, 0))
}

func validPath(n int, edges [][]int, start int, end int) bool {
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

	hash := make(map[int]bool)
	exec := make(map[int]bool)
	toExec := make(map[int]bool)
	net := make([][]int, n)
	for k,v := range ededges {
		if()
	}


}
