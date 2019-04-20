// Package algebra provides sparse vector implementation.
package algebra

// SparseVector is used in matrix-vector multiplication when huge number of matrix entries are zero.
// For example, when matrix size is 10 billion, but the average number of nonzero elements
// per row is less than 10. For such an application, using symbol tables speeds up
// matrix-vector multiplication by a factor of a billion.
type SparseVector struct {
	st map[int]float64
}

// NewSparseVector returns a new sparse vector created from the row.
func NewSparseVector(row []float64) *SparseVector {
	vec := SparseVector{st: make(map[int]float64)}
	for i := range row {
		if row[i] == 0 {
			continue
		}
		vec.st[i] = row[i]
	}
	return &vec
}

// Dot computes a dot product of the sparse vector and the column.
func (vec *SparseVector) Dot(column []float64) (sum float64) {
	for k, v := range vec.st {
		sum += column[k] * v
	}
	return sum
}
