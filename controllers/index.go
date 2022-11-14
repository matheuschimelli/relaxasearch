package controller

import (
	"html/template"

	"net/http"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type SearchData struct {
	Title   string
	Content string
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

// func GetSearch(w http.ResponseWriter, r *http.Request) {
// 	tmpl := template.Must(template.ParseFiles("index.html"))

// 	if r.Method != http.MethodGet {
// 		tmpl.Execute(w, nil)
// 		return
// 	}

// 	queryString := r.FormValue("q")

// 	data := TodoPageData{
// 		PageTitle: "My TODO list",
// 		Todos: []Todo{
// 			{Title: utils.Normalize(queryString), Done: false},
// 			{Title: "Task 2", Done: true},
// 			{Title: "Task 3", Done: true},
// 		},
// 	}
// 	tmpl.Execute(w, data)
// }

// func HandleIndex(w http.ResponseWriter, r *http.Request) {

// 	if r.Method == http.MethodGet {
// 		indexName := strings.TrimPrefix(r.URL.Path, urls.UrlsV1.Index)
// 		index, err := bleve.GetIndex(indexName)

// 		if err != nil {
// 			fmt.Println(err)
// 			j, _ := json.Marshal(response{
// 				Success: false,
// 				Message: "Cannot find a index with provided name",
// 				Data:    "",
// 			})
// 			w.Header().Set("Content-Type", "application/json")
// 			w.Write(j)
// 			return
// 		}

// 		j, _ := json.Marshal(index)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(j)
// 	}

// 	if r.Method == http.MethodPost {
// 		indexName := strings.TrimPrefix(r.URL.Path, urls.UrlsV1.Index)
// 		docId := r.URL.Query().Get("id")
// 		title := r.FormValue("title")
// 		content := r.FormValue("content")

// 		index, err := bleve.IndexDocument(indexName, docId, SearchData{Title: title, Content: content})

// 		fmt.Println(index)

// 		fmt.Println(title)
// 		fmt.Println(content)

// 		if err != nil {
// 			fmt.Println(err)
// 			j, _ := json.Marshal(response{
// 				Success: false,
// 				Message: fmt.Sprintf("Error: %s", err),
// 				Data:    "",
// 			})
// 			w.Header().Set("Content-Type", "application/json")
// 			w.Write(j)
// 			return
// 		}

// 		j, _ := json.Marshal(index)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(j)

// 	}

// }
