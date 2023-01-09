package twitter

import (
	"fmt"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := Constructor()
	twitter.PostTweet(1, 5)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Follow(1, 2)
	twitter.PostTweet(2, 6)
	fmt.Println(twitter.GetNewsFeed(1))
	twitter.Unfollow(1, 2)
	fmt.Println(twitter.GetNewsFeed(1))
}
