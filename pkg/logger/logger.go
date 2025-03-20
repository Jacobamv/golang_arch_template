package logger

import (
	"io"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type TracingHook struct{}

// TracingHook
func (h TracingHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()

	var requestID string

	requestID, _ = ctx.Value("request_id").(string)

	e.Str("request_id", requestID)
}

// InitLogger ...
func InitLogger(params Dependencies) (log zerolog.Logger) {
	var logger zerolog.Logger
	if params.Config.Logger.Type == "file" {
		// Required to get error stack details
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		f, err := os.OpenFile(params.Config.Logger.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
		if err != nil {

			panic(err)
		}

		// Associate the open file with the recording source, which will be
		// registered in the zerolog
		var outputPull io.Writer = f

		// Get the used version of golanag from runtime
		logger = zerolog.New(outputPull)
	} else if params.Config.Logger.Type == "stdout" {
		logger = zerolog.New(os.Stdout)
	} else {
		panic("No log type")
	}

	buildInfo, _ := debug.ReadBuildInfo()

	logger = logger.
		Level(zerolog.Level(zerolog.DebugLevel)).
		With().
		Caller().
		Stack().
		Timestamp().
		Str("go_version", buildInfo.GoVersion).
		Str("service", "example").
		Logger()

	logger = logger.Hook(TracingHook{})

	return logger
}
