package file

// 文件操作者描述
//
// 每一个文件操作者视为一个线程。
//
// todo即表示了这个文件操作者所需要处理的文件处理待办事项。
type filer struct {
	todo 			chan Handler
}

// 添加一个待办事项。
func (slf *filer) add(handler Handler) {
	slf.todo <- handler
}

// 开始运行文件操作者。
func (slf *filer) run(idle chan <- *filer) *filer {
	go func() {
		for handler := range slf.todo {
			handler.Exec()
			idle <- slf
		}
	}()
	return slf
}

// 构建一个文件操作者实例并返回。
func newFiler() *filer {
	slf := &filer{
		todo: make(chan Handler),
	}
	return slf
}