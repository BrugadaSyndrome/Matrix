package matrix

import (
	"testing"
)

// Helper function to compare two matrices
func matricesEqual(t *testing.T, expected, actual *Matrix) bool {
	if expected == nil && actual == nil {
		return true
	}

	if expected == nil || actual == nil {
		t.Errorf("One matrix is nil and the other is not")
		return false
	}

	if expected.rows != actual.rows || expected.columns != actual.columns {
		t.Errorf("Matrix dimensions don't match: expected %dx%d, got %dx%d",
			expected.rows, expected.columns, actual.rows, actual.columns)
		return false
	}

	for i := 0; i < expected.rows; i++ {
		for j := 0; j < expected.columns; j++ {
			if expected.values[i][j] != actual.values[i][j] {
				t.Errorf("Matrix values don't match at position [%d][%d]: expected %f, got %f",
					i, j, expected.values[i][j], actual.values[i][j])
				return false
			}
		}
	}

	return true
}

// TestAddMatrices tests the AddMatrices function
func TestAddMatrices(t *testing.T) {
	// Test case 1: Adding two 2x2 matrices
	a1 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	b1 := NewMatrix(2, 2, [][]float64{
		{5, 6},
		{7, 8},
	})

	expected1 := &Matrix{
		rows:    2,
		columns: 2,
		values: [][]float64{
			{6, 8},
			{10, 12},
		},
	}

	result1 := AddMatrices(a1, b1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("AddMatrices failed for 2x2 matrices")
	}

	// Test case 2: Adding two 3x3 matrices
	a2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	b2 := NewMatrix(3, 3, [][]float64{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{10, 10, 10},
			{10, 10, 10},
			{10, 10, 10},
		},
	}

	result2 := AddMatrices(a2, b2)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("AddMatrices failed for 3x3 matrices")
	}

	// Test case 3: Adding matrices with different dimensions
	a3 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	b3 := NewMatrix(3, 3, [][]float64{
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
	})

	result3 := AddMatrices(a3, b3)
	if result3 != nil {
		t.Errorf("AddMatrices should return nil for matrices with different dimensions")
	}
}

// TestSubtractMatrices tests the SubtractMatrices function
func TestSubtractMatrices(t *testing.T) {
	// Test case 1: Subtracting two 2x2 matrices
	a1 := NewMatrix(2, 2, [][]float64{
		{10, 11},
		{12, 13},
	})

	b1 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	expected1 := &Matrix{
		rows:    2,
		columns: 2,
		values: [][]float64{
			{9, 9},
			{9, 9},
		},
	}

	result1 := SubtractMatrices(a1, b1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("SubtractMatrices failed for 2x2 matrices")
	}

	// Test case 2: Subtracting two 3x3 matrices
	a2 := NewMatrix(3, 3, [][]float64{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	})

	b2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{8, 6, 4},
			{2, 0, -2},
			{-4, -6, -8},
		},
	}

	result2 := SubtractMatrices(a2, b2)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("SubtractMatrices failed for 3x3 matrices")
	}

	// Test case 3: Subtracting matrices with different dimensions
	a3 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	b3 := NewMatrix(3, 3, [][]float64{
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
	})

	result3 := SubtractMatrices(a3, b3)
	if result3 != nil {
		t.Errorf("SubtractMatrices should return nil for matrices with different dimensions")
	}
}

