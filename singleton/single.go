package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("new instance created")
			singleInstance = &single{}
		} else {
			fmt.Println("instance created already")
		}
	} else {
		fmt.Println("instance created already")
	}
	return singleInstance
}
