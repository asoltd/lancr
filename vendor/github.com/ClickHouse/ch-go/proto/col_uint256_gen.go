// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColUInt256 represents UInt256 column.
type ColUInt256 []UInt256

// Compile-time assertions for ColUInt256.
var (
	_ ColInput  = ColUInt256{}
	_ ColResult = (*ColUInt256)(nil)
	_ Column    = (*ColUInt256)(nil)
)

// Rows returns count of rows in column.
func (c ColUInt256) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColUInt256) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of UInt256.
func (ColUInt256) Type() ColumnType {
	return ColumnTypeUInt256
}

// Row returns i-th row of column.
func (c ColUInt256) Row(i int) UInt256 {
	return c[i]
}

// Append UInt256 to column.
func (c *ColUInt256) Append(v UInt256) {
	*c = append(*c, v)
}

// LowCardinality returns LowCardinality for UInt256 .
func (c *ColUInt256) LowCardinality() *ColLowCardinality[UInt256] {
	return &ColLowCardinality[UInt256]{
		index: c,
	}
}

// Array is helper that creates Array of UInt256.
func (c *ColUInt256) Array() *ColArr[UInt256] {
	return &ColArr[UInt256]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(UInt256).
func (c *ColUInt256) Nullable() *ColNullable[UInt256] {
	return &ColNullable[UInt256]{
		Values: c,
	}
}

// NewArrUInt256 returns new Array(UInt256).
func NewArrUInt256() *ColArr[UInt256] {
	return &ColArr[UInt256]{
		Data: new(ColUInt256),
	}
}
