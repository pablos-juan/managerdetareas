package utilities

import "strconv"

// GetIndex devuelve el índice entero representado por el primer argumento en la lista de argumentos.
// Si el argumento no es un número entero válido, se mostrará un mensaje de error y se devolverá -1.
// Si no se proporciona ningún argumento o se proporciona más de un argumento, se mostrará un mensaje de error y se devolverá -1.
//
// Parámetros:
// - args: una lista de argumentos de tipo string.
//
// Retorna:
// - int: el índice entero representado por el primer argumento en la lista de argumentos, o -1 si hay un error.
//
// Ejemplo:
//
//	args := []string{"5"}
//	index := GetIndex(args)
//	// index = 5
func GetIndex(args []string) int {
	if len(args) == 1 {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			ErrorMessage("Error. '%s' no es un número!", args[0])
			return -1
		}

		return i
	}

	ErrorMessage("Error. Necesita un número entero como argumento!")
	return -1
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
