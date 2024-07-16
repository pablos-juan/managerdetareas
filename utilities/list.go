package utilities

func PrintTasks() {
	tasks, err := Load()
	if err != nil {
		panic(err)
	}

	PrintTable(tasks)
}
