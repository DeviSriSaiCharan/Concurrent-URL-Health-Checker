package queue

type Queue[T any] struct {
	elements []T
}

func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{
		elements: make([]T, 0),
	}
}

func (q *Queue[T]) Enque(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Dequeue() {
	q.elements = q.elements[1:]
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
