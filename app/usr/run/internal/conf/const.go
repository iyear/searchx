package conf

import "time"

const (
	ContextScope = "scope"
)

// telegram client options
const (
	RateInterval = 350 * time.Millisecond
	RateBucket   = 3
	MaxRetries   = 15
)
