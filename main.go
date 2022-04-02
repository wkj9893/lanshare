package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := strings.TrimPrefix(r.URL.Path, "/api/upload/")
	f, _, err := r.FormFile(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer f.Close()

	dst, err := os.Create(path.Join("./files", name))
	defer dst.Close()
	if err != nil {
		fmt.Println(err, name)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if _, err := io.Copy(dst, f); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		os.Mkdir("files", 0777)
	}
	http.Handle("/", http.FileServer(http.Dir("")))
	http.HandleFunc("/api/upload/", upload)
	if err := http.ListenAndServe(":3456", nil); err != nil {
		panic(err)
	}
}
