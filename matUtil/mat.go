package matUtil

func MatMul(a1 [][]float64, a2 [][]float64) [][]float64 {
	if a1 == nil || a2 == nil {
		panic("input cannt be nil.")
	}
	if len(a1) == 0 || len(a2) == 0 || len(a1[0]) == 0 || len(a2[0]) == 0 {
		panic("input must has at least one value.")
	}
	var a3 [][]float64
	for i, _ := range a1 {
		var a4 []float64
		for k, _ := range a2[0] {
			sum := float64(0.0)
			for j, _ := range a2 {
				sum = sum + a1[i][j]*a2[j][k]
			}
			a4 = append(a4, sum)
		}
		a3 = append(a3, a4)
	}
	return a3
}
func AscendDimesion(input []float64, axis int) [][]float64 {
	var result [][]float64
	if axis == 0 {
		result = append(result, input)
	}
	if axis == 1 {
		for _, v := range input {
			var result1 []float64
			result = append(result, append(result1, v))
		}
	}
	return result
}
