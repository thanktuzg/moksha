package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"gits.tcbs.com.vn/com/pkg/httputils"
)

var port = "9999"

func main() {
	log.Info("Hello, moksha service is running!")

	log.Fatal(Run())
}

func Run() error {
	r := mux.NewRouter()

	r.HandleFunc("/post", handleHome).Methods("POST")

	http.Handle("/", r)
	log.Infof("moksha service is running at port %s", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	value, err := retryReturnValue(3, time.Duration(2) * time.Second, func() (string, error) {
		var a string
		var err error
		err = nil
		a = "123456"
		if err != nil {
			log.Infof("retry")
			return "", errors.New("aa")
		}
		log.Infof("DDDD %s", a)
		return a, nil
	})

	if err != nil {
		log.WithError(err).Errorf("AAAA")
	} else {
		log.Infof("Value is %s", value)
	}

	log.Infof("LAST %v %v", value, err)

	httputils.RespondNoContent(w, 200)
}

func retryReturnValue(attempts int, sleep time.Duration, f func() (string, error)) (string, error) {
	value, err := f()
	if err != nil {
		for i := 1; i < attempts; i++ {
			v, err := f()
			if err == nil {
				return v, nil
			}
			time.Sleep(sleep)
		}
	}
	return value, err
}
