package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Odvin/go-live-reloading/internal/application"
)

type WebService struct {
	api application.API
	adr int
	ver string
	env string
}

func Init(api application.API, adr int, ver, env string) *WebService {
	return &WebService{
		api: api,
		adr: adr,
		ver: ver,
		env: env,
	}
}

func (web *WebService) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", web.adr),
		Handler:      web.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
