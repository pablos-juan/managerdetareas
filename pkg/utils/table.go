package utils

import (
	"fmt"

	"github.com/alexeyco/simpletable"
	"github.com/pablos-juan/managerdetareas/internal/task"
)

func PrintTasks(tasks []task.Task) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "Tarea"},
			{Align: simpletable.AlignCenter, Text: "Completada"},
		},
	}

	for i, v := range tasks {
		tn := v.Name
		ts := fmt.Sprintf("%t", v.Done)
		if v.Done {
			tn = Green(tn)
			ts = Green(ts)
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprint(i + 1)},
			{Text: tn},
			{Align: simpletable.AlignCenter, Text: ts},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	n := IncompleteTasks(tasks)
	tx := Green("Nada por hacer :)")
	if n > 0 {
		tx = fmt.Sprintf("Tareas por terminar: %d", n)
		tx = Red(tx)
	}
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 3, Text: tx},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func IncompleteTasks(tasks []task.Task) int {
	n := 0
	for _, v := range tasks {
		if !v.Done {
			n++
		}
	}
	return n
}
