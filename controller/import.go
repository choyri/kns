package controller

import (
	"fmt"
	"github.com/choyri/kns/service"
	"net/http"
	"strings"
)

func Import(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("读取文件时出现了错误：%s", err.Error()), http.StatusUnprocessableEntity)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	filenames := strings.Split(fileHeader.Filename, ".")
	if len(filenames) == 0 || filenames[len(filenames)-1] != "xlsx" {
		http.Error(w, "文件格式错误，需要 .xlsx 类型", http.StatusBadRequest)
	}

	err = service.ImportFromReader(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
