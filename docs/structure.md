mini-kafka/
├── cmd/
│   ├── broker/
│   │   └── main.go
│   ├── producer/
│   │   └── main.go
│   ├── consumer/
│   │   └── main.go
│   └── admin/
│       └── main.go
│
├── internal/
│   ├── broker/
│   │   ├── server.go
│   │   ├── lifecycle.go
│   │   ├── router.go
│   │   ├── handlers.go
│   │   ├── errors.go
│   │   └── limits.go
│   │
│   ├── api/
│   │   └── http/
│   │       ├── server.go
│   │       ├── routes.go
│   │       ├── middleware.go
│   │       ├── produce_handler.go
│   │       ├── fetch_handler.go
│   │       ├── commit_handler.go
│   │       ├── group_handler.go
│   │       └── admin_handler.go
│   │
│   ├── protocol/
│   │   ├── request.go
│   │   ├── response.go
│   │   ├── codec.go
│   │   ├── errors.go
│   │   └── types.go
│   │
│   ├── metadata/
│   │   ├── registry.go
│   │   ├── topic_metadata.go
│   │   └── broker_metadata.go
│   │
│   ├── topic/
│   │   ├── manager.go
│   │   ├── topic.go
│   │   ├── config.go
│   │   └── errors.go
│   │
│   ├── partition/
│   │   ├── partition.go
│   │   ├── manager.go
│   │   ├── assigner.go
│   │   ├── watermark.go
│   │   └── errors.go
│   │
│   ├── storage/
│   │   ├── log/
│   │   │   ├── log.go
│   │   │   ├── segment.go
│   │   │   ├── index.go
│   │   │   ├── reader.go
│   │   │   ├── writer.go
│   │   │   ├── retention.go
│   │   │   └── recovery.go
│   │   ├── filesystem/
│   │   │   ├── layout.go
│   │   │   ├── lock.go
│   │   │   └── cleanup.go
│   │   └── errors.go
│   │
│   ├── producer/
│   │   ├── producer.go
│   │   ├── partitioner.go
│   │   ├── batcher.go
│   │   └── errors.go
│   │
│   ├── consumer/
│   │   ├── fetcher.go
│   │   ├── poller.go
│   │   ├── processor.go
│   │   ├── committer.go
│   │   ├── lag.go
│   │   └── errors.go
│   │
│   ├── offset/
│   │   ├── store.go
│   │   ├── file_store.go
│   │   └── errors.go
│   │
│   ├── group/
│   │   ├── group.go
│   │   ├── membership.go
│   │   ├── assignor.go
│   │   ├── rebalance.go
│   │   └── errors.go
│   │
│   ├── config/
│   │   ├── broker.go
│   │   ├── api.go
│   │   ├── storage.go
│   │   ├── producer.go
│   │   ├── consumer.go
│   │   └── loader.go
│   │
│   ├── observability/
│   │   ├── logger.go
│   │   ├── metrics.go
│   │   ├── health.go
│   │   └── errors.go
│   │
│   └── utils/
│       ├── hashing.go
│       ├── retry.go
│       ├── backoff.go
│       ├── clock.go
│       └── errors.go
│
├── pkg/
│   └── client/
│       ├── client.go
│       ├── producer.go
│       ├── consumer.go
│       ├── admin.go
│       ├── config.go
│       └── errors.go
│
├── configs/
│   ├── broker.yaml
│   ├── producer.yaml
│   ├── consumer.yaml
│   └── admin.yaml
│
├── data/
│   └── broker-1/
│       ├── topics/
│       │   └── orders/
│       │       ├── partition-0/
│       │       │   ├── 00000000000000000000.log
│       │       │   └── 00000000000000000000.index
│       │       └── partition-1/
│       │           ├── 00000000000000000000.log
│       │           └── 00000000000000000000.index
│       └── offsets/
│           └── analytics/
│               ├── orders.partition-0.offset
│               └── orders.partition-1.offset
│
├── docs/
│   ├── architecture.md
│   ├── apis.md
│   ├── storage-format.md
│   ├── consumer-groups.md
│   ├── offsets.md
│   ├── retention.md
│   └── roadmap.md
│
├── scripts/
│   ├── start-broker.sh
│   ├── create-topic.sh
│   ├── run-producer.sh
│   ├── run-consumer.sh
│   └── demo.sh
│
├── tests/
│   ├── unit/
│   │   ├── storage_test.go
│   │   ├── partition_test.go
│   │   ├── offset_test.go
│   │   └── protocol_test.go
│   └── integration/
│       ├── produce_fetch_test.go
│       ├── consumer_resume_test.go
│       └── crash_recovery_test.go
│
├── Makefile
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── README.md
├── LICENSE
└── .gitignore
