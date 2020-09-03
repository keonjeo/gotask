package task

type Task struct {
	method func(interface{}) error // Task里面定义一个具体的业务方法
}

// 创建一个Task任务
func NewTask(method func(interface{}) error) *Task {
	return &Task{
		method: method,
	}
}

// Task中定义一个执行业务的方法
func (t*Task) Execute(opts interface{}) {
	t.method(opts) // 调用任务中已经绑定好的业务方法
}