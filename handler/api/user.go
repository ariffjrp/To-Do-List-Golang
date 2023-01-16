package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserAPI interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)

	Delete(w http.ResponseWriter, r *http.Request)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Login(w http.ResponseWriter, r *http.Request) {
	var user entity.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	// user.Email = r.FormValue("email")
	// user.Password = r.FormValue("password")

	if  user.Email == "" || user.Password == ""{
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := entity.ErrorResponse{
			Error: "email or password is empty",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}

	var dataUser entity.User

	dataUser.Email = user.Email
	dataUser.Password = user.Password

	dataValue, dataKeperluan := u.userService.Login(r.Context(), &dataUser)
	if dataKeperluan != nil {
			w.WriteHeader(500)
			w.WriteHeader(http.StatusInternalServerError)
			Err := entity.ErrorResponse{
				Error: "error internal server",
			}
			dataErr, _ := json.Marshal(Err)	
			w.Write(dataErr)
			return	
	}

	hasil := strconv.Itoa(dataValue)

	http.SetCookie(w, &http.Cookie{
		Name	:   "user_id",
		Value 	: 	 hasil,
	})

	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": dataValue, "message": "login success"}) 
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (u *userAPI) Register(w http.ResponseWriter, r *http.Request) {
	var user entity.UserRegister

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	// user.Fullname = r.FormValue("fullname")
	// user.Email = r.FormValue("email")
	// user.Password = r.FormValue("password")
	

	if user.Fullname == "" || user.Email == "" || user.Password == ""{
		w.WriteHeader(400)
		w.WriteHeader(http.StatusBadRequest)
		Alert := entity.ErrorResponse{
			Error: "register data is empty",
		}
		dataAlert, _ := json.Marshal(Alert)
		w.Write(dataAlert)
		return
	}
	
	var dataUser entity.User

	dataUser.Fullname = user.Fullname
	dataUser.Email = user.Email
	dataUser.Password = user.Password

	data, dataEmail := u.userService.Register(r.Context(), &dataUser)
	if dataEmail != nil {
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
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"user_id": data.ID, "message": "register success"})
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (u *userAPI) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name	:   "user_id",
		Value 	: 	 "",
	})
	w.WriteHeader(200)
	w.WriteHeader(http.StatusOK)
	outputKetikaSukses, _ := json.Marshal(map[string]interface{}{"message": "logout success"})
	w.Write(outputKetikaSukses)
	// TODO: answer here
}

func (u *userAPI) Delete(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("user_id is empty"))
		return
	}

	deleteUserId, _ := strconv.Atoi(userId)

	err := u.userService.Delete(r.Context(), int(deleteUserId))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"message": "delete success"})
}
