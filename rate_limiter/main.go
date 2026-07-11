package main

import (
	"errors"
	"time"
)

type RateLimiterAlgorithm interface {
	Name() string
	Allow(key string, config any, stateStore map[string]any) (bool, error)
}

type TokenBucket struct{}

type TokenBucketConfig struct {
	windowLength time.Duration
	bucketSize   int
}

type TokenBucketState struct {
	lastRequestTS time.Time
	tokenLeft     int
}

func (t *TokenBucket) Name() string {
	return "token_bucket"
}

func (t *TokenBucket) Allow(key string, config any, stateStore map[string]any) (bool, error) {
	parsedConfig, ok := config.(TokenBucketConfig)
	if !ok {
		return false, errors.New("Encountered errorr white config type assertion")
	}

	var parsedState TokenBucketState
	state := stateStore[key]

	if state == nil {
		parsedState = TokenBucketState{
			time.Now(),
			parsedConfig.bucketSize - 1,
		}
		stateStore[key] = parsedState
		return true, nil
	}

	parsedState, ok = state.(TokenBucketState)
	if !ok {
		return false, errors.New("Encountered errorr white state type assertion")
	}

	updateTockenBucketState(parsedConfig, &parsedState)

	if parsedState.tokenLeft > 0 {
		parsedState.tokenLeft--
		parsedState.lastRequestTS = time.Now()
		stateStore[key] = parsedState
		return true, nil
	}

	return false, nil
}

func updateTockenBucketState(config TokenBucketConfig, state *TokenBucketState) {
	state.tokenLeft += int(time.Since(state.lastRequestTS) / config.windowLength)

	if state.tokenLeft > config.bucketSize {
		state.tokenLeft = config.bucketSize
	}
}

type RateLimiter struct {
	configStore map[string]any
	stateStore  map[string]any
	algo        RateLimiterAlgorithm
}

func NewRateLimiter(algo string, windowSize time.Duration, bucketSize int) (*RateLimiter, error) {
	switch algo {
	case "token_bucket":
		return &RateLimiter{
			configStore: make(map[string]any),
			stateStore:  make(map[string]any),
			algo:        &TokenBucket{},
		}, nil
	default:
		return nil, errors.ErrUnsupported
	}
}

func (r *RateLimiter) AllowRequest(key string) (bool, error) {
	config, ok := r.configStore[key]
	if !ok {
		return false, errors.New("Config missing for the key")
	}

	return r.algo.Allow(key, config, r.stateStore)
}

func (r *RateLimiter) UpdateConfig(key string, c any) {
	r.configStore[key] = c
}
