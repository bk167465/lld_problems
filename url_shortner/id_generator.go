package main

import "sync/atomic"

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type IdGenerator interface {
	generate() string
}

type Base62IDGenrator struct {
	counter uint64
}

func NewBase62IDGenerator() Base62IDGenrator {
	return Base62IDGenrator{0}
}

func (m Base62IDGenrator) generate() string {
	id := atomic.AddUint64(&m.counter, 1)
	return encodeBase62(id)
}

func encodeBase62(num uint64) string {
	if num == 0 {
		return "0"
	}

	var result []byte

	for num > 0 {
		remainder := num % 62
		result = append(result, base62Chars[remainder])
		num /= 62
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
