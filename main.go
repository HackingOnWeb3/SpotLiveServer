package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	if r.Method == "OPTIONS" {
		//return 200 and ok
		w.WriteHeader(http.StatusOK)
		return
	}

	// 解析文件和表单数据
	r.ParseMultipartForm(10 << 20) // 限制为10MB
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//get filename from request body
	filename := r.FormValue("filename")

	if filename == "" {
		http.Error(w, "no filename name", http.StatusInternalServerError)
		return
	}
	// 检查目录是否存在，如果不存在则创建
	dir := fmt.Sprintf("./%s/", filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 创建目标文件
	dst, err := os.Create(dir + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将上传的文件复制到目标位置
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully\n")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/upload", uploadFile).Methods("POST", "OPTIONS")

	//ignore core is
	http.ListenAndServe(":8080", router)
}
