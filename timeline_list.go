package twitterscraper

type listEntry struct {
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
	Content   struct {
		CursorType  string `json:"cursorType"`
		EntryType   string `json:"entryType"`
		Value       string `json:"value"`
		Typename    string `json:"__typename"`
		ItemContent struct {
			ItemType         string `json:"itemType"`
			Typename         string `json:"__typename"`
			TweetDisplayType string `json:"tweetDisplayType"`
			TweetResults     struct {
				Result result `json:"result"`
			} `json:"tweet_results"`
		} `json:"itemContent"`
	} `json:"content"`
}

type timelineForList struct {
	Data struct {
		List struct {
			TweetsTimeline struct {
				Timeline struct {
					Instructions []struct {
						Entries []listEntry `json:"entries"`
						Entry   listEntry   `json:"entry"`
						Type    string      `json:"type"`
					} `json:"instructions"`
				} `json:"timeline"`
			} `json:"tweets_timeline"`
		} `json:"list"`
	} `json:"data"`
}

func (timeline *timelineForList) parseTweets() ([]*Tweet, string) {
	var cursor string
	var tweets []*Tweet
	for _, instruction := range timeline.Data.List.TweetsTimeline.Timeline.Instructions {
		for _, entry := range instruction.Entries {
			if entry.Content.CursorType == "Bottom" {
				cursor = entry.Content.Value
				continue
			}
			if entry.Content.ItemContent.TweetResults.Result.Typename == "Tweet" {
				if tweet := entry.Content.ItemContent.TweetResults.Result.parse(); tweet != nil {
					tweets = append(tweets, tweet)
				}
			}
		}
	}
	return tweets, cursor
}
