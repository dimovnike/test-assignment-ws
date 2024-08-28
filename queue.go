package main

/* implements a simle queue for better code clarity in the algorithm itself */

type Queue[T any] struct {
	q []T
}

func (q *Queue[T]) Queue(v ...T) {
	q.q = append(q.q, v...)
}

func (q *Queue[T]) Deqeue() (v T, found bool) {
	if len(q.q) == 0 {
		return v, false
	}

	v, q.q = q.q[0], q.q[1:]
	return v, true
}