// TestMultiplyMatrices tests the MultiplyMatrices function
func TestMultiplyMatrices(t *testing.T) {
	// Test case 1: Multiplying a 2x3 matrix by a 3x2 matrix
	a1 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	b1 := NewMatrix(3, 2, [][]float64{
		{7, 8},
		{9, 10},
		{11, 12},
	})

	expected1 := &Matrix{
		rows:    2,
		columns: 2,
		values: [][]float64{
			{58, 64},
			{139, 154},
		},
	}

	result1 := MultiplyMatrices(a1, b1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("MultiplyMatrices failed for 2x3 * 3x2 matrices")
	}

	// Test case 2: Multiplying two 2x2 matrices
	a2 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	b2 := NewMatrix(2, 2, [][]float64{
		{5, 6},
		{7, 8},
	})

	expected2 := &Matrix{
		rows:    2,
		columns: 2,
		values: [][]float64{
			{19, 22},
			{43, 50},
		},
	}

	result2 := MultiplyMatrices(a2, b2)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("MultiplyMatrices failed for 2x2 * 2x2 matrices")
	}

	// Test case 3: Multiplying matrices with incompatible dimensions
	a3 := NewMatrix(2, 2, [][]float64{
		{1, 2},
		{3, 4},
	})

	b3 := NewMatrix(3, 3, [][]float64{
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
	})

	result3 := MultiplyMatrices(a3, b3)
	if result3 != nil {
		t.Errorf("MultiplyMatrices should return nil for matrices with incompatible dimensions")
	}
}

// TestAppendRow tests the AppendRow function
func TestAppendRow(t *testing.T) {
	// Test case 1: Appending a row to a 2x3 matrix
	m1 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	newRow1 := []float64{7, 8, 9}

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	result1 := AppendRow(m1, newRow1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("AppendRow failed for 2x3 matrix with matching row length")
	}

	// Test case 2: Appending a row with incorrect length
	m2 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	newRow2 := []float64{7, 8} // Only 2 elements, but matrix has 3 columns

	result2 := AppendRow(m2, newRow2)
	if result2 != nil {
		t.Errorf("AppendRow should return nil when the new row length doesn't match the matrix columns")
	}
}

// TestAppendColumn tests the AppendColumn function
func TestAppendColumn(t *testing.T) {
	// Test case 1: Appending a column to a 2x3 matrix
	m1 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	newCol1 := []float64{7, 8}

	expected1 := &Matrix{
		rows:    2,
		columns: 4,
		values: [][]float64{
			{1, 2, 3, 7},
			{4, 5, 6, 8},
		},
	}

	result1 := AppendColumn(m1, newCol1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("AppendColumn failed for 2x3 matrix with matching column length")
	}

	// Test case 2: Appending a column with incorrect length
	m2 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	newCol2 := []float64{7, 8, 9} // 3 elements, but matrix has 2 rows

	result2 := AppendColumn(m2, newCol2)
	if result2 != nil {
		t.Errorf("AppendColumn should return nil when the new column length doesn't match the matrix rows")
	}
}

// TestSwapRows tests the SwapRows function
func TestSwapRows(t *testing.T) {
	// Test case 1: Swapping rows in a 3x3 matrix
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{7, 8, 9},
			{4, 5, 6},
			{1, 2, 3},
		},
	}

	result1 := SwapRows(m1, 0, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("SwapRows failed for 3x3 matrix with valid row indices")
	}

	// Test case 2: Swapping rows with invalid indices
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative index
	result2 := SwapRows(m2, -1, 2)
	if result2 != nil {
		t.Errorf("SwapRows should return nil when given a negative row index")
	}

	// Test with out-of-bounds index
	result3 := SwapRows(m2, 0, 3)
	if result3 != nil {
		t.Errorf("SwapRows should return nil when given an out-of-bounds row index")
	}
}

// TestSwapColumns tests the SwapColumns function
func TestSwapColumns(t *testing.T) {
	// Test case 1: Swapping columns in a 3x3 matrix
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{3, 2, 1},
			{6, 5, 4},
			{9, 8, 7},
		},
	}

	result1 := SwapColumns(m1, 0, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("SwapColumns failed for 3x3 matrix with valid column indices")
	}

	// Test case 2: Swapping columns with invalid indices
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative index
	result2 := SwapColumns(m2, -1, 2)
	if result2 != nil {
		t.Errorf("SwapColumns should return nil when given a negative column index")
	}

	// Test with out-of-bounds index
	result3 := SwapColumns(m2, 0, 3)
	if result3 != nil {
		t.Errorf("SwapColumns should return nil when given an out-of-bounds column index")
	}
}

