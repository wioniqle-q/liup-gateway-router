package main

import (
	"fmt"
	"net/http"

	"Liup.Gateway.Golang/src/handlers"
	"Liup.Gateway.Golang/src/middlewares"
	"Liup.Gateway.Golang/src/utils"
	"github.com/gorilla/mux"
	"github.com/unrolled/secure"
)

func main() {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:8080"},
		AllowedHostsAreRegex:    false,
		AllowRequestFunc:        nil,
		HostsProxyHeaders:       []string{"X-Forwarded-Host", "X-Forwarded-Server", "X-Forwarded-For", "X-Forwarded-Proto", "X-Forwarded-Port", "X-Real-Ip"},
		SSLRedirect:             false,
		SSLTemporaryRedirect:    false,
		SSLHost:                 "",
		SSLProxyHeaders:         map[string]string{},
		STSSeconds:              31536000,
		STSIncludeSubdomains:    false,
		STSPreload:              false,
		ForceSTSHeader:          false,
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
		ContentSecurityPolicy:   "default-src 'self'",
		ReferrerPolicy:          "same-origin",
		FeaturePolicy:           "vibrate 'none'; geolocation 'none'; microphone 'none'; camera 'none'; magnetometer 'none'; gyroscope 'none'; speaker 'none'; fullscreen 'none'; payment 'none';",
		PermissionsPolicy:       "fullscreen=(), payment=(), geolocation=(), microphone=(), camera=(), magnetometer=(), gyroscope=(), speaker=(), vibrate=()",
		CrossOriginOpenerPolicy: "same-origin",
		ExpectCTHeader:          "",
		IsDevelopment:           false,
	})

	r := mux.NewRouter()

	r.Use(secureMiddleware.Handler)
	http.Handle("/", r)

	fmt.Println("Registering the request handler at /register")
	r.HandleFunc("/router", handlers.HandleRequest).Methods(http.MethodPost)

	r.Use(middlewares.AuthMiddleware)
	r.Use(middlewares.RecoveryMiddleware)
	r.Use(middlewares.LoggingMiddleware)
	utils.InitLogger()

	fmt.Println("Starting the HTTP server at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println("Error starting the HTTP server:", err)
	}
}
