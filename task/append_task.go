package task

//AppendTasks will append the tasks
type AppendTasks struct{}

//GetName will get the name
func (p *AppendTasks) GetName() string {
	return "Grape_Pie"
}

//GetDescr will get the description
func (p *AppendTasks) GetDescr() string {
	return "appends the tasks"
}

//Run will execute the AppendTasks
func (p *AppendTasks) Run() error {

	// app := task.Application{}
	// app.AddTask(&task.DoubleListCounter{})

	// app.AddTask(&task.MultiTasking{})

	// app.AddTask(&task.NumberConvert{})

	// app.AddTask(&task.WordConvert{})

	// app.AddTask(&task.ConvertInput{})

	// app.AddTask(&task.RespondInput{})

	// app.AddTask(&task.Summing{})

	// app.AddTask(&task.FirstCode{})

	// app.AddTask((&task.ReadWrite{}))

	// taskName := ""
	// if len(os.Args) > 1 {
	// 	taskName = os.Args[1]
	// }
	// err := app.RunTask(taskName)
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	return
	// }
	// return
	return ErrUnknownTask
}