// TestMultiplyRow tests the MultiplyRow function
func TestMultiplyRow(t *testing.T) {
	// Test case 1: Multiplying a row by a positive scalar
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 2, 3},
			{8, 10, 12}, // Row 1 multiplied by 2
			{7, 8, 9},
		},
	}

	result1 := MultiplyRow(m1, 1, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("MultiplyRow failed for 3x3 matrix with positive scalar")
	}

	// Test case 2: Multiplying a row by a negative scalar
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{-7, -8, -9}, // Row 2 multiplied by -1
		},
	}

	result2 := MultiplyRow(m2, 2, -1)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("MultiplyRow failed for 3x3 matrix with negative scalar")
	}

	// Test case 3: Multiplying a row with invalid index
	m3 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative index
	result3 := MultiplyRow(m3, -1, 2)
	if result3 != nil {
		t.Errorf("MultiplyRow should return nil when given a negative row index")
	}

	// Test with out-of-bounds index
	result4 := MultiplyRow(m3, 3, 2)
	if result4 != nil {
		t.Errorf("MultiplyRow should return nil when given an out-of-bounds row index")
	}
}

// TestMultiplyColumn tests the MultiplyColumn function
func TestMultiplyColumn(t *testing.T) {
	// Test case 1: Multiplying a column by a positive scalar
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 4, 3}, // Column 1 multiplied by 2
			{4, 10, 6},
			{7, 16, 9},
		},
	}

	result1 := MultiplyColumn(m1, 1, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("MultiplyColumn failed for 3x3 matrix with positive scalar")
	}

	// Test case 2: Multiplying a column by a negative scalar
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 2, -3}, // Column 2 multiplied by -1
			{4, 5, -6},
			{7, 8, -9},
		},
	}

	result2 := MultiplyColumn(m2, 2, -1)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("MultiplyColumn failed for 3x3 matrix with negative scalar")
	}

	// Test case 3: Multiplying a column with invalid index
	m3 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative index
	result3 := MultiplyColumn(m3, -1, 2)
	if result3 != nil {
		t.Errorf("MultiplyColumn should return nil when given a negative column index")
	}

	// Test with out-of-bounds index
	result4 := MultiplyColumn(m3, 3, 2)
	if result4 != nil {
		t.Errorf("MultiplyColumn should return nil when given an out-of-bounds column index")
	}
}

// TestAddScaledRow tests the AddScaledRow function
func TestAddScaledRow(t *testing.T) {
	// Test case 1: Adding a scaled row to another row with positive scalar
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{15, 18, 21}, // Row 0 + 2*Row 2
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	result1 := AddScaledRow(m1, 0, 2, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("AddScaledRow failed for 3x3 matrix with positive scalar")
	}

	// Test case 2: Adding a scaled row to another row with negative scalar
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 2, 3},
			{0, 0, 0}, // Row 1 + (-1)*Row 1
			{7, 8, 9},
		},
	}

	result2 := AddScaledRow(m2, 1, 1, -1)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("AddScaledRow failed for 3x3 matrix with negative scalar")
	}

	// Test case 3: Adding a scaled row with invalid target index
	m3 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative target index
	result3 := AddScaledRow(m3, -1, 1, 2)
	if result3 != nil {
		t.Errorf("AddScaledRow should return nil when given a negative target row index")
	}

	// Test with out-of-bounds target index
	result4 := AddScaledRow(m3, 3, 1, 2)
	if result4 != nil {
		t.Errorf("AddScaledRow should return nil when given an out-of-bounds target row index")
	}

	// Test case 4: Adding a scaled row with invalid source index
	m4 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative source index
	result5 := AddScaledRow(m4, 1, -1, 2)
	if result5 != nil {
		t.Errorf("AddScaledRow should return nil when given a negative source row index")
	}

	// Test with out-of-bounds source index
	result6 := AddScaledRow(m4, 1, 3, 2)
	if result6 != nil {
		t.Errorf("AddScaledRow should return nil when given an out-of-bounds source row index")
	}
}

