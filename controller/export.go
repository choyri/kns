package controller

import (
	"encoding/json"
	"fmt"
	"github.com/choyri/kns/service"
	"net/http"
	"os"
)

func Export(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var (
		err  error
		data struct {
			IDs []uint `json:"ids"`
		}
		file *os.File
	)

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, fmt.Sprintf("读取 body 错误：%s", err.Error()), http.StatusBadRequest)
		return
	}

	file, err = service.Export(data.IDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		_ = file.Close()
	}()

	fi, _ := file.Stat()

	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), file)
}
