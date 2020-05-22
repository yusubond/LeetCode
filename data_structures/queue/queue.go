package queue

//---------------------------------------------------------------
// Queue
// FIFO先入先出的数据结构，属于线性表的一种
//---------------------------------------------------------------

type Queue struct {
	queue []interface{}
	front, rear int
	size int
}

func NewQueue(s int) *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, s, s)
	q.front = -1
	q.rear  = -1
	q.size = s
	return q
}

func (q *Queue) Push(e interface{}) {
	q.rear = (q.rear+1) % q.size
	// 为空，更新队头
	if q.front == -1 {
		q.front = 0
	}
	q.queue[q.rear] = e
}

func (q *Queue) Pop() {
	// 为空，或者只有一个元素
	if q.front == -1 || q.front == q.rear {
		q.front, q.rear = -1, -1
		return
	}
	q.front = (q.front+1) % q.size
}

func (q *Queue) Front() interface{} {
	if q.front == -1 { return nil }
	return q.queue[q.front]
}

func (q *Queue) Back() interface{} {
	if q.front == -1 { return nil }
	return q.queue[q.rear]
}

func (q *Queue) IsEmpty() bool {
	return q.front == -1
}

func (q *Queue) IsFull() bool {
	return ((q.rear+1) % q.size) == q.front
}

// 返回对列中的元素个数
func (q *Queue) Size() int {
	if q.front == -1 { return 0 }
	if q.rear == q.front { return 1 }
	if q.rear > q.front { return q.rear - q.front + 1 }
	return q.rear + q.size - q.front
}

