package task

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

// Save guarda las tareas en un archivo JSON.
// Recibe un slice de tareas y las guarda en un archivo con formato JSON.
// Devuelve un error si ocurre algún problema durante el proceso de guardado.
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

// Load lee los datos de las tareas desde un archivo JSON y devuelve un slice de tareas.
// Si el archivo no existe, se devuelve un slice vacío.
// Si hay un error al leer o deserializar los datos, se devuelve el error.
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

func PrintTasks() {
	tasks, err := Load()
	if err != nil {
		panic(err)
	}

	PrintTable(tasks)
}

// Filter filtra una lista de tareas utilizando una función de predicado.
// Recibe un slice de tareas y una función de predicado que toma una tarea y devuelve un valor booleano.
// Devuelve un nuevo slice de tareas que cumplen con el predicado.
func Filter(tasks []Task, p func(Task) bool) []Task {
	var result []Task
	for _, task := range tasks {
		if p(task) {
			result = append(result, task)
		}
	}

	return result
}
