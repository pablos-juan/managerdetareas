package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const FILE_NAME = "mann.json"

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func Save(nt Task) error {
	tasks, err := load()
	if err != nil {
		return err
	}

	tasks = append(tasks, nt)
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	err = os.WriteFile(FILE_NAME, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func load() ([]Task, error) {
	tasks := []Task{}

	data, err := os.ReadFile("mann.json")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return tasks, nil
		}
		return tasks, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}
