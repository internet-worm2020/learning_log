package jobs

import "fmt"

type TestJob struct {
	Id   int
	Name string
}

func (this TestJob) Run() {
	fmt.Println(this.Id, this.Name)
	fmt.Println("testJob1...")
}
