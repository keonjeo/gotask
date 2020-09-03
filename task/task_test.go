package task

import (
	"fmt"
	"testing"
)

func TestExecute(t * testing.T) {
	sayHi := func (name interface{}) error {
		fmt.Printf("hi, %v\n", name)
		return nil
	}

	t1 := NewTask(sayHi)
	t1.Execute("Tony")
}
