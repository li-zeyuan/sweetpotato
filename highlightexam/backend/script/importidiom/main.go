package main

import (
	"fmt"

	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
	"gorm.io/gorm"
)

var cfg = Config{
	Tasks: []*Task{
		{
			SubjectName:     "subject_test",
			SubjectDescribe: "subject_describe_test",
			File:            "1.txt",
		},
	},
}

func main() {
	_, err := testdata.InitServer()
	if err != nil {
		fmt.Println("init server error", err)
		return
	}

	for _, task := range cfg.Tasks {
		idioms, err := GetIdiomFromFile(task.File)
		if err != nil {
			return
		}

		err = handle(task, idioms)
		if err != nil {
			return
		}
	}

}

func handle(task *Task, idioms []string) error {
	batcher, err := utils.NewBatcher(len(idioms), utils.DefaultBatchSize)
	if err != nil {
		fmt.Println("new batch error", err)
		return err
	}

	start, length := 0, 0
	for batcher.Iter(&start, &length) {
		batchIdioms := utils.UniqueStr(idioms[start : start+length])
		batchIdioms, err = uniqueByDB(task, batchIdioms)
		if err != nil {
			return err
		}

		// todo 爬虫 https://github.com/gocolly/colly
		// todo insert knowledge
	}

	// todo upsert subject

	return nil
}

func uniqueByDB(task *Task, idioms []string) ([]string, error) {
	subject, err := getSubject(task)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idioms, nil
		}

		return nil, err
	}

	batchIdiom, err := getKnowledge(idioms, subject.ID)
	if err != nil {
		return nil, err
	}

	return batchIdiom, nil
}
