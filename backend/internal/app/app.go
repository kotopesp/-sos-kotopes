package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/kotopesp/sos-kotopes/internal/controller/http/model/validator"
	"github.com/kotopesp/sos-kotopes/internal/migrate"
	rolesService "github.com/kotopesp/sos-kotopes/internal/service/role"
	usersService "github.com/kotopesp/sos-kotopes/internal/service/user"
	vetservice "github.com/kotopesp/sos-kotopes/internal/service/vet"

	"github.com/kotopesp/sos-kotopes/config"
	v1 "github.com/kotopesp/sos-kotopes/internal/controller/http"
	"github.com/kotopesp/sos-kotopes/internal/core"
	"github.com/kotopesp/sos-kotopes/internal/service/auth"
	rolesStore "github.com/kotopesp/sos-kotopes/internal/store/role"
	"github.com/kotopesp/sos-kotopes/internal/store/user"
	userFavouriteStore "github.com/kotopesp/sos-kotopes/internal/store/userfavourite"
	vetstore "github.com/kotopesp/sos-kotopes/internal/store/vet"
	vetreviewsstore "github.com/kotopesp/sos-kotopes/internal/store/vet_review"
	"github.com/kotopesp/sos-kotopes/pkg/logger"
	"github.com/kotopesp/sos-kotopes/pkg/postgres"

	commentservice "github.com/kotopesp/sos-kotopes/internal/service/comment"
	postservice "github.com/kotopesp/sos-kotopes/internal/service/post"
	animalstore "github.com/kotopesp/sos-kotopes/internal/store/animal"
	commentstore "github.com/kotopesp/sos-kotopes/internal/store/comment"
	poststore "github.com/kotopesp/sos-kotopes/internal/store/post"
	postfavouritestore "github.com/kotopesp/sos-kotopes/internal/store/postfavourite"
	refreshsessionstore "github.com/kotopesp/sos-kotopes/internal/store/refresh_session"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	ctx := context.Background()

	// Init logger
	logger.New(cfg.Log.Level)

	// Postgres connection
	pg, err := postgres.New(ctx, cfg.DB.URL)
	if err != nil {
		logger.Log().Fatal(ctx, "error with connection to database: %s", err.Error())
	}
	defer pg.Close(ctx)

	// Migrate up
	if err := migrate.Up(cfg.DB.URL); err != nil {
		logger.Log().Fatal(ctx, "error with up migrations for database: %s", err.Error())
	}

	// Stores
	userStore := user.New(pg)
	commentStore := commentstore.New(pg)
	roleStore := rolesStore.New(pg)
	favouriteUserStore := userFavouriteStore.New(pg)
	postStore := poststore.New(pg)
	postFavouriteStore := postfavouritestore.New(pg)
	animalStore := animalstore.New(pg)
	refreshSessionStore := refreshsessionstore.New(pg)
	vetStore := vetstore.New(pg)
	vetReviewsStore := vetreviewsstore.New(pg)

	// Services
	commentService := commentservice.New(
		commentStore,
		postStore,
	)
	roleService := rolesService.New(roleStore, userStore)
	userService := usersService.New(userStore, favouriteUserStore)
	authService := auth.New(
		userStore,
		refreshSessionStore,
		core.AuthServiceConfig{
			JWTSecret:            cfg.JWTSecret,
			VKClientID:           cfg.VKClientID,
			VKClientSecret:       cfg.VKClientSecret,
			VKCallback:           cfg.VKCallback,
			AccessTokenLifetime:  cfg.AccessTokenLifetime,
			RefreshTokenLifetime: cfg.RefreshTokenLifetime,
		},
	)
	postService := postservice.New(postStore, postFavouriteStore, animalStore, userStore)
	vetService := vetservice.New(vetStore, vetReviewsStore, userStore)

	// Validator
	formValidator := validator.New(ctx, baseValidator.New())
	// HTTP Server
	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		StrictRouting:            false,
		EnableSplittingOnParsers: true,
	})
	app.Use(recover.New())

	// This configuration is necessary so that the frontend can send requests with cookies.
	app.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowedOrigins,
		AllowCredentials: true,
		AllowHeaders:     "*",
	}))

	v1.NewRouter(
		app,
		authService,
		commentService,
		postService,
		userService,
		roleService,
		vetService,
		formValidator,
	)

	logger.Log().Info(ctx, "server was started on %s", cfg.HTTP.Port)
	err = app.Listen(cfg.HTTP.Port)
	if err != nil {
		logger.Log().Fatal(ctx, "server was stopped: %s", err.Error())
	}

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Log().Info(ctx, "signal %s received", s.String())
	case <-ctx.Done():
		return
	}

	// Shutdown
	err = app.Shutdown()
	if err != nil {
		logger.Log().Fatal(ctx, "error with shutdown server: %s", err.Error())
	}
}
