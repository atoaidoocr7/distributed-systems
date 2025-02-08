package app

import "fmt"

type Master struct {
}

func (m *Master) Start() {

	fmt.Println("We are starting the masterv2")
}
