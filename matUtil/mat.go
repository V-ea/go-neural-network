package matUtil

func MatMul(a1 [][]float32, a2 [][]float32) [][]float32 {
	if a1 == nil || a2 == nil {
		panic("input cannt be nil.")
	}
	if len(a1) == 0 || len(a2) == 0 || len(a1[0]) == 0 || len(a2[0]) == 0 {
		panic("input must has at least one value.")
	}
	var a3 [][]float32
	for i, _ := range a1 {
		var a4 []float32
		for k, _ := range a2[0] {
			sum := float32(0.0)
			for j, _ := range a2 {
				sum = sum + a1[i][j]*a2[j][k]
			}
			a4 = append(a4, sum)
		}
		a3 = append(a3, a4)
	}
	return a3
}
