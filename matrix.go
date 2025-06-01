package matrix

type Matrix struct {
	rows    int
	columns int
	values  [][]float64
}

func NewMatrix(rows, cols int, values [][]float64) Matrix {
	m := Matrix{
		rows:    rows,
		columns: cols,
		values:  make([][]float64, rows),
	}

	for i := 0; i < rows; i++ {
		m.values[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			m.values[i][j] = values[i][j]
		}
	}

	return m
}

// AddMatrices adds two matrices and returns a pointer to the resulting matrix.
// Returns nil if the matrices have different dimensions.
func AddMatrices(a, b Matrix) *Matrix {
	// Check if matrices have the same dimensions
	if a.rows != b.rows || a.columns != b.columns {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    a.rows,
		columns: a.columns,
		values:  make([][]float64, a.rows),
	}

	// Initialize the values slice and add corresponding elements
	for i := 0; i < a.rows; i++ {
		result.values[i] = make([]float64, a.columns)
		for j := 0; j < a.columns; j++ {
			result.values[i][j] = a.values[i][j] + b.values[i][j]
		}
	}

	return result
}

// SubtractMatrices subtracts the second matrix from the first and returns a pointer to the resulting matrix.
// Returns nil if the matrices have different dimensions.
func SubtractMatrices(a, b Matrix) *Matrix {
	// Check if matrices have the same dimensions
	if a.rows != b.rows || a.columns != b.columns {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    a.rows,
		columns: a.columns,
		values:  make([][]float64, a.rows),
	}

	// Initialize the values slice and subtract corresponding elements
	for i := 0; i < a.rows; i++ {
		result.values[i] = make([]float64, a.columns)
		for j := 0; j < a.columns; j++ {
			result.values[i][j] = a.values[i][j] - b.values[i][j]
		}
	}

	return result
}

// MultiplyMatrices multiplies two matrices and returns a pointer to the resulting matrix.
// Returns nil if the matrices cannot be multiplied (i.e., if the number of columns in the first matrix
// does not equal the number of rows in the second matrix).
func MultiplyMatrices(a, b Matrix) *Matrix {
	// Check if matrices can be multiplied
	if a.columns != b.rows {
		return nil
	}

	// Create a new matrix with dimensions (a.rows × b.columns)
	result := &Matrix{
		rows:    a.rows,
		columns: b.columns,
		values:  make([][]float64, a.rows),
	}

	// Initialize the values slice and perform matrix multiplication
	for i := 0; i < a.rows; i++ {
		result.values[i] = make([]float64, b.columns)
		for j := 0; j < b.columns; j++ {
			// Calculate the dot product of row i from matrix a and column j from matrix b
			sum := 0.0
			for k := 0; k < a.columns; k++ {
				sum += a.values[i][k] * b.values[k][j]
			}
			result.values[i][j] = sum
		}
	}

	return result
}

// AppendRow appends a new row to the matrix and returns a pointer to the resulting matrix.
// Returns nil if the length of the new row doesn't match the number of columns in the matrix.
func AppendRow(m Matrix, row []float64) *Matrix {
	// Check if the length of the new row matches the number of columns in the matrix
	if len(row) != m.columns {
		return nil
	}

	// Create a new matrix with dimensions (m.rows + 1) × m.columns
	result := &Matrix{
		rows:    m.rows + 1,
		columns: m.columns,
		values:  make([][]float64, m.rows+1),
	}

	// Copy values from the original matrix
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			result.values[i][j] = m.values[i][j]
		}
	}

	// Add the new row at the end
	result.values[m.rows] = make([]float64, m.columns)
	for j := 0; j < m.columns; j++ {
		result.values[m.rows][j] = row[j]
	}

	return result
}

// AppendColumn appends a new column to the matrix and returns a pointer to the resulting matrix.
// Returns nil if the length of the new column doesn't match the number of rows in the matrix.
func AppendColumn(m Matrix, column []float64) *Matrix {
	// Check if the length of the new column matches the number of rows in the matrix
	if len(column) != m.rows {
		return nil
	}

	// Create a new matrix with dimensions m.rows × (m.columns + 1)
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns + 1,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix and add the new column
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns+1)
		// Copy existing values
		for j := 0; j < m.columns; j++ {
			result.values[i][j] = m.values[i][j]
		}
		// Add the new column value at the end
		result.values[i][m.columns] = column[i]
	}

	return result
}

