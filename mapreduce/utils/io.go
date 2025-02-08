package utils

import application "github.com/atoaidoocr7/distributed-systems/mapreduce/app"

// Everything io related

type IOUtils struct {
}

func NewIOUtils() {
	m := application.Master{}

	m.Start()
}
