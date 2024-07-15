package utilities

import (
	"encoding/json"
	"errors"
	"os"
)

const FILE_NAME = "mann.json"

type Task struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

func Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(FILE_NAME, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Load() ([]Task, error) {
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
