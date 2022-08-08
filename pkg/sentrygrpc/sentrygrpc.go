package sentrygrpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor is a grpc interceptor that reports errors and panics
// to sentry. It also sets *sentry.Hub to context.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		hub := sentry.GetHubFromContext(ctx)
		if hub == nil {
			hub = sentry.CurrentHub().Clone()
			ctx = sentry.SetHubOnContext(ctx, hub)
		}

		defer func() {
			if r := recover(); r != nil {
				hub.RecoverWithContext(ctx, r)

				err = status.Errorf(codes.Internal, "%s", r)
			}
		}()

		resp, err = handler(ctx, req)
		if err != nil {
			hub.Scope().SetExtras(map[string]interface{}{
				"grpc.method": info.FullMethod,
				"req":         req,
			})
			hub.Scope().SetUser(sentry.User{
				Email:     "",
				ID:        fmt.Sprint(ctx.Value("userId")),
				IPAddress: "",
				Username:  "",
			})
			statusError, ok := status.FromError(err)
			if ok {
				hub.Scope().SetExtra("grpc.code", statusError.Code())
				hub.CaptureMessage(statusError.Message())

				// enrich the error message with breadcrumbs
				for _, detail := range statusError.Details() {
					switch t := detail.(type) {
					case *errdetails.DebugInfo:
						hub.AddBreadcrumb(&sentry.Breadcrumb{
							Type:      "debug",
							Category:  "server",
							Message:   t.Detail,
							Data:      map[string]interface{}{"stackTrace": strings.Join(t.GetStackEntries(), "")},
							Level:     "error",
							Timestamp: time.Time{},
						}, &sentry.BreadcrumbHint{})
					}
				}
			} else {
				hub.CaptureMessage(fmt.Sprintf("handler error is not compatible with status package: %s", err))
			}
		}
		return resp, err
	}
}
