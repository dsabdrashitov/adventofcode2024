package comparableencoder

type ComparableEncoder[T comparable] struct {
	code   map[T]int
	decode []T
}

func New[T comparable]() *ComparableEncoder[T] {
	return &ComparableEncoder[T]{make(map[T]int), make([]T, 0)}
}

func (encoder *ComparableEncoder[T]) Id(item T) int {
	if id, ok := encoder.code[item]; ok {
		return id
	}
	result := len(encoder.decode)
	encoder.decode = append(encoder.decode, item)
	encoder.code[item] = result
	return result
}

func (encoder *ComparableEncoder[T]) Item(id int) T {
	return encoder.decode[id]
}