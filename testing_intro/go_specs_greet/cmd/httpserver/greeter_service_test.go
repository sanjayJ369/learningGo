package main_test

import (
	"fmt"
	"go_specs_greet/adapters"
	"go_specs_greet/adapters/httpserver"
	"go_specs_greet/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)

		driver = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath, "httpserver")
	specifications.GreetSpecification(t, driver)
	specifications.CurseSepcification(t, driver)
}
