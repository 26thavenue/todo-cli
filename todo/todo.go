package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"
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

	if index < 0 || index >= len(ls) {
		return errors.New("invalid index")
	}

	ls[index -1].Done = true
	ls[index -1].CompletedAt = time.Now()

	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t

	if index < 0 || index >= len(ls) {
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
