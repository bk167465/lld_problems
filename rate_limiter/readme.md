# Requirements
- In-process component, allowRequest(key) → boolean
- Key is caller-supplied (user, IP, endpoint, composite — doesn't matter to the limiter)
- Config (limit, window, algorithm) is per-client and updatable at runtime via some config API
- Pluggable algorithm (token bucket, sliding window, etc.)
- Reject-only, no queuing
- Single node, in-memory

# Non functionaal requirements
- Low latency
- Memory bounded
- Extensibility
- Thread safety

Algorithms:
• Token bucket - Bucket size, Refill rate
• Leaking bucket - Bucket size, Outflow rate
• Fixed window counter
• Sliding window log
• Sliding window counter