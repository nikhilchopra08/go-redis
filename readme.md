# GoRedis

A Redis-compatible, in-memory key-value store built from scratch in Go — implementing the RESP (Redis Serialization Protocol) wire protocol, AOF (Append-Only File) persistence, and concurrent client handling.

> Built as a deep-dive into distributed systems internals: protocol design, persistence strategies, and high-throughput concurrent I/O in Go.

---

## Features

- **RESP Protocol** — Full implementation of the Redis Serialization Protocol for wire-compatible communication
- **AOF Persistence** — Append-Only File logging with replay on startup for durable state recovery
- **Concurrent Connections** — Goroutine-per-client architecture with safe concurrent access to the in-memory store
- **Redis CLI Compatible** — Works with `redis-cli` and standard Redis clients out of the box

---

## Benchmarks

Measured using `redis-benchmark` against a local GoRedis instance:

| Operation | Throughput        |
|-----------|-------------------|
| `SET`     | ~73,000 ops/sec   |
| `GET`     | ~138,000 ops/sec  |

---

## Supported Commands

| Command              | Description                        |
|----------------------|------------------------------------|
| `SET key value`      | Set a key to a string value        |
| `GET key`            | Get the value of a key             |
| `PING`               | Test connection liveness           |

---

## Getting Started

### Prerequisites

- Go 1.21+

### Run

```bash
git clone https://github.com/nikhilchopra08/go-redis
cd goredis
go run *.go
```

The server starts on port `6379` by default.

### Connect via redis-cli

```bash
redis-cli -p 6379
```
---

## Architecture

```
client (TCP)
     │
     ▼
RESP Parser          ← parses incoming commands from raw bytes
     │
     ▼
Command Handler      ← dispatches to SET / GET / DEL / etc.
     │
     ▼
In-Memory Store      ← concurrent map with mutex/RWMutex guards
     │
     ▼
AOF Writer           ← appends each write command to disk
```

On startup, GoRedis replays the AOF log to restore state before accepting connections.
---

## What I Learned

- Implementing a binary-safe wire protocol (RESP) from raw TCP bytes
- Designing safe concurrent access patterns in Go (sync.RWMutex)
- AOF persistence trade-offs: durability vs. write amplification
- Debugging subtle concurrency bugs: premature replay loop exits, single-connection bottlenecks

---