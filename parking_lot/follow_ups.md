- Concurrency: What if two cars try to park at the same time?
-- global locks, local locks, mutex

- Nearest entrance: Always allocate the closest spot instead of FIFO.
-- Per gate heap with lazy deletion with help of global occupancy map

- Reservations: Allow reserving a spot for 15 minutes.
-- separte reservations with 15 minute ttl handeled by min heap or cron job

- Electric vehicles: Add EV charging spots.
-- Just a new parking type with parking rule

- Multiple entry/exit gates: How would you scale this?

- Pricing strategy: Weekend pricing, night pricing, dynamic pricing.
-- Strategy pattern and multiple pricing handeled by decorator or rule engine