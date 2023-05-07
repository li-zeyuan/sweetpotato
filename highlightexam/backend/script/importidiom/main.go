package main

import (
	"fmt"

	"github.com/li-zeyuan/common/utils"
	"github.com/li-zeyuan/sun/highlightexam/onrequest"
	"github.com/li-zeyuan/sweetpotato/highlightexam/backend/testdata"
)

var cfg = Config{
	Tasks: []*Task{
		{
			SubjectName:     "言语理解易错成语",
			SubjectDescribe: "言语理解易错成语积累",
			File:            "1言语理解易错成语.txt",
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

		idioms = utils.UniqueStr(idioms)
		err = handle(task, idioms)
		if err != nil {
			return
		}
	}

	fmt.Println("succesfully....")
}

func handle(task *Task, idioms []string) error {
	subject, err := upsertSubject(task)
	if err != nil {
		return err
	}

	batcher, err := utils.NewBatcher(len(idioms), utils.DefaultBatchSize)
	if err != nil {
		fmt.Println("new batch error", err)
		return err
	}

	start, length := 0, 0
	for batcher.Iter(&start, &length) {
		batchIdioms := idioms[start : start+length]
		batchIdioms, err = uniqueByDB(task, batchIdioms)
		if err != nil {
			return err
		}

		IdiomList, err := onrequest.Handle(batchIdioms)
		if err != nil {
			return err
		}

		err = createKnowledge(subject.ID, IdiomList)
		if err != nil {
			return err
		}
	}

	err = updateSubjectTotal(subject.ID)
	if err != nil {
		return err
	}

	return nil
}

func uniqueByDB(task *Task, idioms []string) ([]string, error) {
	subject, err := getSubject(task)
	if err != nil {
		return nil, err
	}
	if subject.ID == 0 {
		return idioms, nil
	}

	batchIdiom, err := getKnowledge(idioms, subject.ID)
	if err != nil {
		return nil, err
	}

	return batchIdiom, nil
}
