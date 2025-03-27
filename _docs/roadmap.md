# Cicadian Roadmap

This document outlines upcoming features, improvements, and integrations for **Cicadian**. Priorities may shift based on community feedback and project needs.

---

## 1. Twitter Enhancements

### Twitter List Support
- **Goal**: Add the ability to fetch and broadcast tweets from specific Twitter Lists.
- **Tasks**:
    - Extend `TwitterFeed` (or create `TwitterListFeed`) to handle list IDs.
    - Use Twitter’s list/tweets API to retrieve posts.
    - Convert tweets into `FeedItem` objects, handle pagination if necessary.

### API Rate Limiting & Back-Off
- **Goal**: Properly handle rate limits and avoid 429 (Too Many Requests) errors.
- **Tasks**:
    - Detect 429 responses and implement exponential back-off or a token bucket approach.
    - Log rate-limit triggers and display comprehensive warnings or retry information.

### Persistent “Seen” Cache
- **Goal**: Replace in-memory caching of tweet IDs with a small database or file-based approach.
- **Tasks**:
    - Implement a store (e.g., SQLite, BoltDB, or JSON file).
    - Record previously broadcast tweets to avoid duplicates.
    - Load “seen” data at start-up, persist new entries as they arrive.

---

## 2. Ethereum Logs Improvements

### Block Timestamp Caching
- **Goal**: Reduce repeated RPC calls by caching block numbers and their corresponding timestamps.
- **Tasks**:
    - Maintain a map of `{blockNumber: timestamp}` entries.
    - If a block has been seen, reuse the cached timestamp.
    - Consider pruning or limiting the cache’s size for memory efficiency.

---

## 3. New Feeds & Protocols

### Reddit
- **Goal**: Fetch posts from a chosen subreddit (e.g., `/r/<subreddit>/new.json`) and broadcast them.
- **Tasks**:
    - Parse post content (titles, links, timestamps) and extract images if available.
    - Handle potential rate limits, private subreddits, or errors from Reddit.
    - Prevent duplicates through caching.

### Farcaster
- **Goal**: Integrate with Farcaster’s API or aggregator endpoints to fetch user feeds or casts.
- **Tasks**:
    - Create a `FarcasterFeed` type.
    - Parse retrieved data into `FeedItem` objects, including images or media if present.
    - Include any necessary authentication or token usage.

### Bluesky
- **Goal**: Implement a `BlueskyFeed` to pull posts from Bluesky’s protocol or API.
- **Tasks**:
    - Investigate available official or community APIs.
    - Map text, images, and other media into `FeedItem` objects.
    - Handle pagination or streaming updates if supported.

### Lens Protocol
- **Goal**: Subscribe to Lens Protocol updates via their GraphQL or REST interfaces.
- **Tasks**:
    - Implement a `LensFeed` to fetch and parse posts into `FeedItem` objects.
    - Process any embedded images, timestamps, or links.
    - Integrate caching and deduplication.

---

## 4. Signature Support

### Sender Signature
- **Goal**: Use cryptographic signatures to verify authenticity of each broadcast.
- **Tasks**:
    - Sign each `WakuFeedBroadcast` with a private key (e.g., Ethereum).
    - Store the signature in `sender_signature`.
    - (Optional) Provide a verification helper or guide for consumers.

---

## 5. Caching & Deduplication

### Persistent State Tracking
- **Goal**: Standardise how the system remembers and avoids broadcasting the same item multiple times.
- **Tasks**:
    - Create a unified caching interface that supports each feed.
    - Store “seen” item IDs/hashes in a database or file.
    - Ensure concurrency safety if multiple feeds or goroutines access the same store.

---

## 6. Poll Intervals

### User-Defined Poll Intervals
- **Goal**: Allow custom scheduling so users can choose how frequently each feed is polled.
- **Tasks**:
    - Introduce configuration settings (e.g., in `feed_config.yaml`) for defining poll intervals per feed.
    - Implement a scheduling mechanism that respects these intervals without overlap.
    - Ensure changes can be made at runtime or via restart with new config.

---

## 7. Configuration & CLI

- **Goal**: Offer flexible runtime configuration and command-line options.
- **Tasks**:
    - Create or expand a configuration file (`.yaml`, `.toml`, or `.json`) for feed settings and Waku details.
    - Provide CLI flags to override or supply specific settings (API keys, feed URLs, poll intervals, etc.).
    - Validate configuration on start-up and document usage examples.

---

## 8. Dynamic Topic Routing

- **Goal**: Allow different feed types or categories to broadcast over distinct Waku topics.
- **Tasks**:
    - Extend the feed configuration (e.g., `feed_config.yaml`) to specify a custom Waku topic.
    - On broadcast, publish to the configured topic via `go-waku`.
    - Enable consumers to selectively subscribe to specific topics (e.g., “GitHub feed” vs “Twitter feed”).

---

## 9. UI or Dashboard Integration

- **Goal**: Provide a simple interface or dashboard to inspect the status of each feed, last broadcast times, and any errors.
- **Tasks**:
    - (Optional) Set up a minimal HTTP server that displays feed statuses, item counts, error logs, etc.
    - Consider integrating basic metrics for monitoring (e.g., success rates, average fetch times).

---

## 10. Waku Message Formatting Standard

- **Goal**: Define or adopt a standard structure for messages sent over Waku.
- **Tasks**:
    - Ensure `WakuFeedBroadcast` includes all necessary fields (ID, timestamp, link, image data, etc.).
    - Consistently include optional fields (e.g., `sender_signature`) when available.
    - Document the format so third-party tools can parse it reliably.

---

## 11. Documentation & Testing

- **Goal**: Ensure reliability, clarity, and ease of use.
- **Tasks**:
    - Expand unit and integration tests for both existing and new feeds.
    - Document environment variables, configuration files, and feed usage in the README or separate docs.
    - Provide clear examples or a quick-start guide for each feed type.

---

_We welcome feedback and contributions! If you have suggestions or encounter issues, please open an [issue](https://github.com/status-im/cicadian/issues) or submit a pull request._
