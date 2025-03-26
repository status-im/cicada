package feeds

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SnapshotFeed struct {
	space    string
	lastSeen time.Time
}

func NewSnapshotFeed(space string) Feed {
	return &SnapshotFeed{space: space}
}

func (s *SnapshotFeed) Name() string {
	return fmt.Sprintf("snapshot:%s", s.space)
}

func (s *SnapshotFeed) FetchItems() ([]FeedItem, error) {
	url := fmt.Sprintf("https://hub.snapshot.org/api/ens/proposals?space=%s", s.space)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		Proposals []struct {
			ID      string `json:"id"`
			Title   string `json:"title"`
			Created int64  `json:"created"`
		} `json:"proposals"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var items []FeedItem
	for _, p := range data.Proposals {
		t := time.Unix(p.Created, 0)
		if t.After(s.lastSeen) {
			s.lastSeen = t
			items = append(items, FeedItem{
				ID:        p.ID,
				Title:     p.Title,
				Link:      fmt.Sprintf("https://snapshot.org/#/%s/proposal/%s", s.space, p.ID),
				Timestamp: t,
			})
		}
	}

	return items, nil
}
