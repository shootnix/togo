package models

import (
	"database/sql"
	"errors"
)

// Task object
type Task struct {
	id       int64
	task     sql.NullString
	created  string
	resolved sql.NullInt64
	priority int64
}

// NewTask : task object constructor
func NewTask() (t *Task) {
	t = &Task{}

	return
}

// Task : get task
func (t *Task) Task() string {

	return t.task.String
}

// SetTask : set task
func (t *Task) SetTask(tstring string) {
	valid := tstring != ""
	t.task = sql.NullString{
		String: tstring,
		Valid:  valid,
	}
}

// ID : get task id
func (t *Task) ID() int64 {
	return t.id
}

// SetID : set task id
func (t *Task) SetID(uid int64) {
	t.id = uid
}

// Priority : get priority value
func (t *Task) Priority() int64 {
	return t.priority
}

// SetPriority : set priority value
func (t *Task) SetPriority(priority int64) {
	t.priority = priority
}

// GetTask : load task object from database
func GetTask(taskID int64) (*Task, error) {
	//println("sorage workdir =", storage.WorkDir)
	t := &Task{}
	selectSQLQuery := `
		SELECT id, task, priority
		  FROM tasks
		 WHERE id = ?;
	`
	var id int64
	var priority int64
	var task sql.NullString
	err := storage.DbHandler.QueryRow(selectSQLQuery, taskID).Scan(&id, &task, &priority)
	if err != nil {
		return t, errors.New("Task not found")
	}
	t.SetID(id)
	t.SetTask(task.String)
	t.SetPriority(priority)

	return t, nil
}

// Resolve : mark task as a resolved
func (t *Task) Resolve() (bool, error) {
	updateSQLQuery := `
		UPDATE tasks
		   SET resolved = CURRENT_DATE
		 WHERE id = ?;
	`
	_, err := storage.DbHandler.Exec(updateSQLQuery, t.ID())
	if err != nil {
		return falseAndError("Cannot update task")
	}

	return true, nil
}

// Delete : delete task
func (t *Task) Delete() (bool, error) {
	deleteSQLQuery := `
		DELETE FROM tasks
		WHERE id = ?;
	`
	_, err := storage.DbHandler.Exec(deleteSQLQuery, t.ID())
	if err != nil {
		return falseAndError("Cannot delete task")
	}

	return true, nil
}

// Mark : mark task (add some string in the beginning of task)
func (t *Task) Mark(mark string) (bool, error) {
	t.SetTask(mark + " " + t.Task())
	_, err := t.Save()
	if err != nil {
		return falseAndError("Can't save task")
	}

	return true, nil
}

// UnMark : delete some string in the beginning of task
func (t *Task) UnMark(mark string) (bool, error) {
	// TODO: ...
	return true, nil
}

// Up : raise priority of the task (increase value by 1)
func (t *Task) Up() (bool, error) {
	t.SetPriority(t.Priority() + 1)
	_, err := t.Save()
	if err != nil {
		return falseAndError("Can't save task")
	}

	return true, nil
}

// Down : reduce priority of the task (by 1)
func (t *Task) Down() (bool, error) {
	t.SetPriority(t.Priority() - 1)
	_, err := t.Save()
	if err != nil {
		return falseAndError("Can't save task")
	}

	return true, nil
}

// Save : save task
func (t *Task) Save() (bool, error) {
	insertSQLQuery := `
		INSERT INTO tasks (task)
		VALUES (?);
	`
	checkSQLQuery := `
		SELECT id
		  FROM tasks
		 WHERE task = ?;
	`
	updateSQLQuery := `
		UPDATE tasks 
		SET 
			task = ?,
			priority = ?
		WHERE
			id = ?;
	`

	if t.ID() == 0 {
		// new task
		// check
		var id int64
		err := storage.DbHandler.QueryRow(checkSQLQuery, t.Task()).Scan(&id)
		//fmt.Println("id =", id)
		if err == nil && id != 0 {
			return falseAndError("Task is already exists")
		}

		// insert new:
		stm, err := storage.DbHandler.Prepare(insertSQLQuery)
		if err != nil {
			return false, errors.New("Cannot prepare statement: " + err.Error())
		}
		defer stm.Close()
		res, err := stm.Exec(t.Task())
		if err != nil {
			return false, errors.New("Cannot execute statement: " + err.Error())
		}
		id, err = res.LastInsertId()
		if err != nil {
			return false, errors.New("Cannot get last insert id: " + err.Error())
		}
		t.SetID(id)
	} else {
		// task from db
		_, err := storage.DbHandler.Exec(updateSQLQuery, t.Task(), t.Priority(), t.ID())
		if err != nil {
			return falseAndError("Can not update task:" + err.Error())
		}
	}

	return true, nil
}

// ListTasks : get list of resolved task ordered by priority
func ListTasks() []*Task {
	var list []*Task
	sqlQuery := `
		SELECT id, task
		  FROM tasks
   		 WHERE resolved IS NULL
   		ORDER BY priority DESC; 
	`
	rows, err := storage.DbHandler.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var id int64
	var task sql.NullString
	for rows.Next() {
		err = rows.Scan(&id, &task)
		if err != nil {
			panic(err)
		}

		t := NewTask()
		t.SetID(id)
		t.SetTask(task.String)

		list = append(list, t)
	}

	return list
}
