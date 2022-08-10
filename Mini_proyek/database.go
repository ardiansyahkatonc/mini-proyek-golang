package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Task_pegawai struct {
	Id_task      int    `json: "id_task"`
	Nama_pegawai string `json: "nama_pegawai"`
	Pesan        string `json: "pesan"`
	Deadline     string `json: "deadline"`
}

func main() {
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
		fmt.Println(task_pegawai.Id_task, task_pegawai.Nama_pegawai, task_pegawai.Pesan, task_pegawai.Deadline)
	}
}
