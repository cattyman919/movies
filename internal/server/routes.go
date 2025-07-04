package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// ErrorHandler captures errors and returns a consistent JSON error response
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Step1: Process the request first.

		// Step2: Check if any errors were added to the context
		if len(c.Errors) > 0 {
			// Step3: Use the last error
			err := c.Errors.Last().Err

			// Step4: Respond with a generic error message
			c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"error":   err.Error(),
			})
		}

		// Any other steps if no errors are found
	}
}

// Controller
func (s *Server) RegisterRoutes() http.Handler {

	r := gin.Default()

	r.Use(ErrorHandler())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	// Register Handlers

	r.GET("/health", s.healthHandler)
	api := r.Group("/api")
	api.GET("/movies/now_playing", s.tmdb.NowPlaying)
	api.GET("/get_guest_session", s.tmdb.GetGuestSessionId)

	return r
}

func (s *Server) healthHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Healthy"
	c.JSON(http.StatusOK, resp)
}
