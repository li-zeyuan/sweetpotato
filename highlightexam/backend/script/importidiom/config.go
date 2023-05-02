package main

type Config struct {
	Tasks []*Task
}

type Task struct {
	SubjectName     string
	SubjectDescribe string
	File            string
}
