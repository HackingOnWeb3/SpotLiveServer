package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// 解析文件和表单数据
	r.ParseMultipartForm(10 << 20) // 限制为10MB
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// 创建 IPFS shell
	sh := shell.NewShell("https://ipfs.infura.io:5001")

	// 添加文件到 IPFS
	ipfsHash, err := sh.Add(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: https://ipfs.io/ipfs/%s\n", ipfsHash)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/upload", uploadFile).Methods("POST")

	http.ListenAndServe(":8080", router)
}
