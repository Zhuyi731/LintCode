package main

func main() {

}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	mapper := make(map[int]*Employee, len(employees))
	for _, v := range employees {
		mapper[v.Id] = v
	}
	result := 0
	var dfs func(i int)
	dfs = func(i int) {
		result += mapper[i].Importance
		for _, v := range mapper[i].Subordinates {
			dfs(v)
		}
	}
	dfs(id)
	return result
}
