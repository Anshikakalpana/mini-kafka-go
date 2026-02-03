mini-kafka/
├── cmd/
│   ├── broker/
│   │   └── main.go             # entrypoint: starts the TCP broker
│   ├── producer/
│   │   └── main.go             # simple CLI producer (for demo & testing)
│   ├── consumer/
│   │   └── main.go             # simple CLI consumer (pull from offset)
│   └── admin/
│       └── main.go             # CLI: create-topic, list-topics, etc.
│
├── internal/
│   ├── broker/
│   │   ├── server.go           # TCP listener + connection handling loop
│   │   ├── handlers.go         # produce, fetch, commit-offset, create-topic logic
│   │   ├── router.go           # maps request type → handler
│   │   └── config.go           # broker config struct + loading
│   │
│   ├── protocol/
│   │   ├── codec.go            # length-prefixed JSON encode/decode
│   │   └── types.go            # Request, Response, Record structs
│   │
│   ├── topic/
│   │   ├── manager.go          # in-memory topic → partitions map + CreateTopic
│   │   └── partition.go        # per-partition state (log + mutex)
│   │
│   ├── storage/
│   │   ├── log.go              # core append-only log: Append, ReadFrom(offset)
│   │   ├── segment.go          # segment file management + rolling
│   │   └── retention.go        # optional: size/time-based cleanup
│   │
│   ├── offset/
│   │   └── store.go            # file-based offset storage per group/topic/partition
│   │
│   └── utils/
│       ├── hashing.go          # key → partition (murmur or simple mod)
│       └── errors.go           # custom error types + wrappers
│
├── configs/
│   └── broker.yaml             # port, data dir, segment size, etc.
│
├── data/                       # runtime: gitignore this
│   └── broker-1/               # (topics/, offsets/ created automatically)
│
├── scripts/
│   └── demo.sh                 # one-command: start broker + produce + consume + restart
│
├── tests/
│   ├── storage_test.go         # log append/read, segment rolling, recovery
│   ├── broker_test.go          # end-to-end: produce → fetch → commit → restart
│   └── recovery_test.go        # (optional but strong) crash simulation
│
├── Dockerfile
├── docker-compose.yml          # (optional) for multi-broker later or easy local run
├── Makefile                    # test, bench, lint, run-broker, etc.
├── README.md                   # architecture, usage, demo, benchmarks
├── go.mod
├── go.sum
└── .gitignore

V1 Features Checklist 

✅ Create topic (fixed partitions)
✅ Produce → returns (partition, offset)
✅ Fetch → from (topic, partition, offset)
✅ Persistent append-only log
✅ Offset commit per consumer group (simple file)

Extra 

✅ segment rolling
✅ retention by size
✅ simple benchmark script (100k msgs)