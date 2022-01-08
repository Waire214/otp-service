package routes

import (
	"net/http"
	"otp/interfaces"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRouter(appPort, hostAddress string, otp interfaces.OtpInterface) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: customLogger, NoColor: true})
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Mount("/otp", otpEndpoint(otp))

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(hostAddress+":"+appPort+"/swagger/doc.json"),
	))

	return router
}

func otpEndpoint(otp interfaces.OtpInterface) http.Handler {
	r := chi.NewRouter()
	r.Post("/requests", otp.CreateOtp)
	r.Post("/requests/validate", otp.ValidateOtp)
	return r
}
