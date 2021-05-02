package faceeasco

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var (
	Version = "0.0.1"
	Debug   = false
)

var defaultConnectTimeout = 5 * time.Second
var defaultReadTimeout = 10 * time.Second

var DefaultUserAgent = fmt.Sprintf("Faceeasco (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

type Faceeasco struct {
	httpClient *http.Client
	config     *Config
}
