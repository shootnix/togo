package main

import (
	"errors"
	"g0/models"
	"g0/views"
	"os"
	"strconv"
)

func main() {
	models.InitStorage()

	var command string
	if len(os.Args) <= 1 {
		command = "-l"
	} else {
		command = os.Args[1]
	}

	cmd, err := whatCanIDo(command)
	if err != nil {
		// TODO: no panic at the disco!
		panic(err)
	}

	screen := views.NewScreen()

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			println("Not enough arguments!")
			break
		}
		taskStr := os.Args[2]
		t := models.NewTask()
		t.SetTask(taskStr)
		_, err := t.Save()
		if err != nil {
			panic(err)
		}

	case "list":
		break

	case "resolve":
		if len(os.Args) < 3 {
			println("Not enough arguments!")
			break
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		//screen.Render(t)
		_, err = t.Resolve()
		if err != nil {
			println(err.Error())
			break
		}

	case "delete":
		if len(os.Args) < 3 {
			println("Not enough argumens!")
			break
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		_, err = t.Delete()
		if err != nil {
			println(err.Error())
			break
		}

	case "mark":
		if len(os.Args) < 4 {
			println("Not enough arguments!")
		}
		id, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		_, err = t.Mark(os.Args[2])
		if err != nil {
			println(err.Error())
			break
		}

	case "unmark":
		if len(os.Args) < 4 {
			println("Not enough arguments!")
			break
		}
		id, err := strconv.ParseInt(os.Args[3], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		_, err = t.UnMark(os.Args[2])
		if err != nil {
			println(err.Error())
			break
		}

	case "up":
		if len(os.Args) < 3 {
			println("Not enough arguments!")
			break
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		_, err = t.Up()
		if err != nil {
			println(err.Error())
			break
		}

	case "down":
		if len(os.Args) < 3 {
			println("Not enough arguments!")
			break
		}
		id, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			println("Error:", err.Error())
			break
		}
		t, err := models.GetTask(id)
		if err != nil {
			println(err.Error())
			break
		}
		_, err = t.Down()
		if err != nil {
			println(err.Error())
			break
		}
	}

	screen.RenderList(models.ListTasks())
}

func whatCanIDo(cmd string) (string, error) {
	m := map[string]string{
		"-l":    "list",
		"-r":    "resolve",
		"-add":  "add",
		"-m":    "mark",
		"-d":    "delete",
		"-um":   "unmark",
		"-up":   "up",
		"-down": "down",
		"-h":    "help",
	}
	command := m[cmd]
	if command == "" {
		return "", errors.New("Unknown command: " + cmd)
	}

	return m[cmd], nil
}
