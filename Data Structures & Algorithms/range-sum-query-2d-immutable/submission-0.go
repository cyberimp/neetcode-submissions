type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n, m:= len(matrix), len(matrix[0])
	prefix := make([][]int, n+1)
	for i:=0; i <= n; i++ {
		prefix[i] = make([]int, m+1)
	}
	for i:=1; i <= n; i++ {
		for j:=1; j <= m; j++ {
			prefix[i][j] = matrix[i-1][j-1] + prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1]
		}
	}
	return NumMatrix{prefix: prefix}
}

func (this *NumMatrix) SumRegion(r1 int, c1 int, r2 int, c2 int) int {
	return this.prefix[r2+1][c2+1] - this.prefix[r1][c2+1] - this.prefix[r2+1][c1] + this.prefix[r1][c1]
}

// Your NumMatrix object will be instantiated and called as such:
// obj := Constructor(matrix)
// param_1 := obj.SumRegion(row1,col1,row2,col2)
