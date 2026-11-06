package http

import "encoding/json"

type HTTPHandlers struct {

}

func NewHTTPHandlers(todoList*todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}

}

func (h*HTTPHandlers)HandlerCreateTasks(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO:=ErrorDTO{
			Massage: err.Error(),
			Time:time.Now(),
		}
	}
}

func (h*HTTPHandlers)HandlerGetTasks(w http.ResponseWriter, r *http.Request) {

}

func (h*HTTPHandlers)HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h*HTTPHandlers)HandlerGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}
func (h*HTTPHandlers)HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
}

func (h*HTTPHandlers)HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {