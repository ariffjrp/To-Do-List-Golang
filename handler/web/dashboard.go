package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
	"log"
	"net/http"
	"text/template"
	"path"
)

type DashboardWeb interface {
	Dashboard(w http.ResponseWriter, r *http.Request)
}

type dashboardWeb struct {
	categoryClient client.CategoryClient
	embed          embed.FS
}

func NewDashboardWeb(catClient client.CategoryClient, embed embed.FS) *dashboardWeb {
	return &dashboardWeb{catClient, embed}
}

func (d *dashboardWeb) Dashboard(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	categories, err := d.categoryClient.GetCategories(userId.(string))
	if err != nil {
		log.Println("error get cat: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dataTemplate = map[string]interface{}{
		"categories": categories,
	}

	var funcMap = template.FuncMap{
		"categoryInc": func(catId int) int {
			return catId + 1
		},
		"categoryDec": func(catId int) int {
			return catId - 1
		},
	}

	// ignore this
	_ = dataTemplate
	_ = funcMap
	//

	var filelogin = path.Join("views", "main", "dashboard.html")
	var fileheader = path.Join("views", "general", "header.html")
	webLogin := template.Must(template.New("dashboard.html").Funcs(funcMap).ParseFS(d.embed, filelogin, fileheader))

	// var data = map[string]interface{}{
	// 	"title": ">Kanban App",
	// }	

	
	Alert := webLogin.ExecuteTemplate(w, "dashboard.html",dataTemplate)

	if Alert != nil {
		http.Error(w, Alert.Error(), http.StatusInternalServerError)
		return
	}
	// TODO: answer here
}
