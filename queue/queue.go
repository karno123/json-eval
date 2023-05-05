package queue

import "errors"

type Queue struct {
	Items []string
}

func (q *Queue) EnQueue(item string) {
	q.Items = append(q.Items, item)
}

func (q *Queue) DeQue() (string, error) {
	if len(q.Items) == 0 {
		return "", errors.New("")
	}

	item := q.Items[len(q.Items)-1]
	q.Items = q.Items[0 : len(q.Items)-2]
	return item, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) <= 0
}
