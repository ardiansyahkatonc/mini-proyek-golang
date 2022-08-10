package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

type Task_pegawai struct {
	Id_task      int    `json: "id_task"`
	Nama_pegawai string `json: "nama_pegawai"`
	Pesan        string `json: "pesan"`
	Deadline     string `json: "deadline"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error Bro", http.StatusInternalServerError)
		return
	}

	db, err := sql.Open("mysql", "root:Katon2022@tcp(127.0.0.1:3306)/mini_proyek")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT id_task, nama_pegawai, pesan, deadline FROM task_pegawai")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var task_pegawai Task_pegawai

		err = results.Scan(&task_pegawai.Id_task, &task_pegawai.Nama_pegawai, &task_pegawai.Pesan, &task_pegawai.Deadline)
		if err != nil {
			panic(err.Error())
		}

	}

	err = tmpl.Execute(w, r)

}
func BrowseHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Browse"))
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Pegawai"))

}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Pesan : %d", idNumb)

}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello About"))

}