// SwapRows swaps two rows in the matrix and returns a pointer to the resulting matrix.
// Returns nil if either row index is out of bounds.
func SwapRows(m Matrix, row1, row2 int) *Matrix {
	// Check if row indices are valid
	if row1 < 0 || row1 >= m.rows || row2 < 0 || row2 >= m.rows {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix, swapping the specified rows
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)

		// Determine which row to copy from
		sourceRow := i
		if i == row1 {
			sourceRow = row2
		} else if i == row2 {
			sourceRow = row1
		}

		// Copy the row
		for j := 0; j < m.columns; j++ {
			result.values[i][j] = m.values[sourceRow][j]
		}
	}

	return result
}

// SwapColumns swaps two columns in the matrix and returns a pointer to the resulting matrix.
// Returns nil if either column index is out of bounds.
func SwapColumns(m Matrix, col1, col2 int) *Matrix {
	// Check if column indices are valid
	if col1 < 0 || col1 >= m.columns || col2 < 0 || col2 >= m.columns {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix, swapping the specified columns
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			// Determine which column to copy from
			sourceCol := j
			if j == col1 {
				sourceCol = col2
			} else if j == col2 {
				sourceCol = col1
			}

			// Copy the value
			result.values[i][j] = m.values[i][sourceCol]
		}
	}

	return result
}

// MultiplyRow multiplies a row in the matrix by a scalar value and returns a pointer to the resulting matrix.
// Returns nil if the row index is out of bounds.
func MultiplyRow(m Matrix, row int, scalar float64) *Matrix {
	// Check if row index is valid
	if row < 0 || row >= m.rows {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix, multiplying the specified row by the scalar
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			if i == row {
				// Multiply the row by the scalar
				result.values[i][j] = m.values[i][j] * scalar
			} else {
				// Copy the value as is
				result.values[i][j] = m.values[i][j]
			}
		}
	}

	return result
}

// MultiplyColumn multiplies a column in the matrix by a scalar value and returns a pointer to the resulting matrix.
// Returns nil if the column index is out of bounds.
func MultiplyColumn(m Matrix, col int, scalar float64) *Matrix {
	// Check if column index is valid
	if col < 0 || col >= m.columns {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix, multiplying the specified column by the scalar
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			if j == col {
				// Multiply the column by the scalar
				result.values[i][j] = m.values[i][j] * scalar
			} else {
				// Copy the value as is
				result.values[i][j] = m.values[i][j]
			}
		}
	}

	return result
}

// AddScaledRow adds a source row multiplied by a scalar to a target row in the matrix and returns a pointer to the resulting matrix.
// Returns nil if either row index is out of bounds.
func AddScaledRow(m Matrix, targetRow, sourceRow int, scalar float64) *Matrix {
	// Check if row indices are valid
	if targetRow < 0 || targetRow >= m.rows || sourceRow < 0 || sourceRow >= m.rows {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			if i == targetRow {
				// Add the source row multiplied by the scalar to the target row
				result.values[i][j] = m.values[i][j] + m.values[sourceRow][j]*scalar
			} else {
				// Copy the value as is
				result.values[i][j] = m.values[i][j]
			}
		}
	}

	return result
}

// AddScaledColumn adds a source column multiplied by a scalar to a target column in the matrix and returns a pointer to the resulting matrix.
// Returns nil if either column index is out of bounds.
func AddScaledColumn(m Matrix, targetCol, sourceCol int, scalar float64) *Matrix {
	// Check if column indices are valid
	if targetCol < 0 || targetCol >= m.columns || sourceCol < 0 || sourceCol >= m.columns {
		return nil
	}

	// Create a new matrix with the same dimensions
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			if j == targetCol {
				// Add the source column multiplied by the scalar to the target column
				result.values[i][j] = m.values[i][j] + m.values[i][sourceCol]*scalar
			} else {
				// Copy the value as is
				result.values[i][j] = m.values[i][j]
			}
		}
	}

	return result
}

