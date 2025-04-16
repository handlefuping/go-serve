package log

import (
	"io"
	stLog "log"
	"net/http"
	"os"
)

var log *stLog.Logger

type fileLog string
func (fl fileLog) Write(p []byte) (int, error)  {
	f, err := os.OpenFile(string(fl), os.O_APPEND | os.O_CREATE | os.O_RDWR, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(p)
}

func Run (des string)  {
	log = stLog.New(fileLog(des), "go-> ", stLog.LstdFlags)
}

func LogHandler () {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost: 
			msg, err := io.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
