package utilities

import "strconv"

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
