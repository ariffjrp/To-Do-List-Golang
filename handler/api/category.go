package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("id").(string)
	if data == "" {
			w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid user id",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	dataID, _ := strconv.Atoi(data)
	
	dataRespoen, err := c.categoryService.GetCategories(r.Context(), dataID)
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
	json.NewEncoder(w).Encode(dataRespoen)
	// TODO: answer here
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	if category.Type == "" {	
		w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid category request",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return
	}

	dataUntukId := r.Context().Value("id").(string)
	
	if dataUntukId == "" {
			w.WriteHeader(400)
			w.WriteHeader(http.StatusBadRequest)
			Err := entity.ErrorResponse{
				Error: "invalid user id",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	dataID, _ := strconv.Atoi(dataUntukId)

	var dataNewCatagories entity.Category

	dataNewCatagories.Type = category.Type
	dataNewCatagories.UserID = dataID


	dataCategoriesValue, Alert := c.categoryService.StoreCategory(r.Context(), &dataNewCatagories)
	if Alert != nil {
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
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataCategoriesValue, "category_id": dataNewCatagories.ID, "message": "success create new category"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	dataUntukId := r.Context().Value("id").(string)
	categoryIDstr := r.URL.Query().Get("category_id")

	categoryIDint, _ := strconv.Atoi(categoryIDstr)

	dataUntukIdInt, _ := strconv.Atoi(dataUntukId)

	dataUntukCategories := c.categoryService.DeleteCategory(r.Context(),categoryIDint)
	if dataUntukCategories != nil {
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
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataUntukCategories, "category_id": dataUntukIdInt, "message": "success delete category"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
