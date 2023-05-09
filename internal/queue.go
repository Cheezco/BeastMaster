package internal

import "fmt"

type Queue struct {
	Items []interface{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) Push(value interface{}) {
	q.Items = append(q.Items, value)
	fmt.Print("Added to queue:")
	fmt.Println(value)
}

func (q *Queue) Pop() interface{} {
	if q.IsEmpty() {
		return nil
	}

	val := q.Items[0]
	q.Items = q.Items[1:]
	return val
}
