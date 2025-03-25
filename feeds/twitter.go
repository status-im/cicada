package feeds

// TODO Add Search Feed (instead of user timeline): Use client.Search.Tweets(&twitter.SearchTweetParams{...}).
// TODO Rate limiting handling. Back off if the Twitter API returns a 429.
// TODO Use a database or file cache instead of in-memory seen map for persistence.

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"

	"github.com/dghubble/oauth1"
)

type TwitterFeed struct {
	client     *twitter.Client
	screenName string
	lastSeenID int64
}

func NewTwitterFeed(consumerKey, consumerSecret, accessToken, accessSecret, screenName string) *TwitterFeed {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return &TwitterFeed{
		client:     twitter.NewClient(httpClient),
		screenName: screenName,
	}
}

func (t *TwitterFeed) Name() string {
	return fmt.Sprintf("twitter:%s", t.screenName)
}

func (t *TwitterFeed) FetchItems() ([]FeedItem, error) {
	tweets, _, err := t.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: t.screenName,
		SinceID:    t.lastSeenID,
		Count:      10,
		TweetMode:  "extended",
	})
	if err != nil {
		return nil, err
	}

	var items []FeedItem
	for _, tweet := range tweets {
		if tweet.ID > t.lastSeenID {
			t.lastSeenID = tweet.ID
		}

		timestamp, err := tweet.CreatedAtTime()
		if err != nil {
			return nil, err
		}
		items = append(items, FeedItem{
			ID:        fmt.Sprint(tweet.ID),
			Title:     tweet.FullText,
			Link:      fmt.Sprintf("https://twitter.com/%s/status/%d", t.screenName, tweet.ID),
			Timestamp: timestamp,
		})
	}
	return items, nil
}