// TransposeMatrix transposes a matrix (rows become columns, columns become rows) and returns a pointer to the resulting matrix.
func TransposeMatrix(m Matrix) *Matrix {
	// Create a new matrix with swapped dimensions
	result := &Matrix{
		rows:    m.columns,
		columns: m.rows,
		values:  make([][]float64, m.columns),
	}

	// Copy values from the original matrix, swapping row and column indices
	for i := 0; i < m.columns; i++ {
		result.values[i] = make([]float64, m.rows)
		for j := 0; j < m.rows; j++ {
			result.values[i][j] = m.values[j][i]
		}
	}

	return result
}

// RowEchelonForm performs Gaussian elimination on a matrix and returns a pointer to the resulting matrix in row echelon form.
func RowEchelonForm(m Matrix) *Matrix {
	// Create a copy of the input matrix to work with
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			result.values[i][j] = m.values[i][j]
		}
	}

	// Current row being processed
	row := 0

	// Process each column
	for col := 0; col < m.columns && row < m.rows; col++ {
		// Find the pivot row (first row with non-zero element in current column)
		pivotRow := -1
		for i := row; i < m.rows; i++ {
			if result.values[i][col] != 0 {
				pivotRow = i
				break
			}
		}

		// If no pivot found in this column, move to the next column
		if pivotRow == -1 {
			continue
		}

		// Swap the pivot row with the current row if they're different
		if pivotRow != row {
			// Create a temporary matrix for the swap
			temp := SwapRows(*result, row, pivotRow)
			if temp == nil {
				return nil // This should not happen, but handle it just in case
			}
			result = temp
		}

		// Scale the pivot row to make the pivot element 1
		pivotValue := result.values[row][col]
		if pivotValue != 1 {
			// Create a temporary matrix for the scaling
			temp := MultiplyRow(*result, row, 1/pivotValue)
			if temp == nil {
				return nil // This should not happen, but handle it just in case
			}
			result = temp
		}

		// Eliminate all other elements in the current column below the pivot
		for i := row + 1; i < m.rows; i++ {
			if result.values[i][col] != 0 {
				// Calculate the scalar to multiply the pivot row by
				scalar := -result.values[i][col]

				// Create a temporary matrix for the elimination
				temp := AddScaledRow(*result, i, row, scalar)
				if temp == nil {
					return nil // This should not happen, but handle it just in case
				}
				result = temp
			}
		}

		// Move to the next row
		row++
	}

	return result
}

// ReducedRowEchelonForm performs Gauss-Jordan elimination on a matrix and returns a pointer to the resulting matrix in reduced row echelon form.
func ReducedRowEchelonForm(m Matrix) *Matrix {
	// Create a copy of the input matrix to work with
	result := &Matrix{
		rows:    m.rows,
		columns: m.columns,
		values:  make([][]float64, m.rows),
	}

	// Copy values from the original matrix
	for i := 0; i < m.rows; i++ {
		result.values[i] = make([]float64, m.columns)
		for j := 0; j < m.columns; j++ {
			result.values[i][j] = m.values[i][j]
		}
	}

	// Current row being processed
	row := 0

	// Process each column
	for col := 0; col < m.columns && row < m.rows; col++ {
		// Find the pivot row (first row with non-zero element in current column)
		pivotRow := -1
		for i := row; i < m.rows; i++ {
			if result.values[i][col] != 0 {
				pivotRow = i
				break
			}
		}

		// If no pivot found in this column, move to the next column
		if pivotRow == -1 {
			continue
		}

		// Swap the pivot row with the current row if they're different
		if pivotRow != row {
			// Create a temporary matrix for the swap
			temp := SwapRows(*result, row, pivotRow)
			if temp == nil {
				return nil // This should not happen, but handle it just in case
			}
			result = temp
		}

		// Scale the pivot row to make the pivot element 1
		pivotValue := result.values[row][col]
		if pivotValue != 1 {
			// Create a temporary matrix for the scaling
			temp := MultiplyRow(*result, row, 1/pivotValue)
			if temp == nil {
				return nil // This should not happen, but handle it just in case
			}
			result = temp
		}

		// Eliminate all other elements in the current column (both above and below the pivot)
		for i := 0; i < m.rows; i++ {
			if i != row && result.values[i][col] != 0 {
				// Calculate the scalar to multiply the pivot row by
				scalar := -result.values[i][col]

				// Create a temporary matrix for the elimination
				temp := AddScaledRow(*result, i, row, scalar)
				if temp == nil {
					return nil // This should not happen, but handle it just in case
				}
				result = temp
			}
		}

		// Move to the next row
		row++
	}

	return result
}
