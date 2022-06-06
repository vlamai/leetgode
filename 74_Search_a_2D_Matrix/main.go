package _4_Search_a_2D_Matrix

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		if row[len(row)-1] >= target {
			for _, cell := range row {
				if cell == target {
					return true
				}
			}
		}
	}
	return false
}
