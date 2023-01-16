package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	dataTaskstr := r.Context().Value("id").(string)
	if dataTaskstr == "" {
			w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid user id",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	dataTaskInt, _ := strconv.Atoi(dataTaskstr)
	
	dataGetTask := r.URL.Query().Get("task_id")	
	if dataGetTask == "" {
		dataResponen, err := t.taskService.GetTasks(r.Context(), dataTaskInt)
		if err != nil {
			w.WriteHeader(500)
			w.WriteHeader(http.StatusInternalServerError)
			Err := entity.ErrorResponse{
				Error: "error internal server",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
		}

		w.WriteHeader(200)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dataResponen)
	}else{
		dataResponen, err := t.taskService.GetTaskByID(r.Context(), dataTaskInt)
		if err != nil {
			w.WriteHeader(500)
			w.WriteHeader(http.StatusInternalServerError)
			Err := entity.ErrorResponse{
				Error: "error internal server",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
		}

		w.WriteHeader(200)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(dataResponen)
	}
	// TODO: answer here

}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	// datatask.Title = r.FormValue("title")
	// datatask.Description = r.FormValue("description")
	// datatask.CategoryID = r.FormValue("category_id")

	if task.Title == "" || task.Description == "" || task.CategoryID == 0 {
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := entity.ErrorResponse{
			Error: "invalid task request",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}


	dataTaskstr := r.Context().Value("id").(string)
	if dataTaskstr == "" {
			w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid user id",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	dataTaskint, _ := strconv.Atoi(dataTaskstr) 

	var hasilTask entity.Task

	hasilTask.Title = task.Title
	hasilTask.Description = task.Description
	hasilTask.CategoryID = task.CategoryID
	hasilTask.UserID = dataTaskint

	dataResponen, err := t.taskService.StoreTask(r.Context(), &hasilTask)
	
	if err != nil {
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Err := entity.ErrorResponse{
			Error: "error internal server",
		}
		dataErr, _ := json.Marshal(Err)	
		w.Write(dataErr)
		return	
	}

	w.WriteHeader(201)
	w.WriteHeader(http.StatusCreated)
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataResponen, "task_id": hasilTask.ID, "message": "success create new task"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	dataUntukId := r.Context().Value("id").(string)
	taskID := r.URL.Query().Get("task_id")

	taskIDInt, _ := strconv.Atoi(taskID)

	dataUntukIdInt, _ := strconv.Atoi(dataUntukId)

	dataDeleteTask:= t.taskService.DeleteTask(r.Context(), taskIDInt)

	if dataDeleteTask != nil {
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Err := entity.ErrorResponse{
			Error: "error internal server",
		}
		dataErr, _ := json.Marshal(Err)	
		w.Write(dataErr)
		return	
	}

	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataDeleteTask, "task_id": dataUntukIdInt, "message": "success delete task"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	dataTaskstr := r.Context().Value("id").(string)
	if dataTaskstr == "" {
			w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid user id",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	var hasilTask entity.Task

	hasilTask.Title = task.Title
	hasilTask.Description = task.Description
	hasilTask.CategoryID = task.CategoryID
	

	dataTaskint, _ := strconv.Atoi(dataTaskstr) 

	dataResponen, err := t.taskService.UpdateTask(r.Context(), &hasilTask)
	if err != nil {
		w.WriteHeader(500)
		w.WriteHeader(http.StatusInternalServerError)
		Err := entity.ErrorResponse{
			Error: "error internal server",
		}
		dataErr, _ := json.Marshal(Err)	
		w.Write(dataErr)
		return	
	}

	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataResponen, "task_id": dataTaskint, "message": "success update task"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     int(idLogin),
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
