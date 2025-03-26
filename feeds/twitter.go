package feeds

// TODO Rate limiting handling. Back off if the Twitter API returns a 429.
// TODO Use a database or file cache instead of in-memory seen map for persistence.

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterClient struct {
	client *twitter.Client
}

func NewTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret string) *TwitterClient {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return &TwitterClient{client: twitter.NewClient(httpClient)}
}

func (c *TwitterClient) Client() *twitter.Client {
	return c.client
}

type TwitterProfileFeed struct {
	client     *twitter.Client
	screenName string
	lastSeenID int64
}

func NewTwitterProfileFeed(client *twitter.Client, screenName string) Feed {
	return &TwitterProfileFeed{
		client:     client,
		screenName: screenName,
	}
}

func (t *TwitterProfileFeed) Name() string {
	return fmt.Sprintf("twitter:profile:%s", t.screenName)
}

func (t *TwitterProfileFeed) FetchItems() ([]FeedItem, error) {
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
		ts, _ := tweet.CreatedAtTime()
		items = append(items, FeedItem{
			ID:        fmt.Sprint(tweet.ID),
			Title:     tweet.FullText,
			Link:      fmt.Sprintf("https://twitter.com/%s/status/%d", t.screenName, tweet.ID),
			Timestamp: ts,
		})
	}
	return items, nil
}

//
// Search Feed
//

type TwitterSearchFeed struct {
	client     *twitter.Client
	query      string
	lastSeenID int64
}

func NewTwitterSearchFeed(client *twitter.Client, query string) Feed {
	return &TwitterSearchFeed{client: client, query: query}
}

func (t *TwitterSearchFeed) Name() string {
	return fmt.Sprintf("twitter:search:%s", t.query)
}

func (t *TwitterSearchFeed) FetchItems() ([]FeedItem, error) {
	result, _, err := t.client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      t.query,
		SinceID:    t.lastSeenID,
		Count:      10,
		TweetMode:  "extended",
		ResultType: "recent",
	})
	if err != nil {
		return nil, err
	}

	var items []FeedItem
	for _, tweet := range result.Statuses {
		if tweet.ID > t.lastSeenID {
			t.lastSeenID = tweet.ID
		}
		ts, _ := tweet.CreatedAtTime()
		items = append(items, FeedItem{
			ID:        fmt.Sprint(tweet.ID),
			Title:     tweet.FullText,
			Link:      fmt.Sprintf("https://twitter.com/i/web/status/%d", tweet.ID),
			Timestamp: ts,
		})
	}
	return items, nil
}
