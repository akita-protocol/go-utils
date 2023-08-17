package server

import (
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"gitlab.algotoolchain.com/keelakay/utils/db"
	"gitlab.algotoolchain.com/keelakay/utils/errors"
	"gitlab.algotoolchain.com/keelakay/utils/natswrapper"
	"gitlab.algotoolchain.com/keelakay/utils/uuid"
)

const (
	// TimeFormat is the server format for time.Now().Format()
	TimeFormat = "2006-01-02 15:04:05"
	// TimeZone is the server time zone
	TimeZone = "America/Los_Angeles"

	pkg errors.Pkg = "server"
)

func New() *GenericServer {
	const op errors.Op = "New"
	var err error
	s := &GenericServer{
		Engine: gin.Default(),
	}

	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN != "" {
		s.Sentry, err = s.SentryInit(sentry.ClientOptions{
			Dsn:         os.Getenv("SENTRY_DSN"),
			Environment: s.Getenv(),
		})
		if err != nil {
			panic(err.Error())
		}

		s.Use(sentrygin.New(sentrygin.Options{
			Repanic: true,
		}))
	}

	s.Use(RequestID())
	s.Use(CORSMiddleware())

	conn, err := db.Connect()
	if err != nil {
		err := errors.E(pkg, op, errors.Database, err, "failure to connect to database")
		s.Sentry.CaptureException(err, nil, nil)
		panic(err)
	}
	s.DB = conn

	s.LocalTime, err = time.LoadLocation(TimeZone)
	if err != nil {
		err := errors.E(pkg, op, errors.Database, err, "failure to to load local time")
		s.Sentry.CaptureException(err, nil, nil)
		panic(err)
	}

	return s
}

type GenericServer struct {
	*gin.Engine
	Name      errors.Sn
	DB        *sqlx.DB
	LocalTime *time.Location
	Sentry    *sentry.Client
	Nats      *natswrapper.NatsWrapper
}

func (s *GenericServer) SentryInit(options sentry.ClientOptions) (*sentry.Client, error) {
	const op errors.Op = "SentryInit"
	hub := sentry.CurrentHub()
	client, err := sentry.NewClient(options)
	if err != nil {
		return nil, errors.E(pkg, op, err)
	}
	hub.BindClient(client)
	return client, nil
}

// NatsConnect connects to the nats cluster
func (s *GenericServer) NatsConnect(namespace natswrapper.Namespace) {
	const op errors.Op = "server.NatsConnect"

	nc, err := nats.Connect("nats.nats", nats.Timeout(10*time.Second))
	if err != nil {
		err := errors.E(op, err, "failure to connect to nats cluster")
		panic(err.Error())
	}

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		err := errors.E(op, err, "failure to setup encoded connection")
		panic(err.Error())
	}

	s.Nats = natswrapper.New(ec, s.Getenv(), namespace)
}

// Close cleanly exits the server
func (s *GenericServer) Close() {
	sentry.Flush(2 * time.Second)
	s.DB.Close()

	if s.Nats != nil {
		for _, sub := range s.Nats.Subs {
			sub.Unsubscribe()
		}

		s.Nats.Close()
	}
}

// IsProduction returns a boolean for whether the request wants to use product or not
func (s *GenericServer) IsProduction() bool {
	env := s.Getenv()
	if env == "production" {
		return true
	}
	return false
}

// Getenv returns the environment the server is running in
func (s *GenericServer) Getenv() string {
	return os.Getenv("GITLAB_ENVIRONMENT_NAME")
}

// GetDatabaseName uses the gin context production header to determine which database to connect to
func (s *GenericServer) GetDatabaseName(c *gin.Context) string {
	prod := c.GetHeader("Production")
	if prod == "true" {
		return "advisor_core"
	}
	return "advisor_test"
}

// RequestID generates a request id to attach our logs/event chains to
// it also verifies that the request came from the headset advisor domain
// and it responds to our option requests
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUUID := uuid.New(uuid.Request)
		c.Set("REQUEST_ID", requestUUID)
		c.Next()
	}
}

// CORSMiddleware for CORS requests
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
