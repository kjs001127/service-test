package datadog

import (
	"sync"
)

const ddServiceName = "ch-app-store"

var once sync.Once
