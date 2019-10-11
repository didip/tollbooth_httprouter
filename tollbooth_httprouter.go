package tollbooth_httprouter

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/julienschmidt/httprouter"
)

// RateLimit is a rate limiting middleware
func LimitHandler(handler httprouter.Handle, lmt *limiter.Limiter) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		httpError := tollbooth.LimitByRequest(lmt, w, r)
		if httpError != nil {
			lmt.ExecOnLimitReached(w, r)
			w.Header().Add("Content-Type", lmt.GetMessageContentType())
			w.WriteHeader(httpError.StatusCode)
			w.Write([]byte(httpError.Message))
			return
		}

		handler(w, r, ps)
	}
}
