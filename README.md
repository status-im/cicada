# 🪰👁️ Cicadian

![42ef9cb0-1afc-4e48-9dc2-c7a41f0a0182 (1)](https://github.com/user-attachments/assets/c9392543-264b-4455-bc99-e63c9557b3b6)


Cicadian is a modular, minimalist Waku bot that listens to real-world and on-chain content feeds and relays them into [Waku](https://waku.org/) content topics.

Designed as a rhythmic bridge between decentralised messaging and the ever-changing web-curating signal from RSS, social, and Web3 event sources.

---

## ✨ Features

- ⏳ **Cicadian rhythm**: Feeds poll on interval, in sync with natural or app-defined rhythms  
- 🪐 **Pluggable feed types**:
  - RSS
  - Twitter/X
  - GitHub Releases
  - Ethereum contract events
  - Farcaster casts
  - Snapshot DAO proposals
  - Bluesky feeds
- 🕸️ **Decentralised relaying**: All messages are published to Waku content topics
- ⚙️ **Configurable and composable**: Easily add new feed types or swap out logic

---

## 🧠 Architecture

Each feed implements a simple interface:

```go
type Feed interface {
    Name() string
    FetchItems() ([]FeedItem, error)
}
```

Each `FeedItem` is published to Waku using a common publisher:

```go
type FeedItem struct {
    ID        string
    Title     string
    Link      string
    Timestamp time.Time
}
```

Messages are serialized and published to a configured content topic.

---

## 🚀 Getting Started

### Requirements

- Go 1.21+
- Waku Relay node or access to a public Waku network
- Optional: API tokens for social/Web3 integrations

### Run the Bot

```bash
go run cmd/main.go
```

### Configure Feeds

Feeds are registered manually in `main.go` for now. Example:

```go
feeds := []Feed{
    NewRSSFeed("https://example.com/rss"),
    NewTwitterFeed(...),
    NewEthereumEventFeed(...),
}
```

---

## 🧩 Adding a New Feed

1. Create a new file in `feeds/`
2. Implement the `Feed` interface
3. Register it in `main.go`

> Feed sources can be anything: APIs, logs, webhooks, or on-chain data.

---

## 📡 Waku Integration

Feeds are published to a specific Waku content topic, such as:

```
/cicadian/1/feed/proto
```

Messages are serialized as plaintext or protobuf.

---

## 🐞 TODO / Roadmap

- [ ] Config file or CLI for runtime configuration
- [ ] Persistent state tracking per feed
- [ ] Dynamic topic routing per feed type
- [ ] UI or dashboard integration
- [ ] Waku message formatting standard

---

## 🪲 Inspiration

Named after the cicada, nature’s rhythmic broadcaster. Circadian echoes Waku signals through time and the network. The logo and theme reflect cycles, decentralisation, and quiet (sometimes at least) resilience.