// TestAddScaledColumn tests the AddScaledColumn function
func TestAddScaledColumn(t *testing.T) {
	// Test case 1: Adding a scaled column to another column with positive scalar
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{7, 2, 3}, // Column 0 + 2*Column 2
			{16, 5, 6},
			{25, 8, 9},
		},
	}

	result1 := AddScaledColumn(m1, 0, 2, 2)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("AddScaledColumn failed for 3x3 matrix with positive scalar")
	}

	// Test case 2: Adding a scaled column to another column with negative scalar
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 0, 3}, // Column 1 + (-1)*Column 1
			{4, 0, 6},
			{7, 0, 9},
		},
	}

	result2 := AddScaledColumn(m2, 1, 1, -1)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("AddScaledColumn failed for 3x3 matrix with negative scalar")
	}

	// Test case 3: Adding a scaled column with invalid target index
	m3 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative target index
	result3 := AddScaledColumn(m3, -1, 1, 2)
	if result3 != nil {
		t.Errorf("AddScaledColumn should return nil when given a negative target column index")
	}

	// Test with out-of-bounds target index
	result4 := AddScaledColumn(m3, 3, 1, 2)
	if result4 != nil {
		t.Errorf("AddScaledColumn should return nil when given an out-of-bounds target column index")
	}

	// Test case 4: Adding a scaled column with invalid source index
	m4 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	// Test with negative source index
	result5 := AddScaledColumn(m4, 1, -1, 2)
	if result5 != nil {
		t.Errorf("AddScaledColumn should return nil when given a negative source column index")
	}

	// Test with out-of-bounds source index
	result6 := AddScaledColumn(m4, 1, 3, 2)
	if result6 != nil {
		t.Errorf("AddScaledColumn should return nil when given an out-of-bounds source column index")
	}
}

// TestTransposeMatrix tests the TransposeMatrix function
func TestTransposeMatrix(t *testing.T) {
	// Test case 1: Transposing a square matrix
	m1 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})

	expected1 := &Matrix{
		rows:    3,
		columns: 3,
		values: [][]float64{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	}

	result1 := TransposeMatrix(m1)
	if !matricesEqual(t, expected1, result1) {
		t.Errorf("TransposeMatrix failed for 3x3 square matrix")
	}

	// Test case 2: Transposing a non-square matrix
	m2 := NewMatrix(2, 3, [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})

	expected2 := &Matrix{
		rows:    3,
		columns: 2,
		values: [][]float64{
			{1, 4},
			{2, 5},
			{3, 6},
		},
	}

	result2 := TransposeMatrix(m2)
	if !matricesEqual(t, expected2, result2) {
		t.Errorf("TransposeMatrix failed for 2x3 non-square matrix")
	}
}

// TestRowEchelonForm tests the RowEchelonForm function
func TestRowEchelonForm(t *testing.T) {
	// Test case 1: Converting a matrix to row echelon form
	m1 := NewMatrix(3, 4, [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	})

	result1 := RowEchelonForm(m1)

	// Check if the result is in row echelon form
	// We need to allow for small floating-point differences
	for i := 0; i < result1.rows; i++ {
		// Check if there are only zeros to the left of the pivot
		for j := 0; j < i; j++ {
			if result1.values[i][j] != 0 {
				t.Errorf("RowEchelonForm failed: non-zero value at position [%d][%d]", i, j)
			}
		}

		// Check if the pivot is 1 (if it exists)
		pivotFound := false
		for j := i; j < result1.columns && !pivotFound; j++ {
			if result1.values[i][j] != 0 {
				if result1.values[i][j] != 1.0 {
					t.Errorf("RowEchelonForm failed: pivot at position [%d][%d] is not 1", i, j)
				}
				pivotFound = true
			}
		}
	}

	// Test case 2: Matrix with a zero row
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{0, 0, 0},
		{4, 5, 6},
	})

	result2 := RowEchelonForm(m2)

	// Check if the result is in row echelon form
	// Zero rows should be at the bottom
	zeroRowsAtBottom := true
	for i := 0; i < result2.rows-1; i++ {
		rowIsZero := true
		for j := 0; j < result2.columns; j++ {
			if result2.values[i][j] != 0 {
				rowIsZero = false
				break
			}
		}
		if rowIsZero {
			// Check if all subsequent rows are also zero
			for k := i + 1; k < result2.rows; k++ {
				for j := 0; j < result2.columns; j++ {
					if result2.values[k][j] != 0 {
						zeroRowsAtBottom = false
						break
					}
				}
				if !zeroRowsAtBottom {
					break
				}
			}
			if !zeroRowsAtBottom {
				t.Errorf("RowEchelonForm failed: zero rows are not at the bottom")
			}
			break
		}
	}
}

