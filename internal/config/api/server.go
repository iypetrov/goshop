package api

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth/v5"
	"github.com/iypetrov/goshop/web"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/iypetrov/goshop/internal/config"
	"github.com/iypetrov/goshop/internal/users"
)

type Server struct{}

func NewServer(ctx context.Context, conn *pgxpool.Pool) *http.Server {
	s := &Server{}

	// init repositories
	users.NewRepository(ctx, conn)

	// init services
	users.NewServer(users.GetRepository())

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", config.Get().App.Port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
}

func (s *Server) RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	users.GetServer().InitAuthProviders()

	r.Mount("/", web.Router())
	r.Route(fmt.Sprintf("/api/v%s", config.Get().App.Version), func(r chi.Router) {
		r.Use(apiVersionCtx(config.Get().App.Version))
		// Public Routes
		r.Group(func(r chi.Router) {
			r.Mount("/auth", users.AuthRouter())
			r.Mount("/onboarding", users.OnboardingRouter())
		})
		// Private Routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(users.GetTokenAuth()))
			r.Use(jwtauth.Authenticator(users.GetTokenAuth()))
			r.Mount("/users", users.Router())

		})
		// Admin Routes
		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(users.GetTokenAuth()))
			r.Use(jwtauth.Authenticator(users.GetTokenAuth()))
			r.Use(users.AdminAuthenticator())
			r.Mount("/admin/users", users.AdminRouter())
		})
	})

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte{})
		if err != nil {
			return
		}
	})

	docgen.PrintRoutes(r)

	return r
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}
