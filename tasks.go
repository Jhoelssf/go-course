package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) addToList(tl *task) {
	t.tasks = append(t.tasks, tl)
}

type task struct {
	name string
	description string
	completed bool
}

func main() {
	t := &task{
		name: "Completar el curso de Go",
		description: "Aprender Go",
	}
	t1 := &task{
		name: "Completar el curso de Python",
		description: "Aprender Python",
	}
	//t2 := &task{
	//	name: "Completar el curso de NodeJS",
	//	description: "Aprender NodeJS",
	//}
	//fmt.Printf("%+v\n", t)
	//t.markAsComplete()
	//t.changeDescription("finaizar Curso GO")
	//t.changeName("Go")
	//fmt.Printf("%+v\n", t)
	list := &taskList{
		tasks: []*task{
			t, t1,
		},
	}

	//fmt.Println(list.tasks[0])
	//list.addToList(t2)
	////fmt.Println(len(list.tasks))
	//list.deleteofList(1)
	//list.tasks[0].completed = true
	//list.printList()
	//fmt.Println("Completed")
	//list.printListCompleted()

	mapTasks := make(map[string]*taskList)

	mapTasks["Jhoel"] = list

	t3 := &task{
		name: "Completar el curso de Java",
		description: "Aprender Java",
	}
	t4 := &task{
		name: "Completar el curso de C#",
		description: "Aprender C#",
	}
	list2 := &taskList{
		tasks: []*task{
			t3, t4,
		},
	}

	mapTasks["Marco"] = list2

	fmt.Println("Tareas de Jhoel")
	mapTasks["Jhoel"].printList()

	fmt.Println("Tareas de Marco")
	mapTasks["Marco"].printList()
}

func (t *taskList) printListCompleted()  {
	for _, value := range t.tasks {
		if value.completed == true {
			fmt.Println("name:", value.name)
			fmt.Println("description:", value.description)
		}
	}
}

func (t *taskList) printList()  {
	for _, value := range t.tasks {
		fmt.Println("name:", value.name)
		fmt.Println("description:", value.description)
	}
}

func (t *taskList) deleteofList(index int)  {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *task) markAsComplete()  {
	t.completed = true
}

func (t *task) changeDescription(description string)  {
	t.description = description
}

func (t *task) changeName(name string)  {
	t.name = name
}