// TestReducedRowEchelonForm tests the ReducedRowEchelonForm function
func TestReducedRowEchelonForm(t *testing.T) {
	// Test case 1: Converting a matrix to reduced row echelon form
	m1 := NewMatrix(3, 4, [][]float64{
		{2, 1, -1, 8},
		{-3, -1, 2, -11},
		{-2, 1, 2, -3},
	})

	result1 := ReducedRowEchelonForm(m1)

	// Check if the result is in reduced row echelon form
	// 1. Each row has a pivot of 1
	// 2. Each pivot is the only non-zero entry in its column
	// 3. Pivots appear in a "staircase" pattern from left to right

	// Find the pivot positions
	pivotCols := make([]int, result1.rows)
	for i := 0; i < result1.rows; i++ {
		pivotCols[i] = -1 // Initialize to -1 (no pivot)
		for j := 0; j < result1.columns; j++ {
			if result1.values[i][j] != 0 {
				pivotCols[i] = j
				// Check if pivot is 1
				if result1.values[i][j] != 1.0 {
					t.Errorf("ReducedRowEchelonForm failed: pivot at position [%d][%d] is not 1", i, j)
				}
				break
			}
		}
	}

	// Check that pivots are in a staircase pattern
	for i := 0; i < result1.rows-1; i++ {
		if pivotCols[i] != -1 && pivotCols[i+1] != -1 && pivotCols[i] >= pivotCols[i+1] {
			t.Errorf("ReducedRowEchelonForm failed: pivots are not in a staircase pattern")
		}
	}

	// Check that each pivot is the only non-zero entry in its column
	for i := 0; i < result1.rows; i++ {
		if pivotCols[i] != -1 {
			for j := 0; j < result1.rows; j++ {
				if j != i && result1.values[j][pivotCols[i]] != 0 {
					t.Errorf("ReducedRowEchelonForm failed: non-zero entry at position [%d][%d] in pivot column", j, pivotCols[i])
				}
			}
		}
	}

	// Test case 2: Matrix with a zero row
	m2 := NewMatrix(3, 3, [][]float64{
		{1, 2, 3},
		{0, 0, 0},
		{4, 5, 6},
	})

	result2 := ReducedRowEchelonForm(m2)

	// Check if the result is in reduced row echelon form
	// Zero rows should be at the bottom
	zeroRowsAtBottom := true
	for i := 0; i < result2.rows-1; i++ {
		rowIsZero := true
		for j := 0; j < result2.columns; j++ {
			if result2.values[i][j] != 0 {
				rowIsZero = false
				break
			}
		}
		if rowIsZero {
			// Check if all subsequent rows are also zero
			for k := i + 1; k < result2.rows; k++ {
				for j := 0; j < result2.columns; j++ {
					if result2.values[k][j] != 0 {
						zeroRowsAtBottom = false
						break
					}
				}
				if !zeroRowsAtBottom {
					break
				}
			}
			if !zeroRowsAtBottom {
				t.Errorf("ReducedRowEchelonForm failed: zero rows are not at the bottom")
			}
			break
		}
	}
}
