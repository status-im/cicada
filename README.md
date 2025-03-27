# 🪰👁️ Cicadian

![42ef9cb0-1afc-4e48-9dc2-c7a41f0a0182 (1)](https://github.com/user-attachments/assets/c9392543-264b-4455-bc99-e63c9557b3b6)


Cicadian is a modular, minimalist Waku bot that listens to real-world and on-chain content feeds and relays them into [Waku](https://waku.org/) content topics.

Designed as a rhythmic bridge between decentralised messaging and the ever-changing web-curating signal from RSS, social, and Web3 event sources.

## 🐜 What’s in a Name?

**“Cicadian”** is a mix of **Cicada** + **Circadian**:
- **Cicada** – A chirping insect that appears in cycles, going dormant before waking up to broadcast its signature calls.
- **Circadian** – Refers to biological rhythms repeating roughly every 24 hours.

Much like a cicada’s periodic songs (and aligned with circadian rhythms), Cicadian awakens at set intervals to fetch feeds and “chirp” their updates out over Waku. This captures both the cyclical, scheduled nature of the bot and the idea of periodic broadcasts.

## 🚧 Project Status: Alpha

Cicadian is in **alpha**, meaning its APIs and configurations may still change without notice. We welcome testing, feedback, and contributions, but please be aware of potential breaking changes as the project evolves.

---

## ✨ Features

### ⏱️ Poll Intervals
Feeds are polled on a user-defined schedule. You can choose short intervals, align them with natural daily cycles, or set any custom timing you prefer.

### ⏬ Feed Types
#### ⚙️ Implemented Feed Types
- **RSS/Atom**  
  Includes basic image extraction via enclosures or HTML fallback.
- **Twitter/X**  
  Extracts tweet data and supports images from extended entities.
- **GitHub Releases**  
  Retrieves release notes and metadata.
- **Ethereum Contract Events**  
  Listens for on-chain logs, with timestamps fetched via RPC.
- **Snapshot DAO Proposals**  
  Extracts proposal details and broadcasts them over Waku.

#### 🔮 Planned Feed Types
- **Farcaster**  
  Fetch casts and associated metadata from Farcaster.
- **Bluesky**  
  Pull posts from Bluesky, including text and media attachments.
- **Lens Protocol**  
  Subscribe to Lens via their GraphQL or REST endpoints.
- **Reddit**  
  Fetch posts from subreddits using the official Reddit JSON API.

### 🕸️ Decentralised Relaying
All messages are broadcast over Waku content topics, enabling censorship-resistant, peer-to-peer distribution.

### 🔧 Configurable and Composable
Feed sources are defined in a YAML file, and new feed types or custom logic can be added with minimal effort.

---

## 🧠 Architecture

### Feed Interface
Each feed in Cicadian implements a straightforward interface:

```go
type Feed interface {
    Name() string
    FetchItems() ([]FeedItem, error)
}
```

Within this interface, the feed is responsible for retrieving new items (e.g., tweets, GitHub release notes, or DAO proposals) and returning them as a list of `FeedItem`.

---

### Common Item Structure
Each `FeedItem` contains the essential metadata needed for broadcasting:

```go
type FeedItem struct {
    ID         string
    Title      string
    Link       string
    Timestamp  time.Time
    ImageData  []byte
}
```

- **ID**: A unique identifier for deduplication (e.g., tweet ID, GitHub release tag).
- **Title**: A human-readable title or summary.
- **Link**: A canonical URL pointing to the original source.
- **Timestamp**: The item’s publication time.
- **ImageData**: An optional byte array for embedded images (or other binary attachments).

---

### Waku Publishing
All items are converted into a serialised broadcast format (e.g., `WakuFeedBroadcast` protobuf). These broadcasts contain the `FeedItem` data (including any images) and are published to a configurable Waku **content topic**. This decentralised approach makes it easy for subscribers to retrieve items in near real time, without relying on a central server.

---

### Extensibility
By implementing `Feed` and returning `FeedItem`, you can easily add new data sources. The rest of the system (including Waku publishing, logging, and potential caching/deduplication) remains consistent across feeds. This modular design allows Cicadian to support a wide variety of feed types with minimal additional code.

---

## 🚀 Getting Started

### Requirements
- **Go 1.21+**  
  Ensure you have Go 1.21 or higher installed.
- **Waku Relay Node**  
  You will need access to a local or public Waku Relay node. If none is available, see [go-waku docs](https://docs.status.im/technical/go-waku/) for instructions on running your own node.
- **(Optional) API Tokens**  
  Some feeds (e.g., Twitter, certain Web3 services) require API keys or tokens. Provide them via environment variables or specify them in the configuration if supported.

### Installation
1. Clone the repository:

```shell
 git clone https://github.com/status-im/cicadian.git
 cd cicadian
```

2. Build or run directly:
```shell
go run ./cmd/main.go
```

Or build an executable:
```shell
go build -o cicadian ./cmd/main.go
```

### Configuration
Feeds and other settings are specified in a YAML file (e.g., `feed_config.yaml`). This file controls which feeds are enabled and how often they are polled (when user-defined polling is available).

1. Copy and edit the example:

```shell
cp feed_config.example.yaml feed_config.yaml
```

2. Adjust feed entries:
  - Provide URLs for RSS feeds.
  - Insert credentials (if needed) for Twitter or Ethereum.
  - Enable or disable certain feeds according to your needs.

3. Environment variables:
  - Some APIs (e.g., Twitter, Infura) require keys. Provide them via environment variables or a secrets manager.

Example `feed_config.yaml`:

```yaml
rss:
  - url: "https://example.com/rss"
  - url: "https://another-feed.example.org/feed"

twitter:
  credentials:
    consumer_key_env: "TWITTER_CONSUMER_KEY"
    consumer_secret_env: "TWITTER_CONSUMER_SECRET"
    access_token_env: "TWITTER_ACCESS_TOKEN"
    access_secret_env: "TWITTER_ACCESS_SECRET"
  targets:
    - name: "SomeUser"
      type: "profile"
    - name: "some_search_term"
      type: "search"
    - name: "#SomeHashtag"
      type: "hashtag"

ethereum:
  - rpc_url: "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    contract: "0x1234567890abcdef1234567890abcdef12345678"
    event: "Transfer(address,address,uint256)"
    start_block: 18000000
  - rpc_url: "https://goerli.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    contract: "0xabcdef1234567890abcdef1234567890abcdef12"
    event: "Approval(address,address,uint256)"
    start_block: 9000000
```

### Running Cicadian
Once your config file is in place and any required environment variables are set, start Cicadian:

```shell
go run ./cmd/main.go --config=feed_config.yaml
```

This will:
1. Connect to the specified Waku node.
2. Initialise each feed (e.g., RSS, Twitter, GitHub).
3. Fetch the latest items and broadcast them on Waku at startup and periodically thereafter.

### Verifying Your Setup
1. **Logs**  
   Cicadian will log each feed poll and any broadcast events. If you see errors, verify your config and credentials.
2. **Waku Relay**  
   Check your Waku node logs or use a Waku client to subscribe to the configured content topic. You should see incoming messages.
3. **Feed Output**  
   Confirm items are being broadcast with the correct titles, timestamps, and any optional images.

### Next Steps
- **Customise Poll Intervals**
  When available, configure different polling frequencies in `feed_config.yaml`.
- **Add New Feeds**
  Implement the `Feed` interface for your custom data source and add its config block.
- **Explore the Roadmap**
  See [`roadmap.md`](_docs/roadmap.md) for future plans and features.

If you encounter any problems or have suggestions, feel free to open an [issue](https://github.com/status-im/cicadian/issues).

---

## 🔍 Code & Explanation

Below is a short reference explaining how the **FeedConfig** is read and how feeds are loaded. This reflects the current structure with `config/`, `loader/`, and `feeds/` packages.

### `FeedConfig` (in `config/`)

The `FeedConfig` struct corresponds to the YAML file structure (`feed_config.yaml`). It is parsed by the `Read` function:

```go
package config

import (
    "os"

    "gopkg.in/yaml.v3"
)

type FeedConfig struct {
    Ethereum []EthereumFeedConfig `yaml:"ethereum"`
    RSS      []RSSFeedConfig      `yaml:"rss"`
    Snapshot []SnapshotFeedConfig `yaml:"snapshot"`
    Twitter  TwitterConfig        `yaml:"twitter"`
}

func Read(path string) (FeedConfig, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return FeedConfig{}, err
    }

    var cfg FeedConfig
    err = yaml.Unmarshal(data, &cfg)
    return cfg, err
}
```
	
**Key Points**
- The fields (`Ethereum`, `RSS`, `Snapshot`, `Twitter`) directly map to top-level keys in the YAML.
- `Read` takes the path to a YAML file, unmarshals it, and returns a `FeedConfig`.

---

### Loading Feeds (in `loader/`)

Once the YAML is parsed into a `FeedConfig`, the `Load` function combines feeds from different loaders:

```go
package loader

import (
    "github.com/status-im/cicadian/config"
    "github.com/status-im/cicadian/feeds"
)

func Load(cfg config.FeedConfig) ([]feeds.Feed, error) {
    var all []feeds.Feed

    all = append(all, LoadEthereumFeeds(cfg.Ethereum)...)
    all = append(all, LoadRSSFeeds(cfg.RSS)...)
    all = append(all, LoadSnapshotFeeds(cfg.Snapshot)...)
    all = append(all, LoadTwitterFeeds(cfg.Twitter)...)

    // TODO: Add Reddit feed (JSON parsing)
    // TODO: Add Farcaster feed
    // TODO: Add Bluesky feed
    // TODO: Add Lens Protocol feed

    return all, nil
}
```

**Key Points**
- Functions like `LoadEthereumFeeds` and `LoadRSSFeeds` each take slices of specific config entries (e.g., `[]EthereumFeedConfig`) and return slices of `feeds.Feed`.
- The final `all` slice represents every feed the bot will poll and broadcast.

---

### 🧩 Adding a New Feed

1. **Extend the Config**
  - If your feed needs new YAML fields, create a corresponding struct (e.g. `MyFeedConfig`) and add it to `FeedConfig`.
  - Update `feed_config.schema.yaml` to validate these fields.
  - Provide an example in `feed_config.example.yaml`.

2. **Implement the Feed**
  - In `feeds/`, create a file (e.g. `myfeed.go`) with a `MyFeed` type that implements the `Feed` interface:
      ```go
        type Feed interface {
            Name() string
            FetchItems() ([]FeedItem, error)
        }
      ```
    - Write the logic to fetch or generate items, returning them as a list of `FeedItem`.

3. **Create a Loader**
  - In `loader/`, add a file (e.g. `myfeed.go`) with a function `LoadMyFeeds(cfg []MyFeedConfig) []feeds.Feed`.
  - This function builds `MyFeed` instances from each config entry.

4. **Register in `Load`**
  - Append the returned slice to `all` in the `Load` function:
    ```go
    func Load(cfg config.FeedConfig) ([]feeds.Feed, error) {
		var all []feeds.Feed
		// ...
		all = append(all, LoadMyFeeds(cfg.MyFeed)...)
		return all, nil
	}
    ```

5. **Test & Document**
  - Add tests for fetching/parsing your new feed.
  - Update the README or other docs to explain how users can configure and enable it.

With these steps, any new source (API, web service, blockchain event, etc.) can be integrated into Cicadian’s Waku-based broadcast flow.

---

## 📡 Waku Integration

Cicadian broadcasts all feed items over Waku, using a **content topic** set by the bot admins (e.g., `"kitten-haiku"`). While you can use any string for this field, note that [go-waku](https://github.com/status-im/go-waku) will internally hash it at the pubsub layer to preserve privacy and mitigate collisions.

### Content Topic Configuration
By default, items are published to a single user defined content topic (e.g., `"kitten-haiku"`). You can change this if you want to separate different feeds or categories under their own topics, enabling more granular organisation or filtering of messages.

### Message Serialisation
Each `FeedItem` is serialised into a protobuf message (e.g. `WakuFeedBroadcast`) containing:
- **ID**, **Title**, and **Link**
- **Timestamp** and (optionally) **ImageData**
- Potential for a **Signature** field (planned for future enhancements)

The bot publishes these serialised messages to the configured content topic. Waku nodes then forward them across the network, ensuring that any subscriber who knows the topic name (and the hashed topic ID behind it) can receive your feed items.

### Receiving the Broadcast
Consumers simply subscribe to the topic (or whichever name you configure), parse the incoming protobuf messages, and handle them as needed (e.g., displaying them in a client, forwarding them to another system, etc.). Because Waku is gossip-based, messages are shared among participating nodes, offering a decentralised, censorship-resistant channel for your feed updates.

---

## 🗺️Roadmap

[Roadmap](./_docs/roadmap.md)
