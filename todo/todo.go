package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)


type item struct{
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}

type Todos []item 


func (t *Todos) Add(task string) {
	todo := item{
		Task: task,
		Done:false,
		CreatedAt:time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	ls := *t

	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index -1].Done = true
	ls[index -1].CompletedAt = time.Now()

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index < 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t= append(ls[:index-1], ls[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error{

	file , err := os.ReadFile(filename)

	if err !=nil{
		if errors.Is(err, os.ErrNotExist){
			return nil
		}
		return errors.New("could not read the file")
	}

	if len(file) == 0{
		return nil
	}

	err = json.Unmarshal(file, &t)

	if err != nil{
		return errors.New("could not unmarshal the file")
	}	


	return nil

}

func (t *Todos) Save(filename string) error{

	file, err:= json.Marshal(t)

	if err !=nil{
		return err
	}

	return os.WriteFile(filename, file,0644)
}

func (t *Todos) Print()  {
	if t == nil || len(*t) == 0{
		fmt.Println("You have no todos")
		return
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "No."},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "Created At"},
			{Align: simpletable.AlignRight, Text: "Completed At"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx,item := range *t{
		idx ++
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.Done)},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: item.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: fmt.Sprintf("You have  %d pending todos", t.CountPending())},
	}}



	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) CountPending() int{
	total:= 0

	for _,item := range *t{
		if !item.Done{
			total ++
		}
	}

	return total
}