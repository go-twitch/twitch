package twitch

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelFeedService_GetMultipleFeedPost(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.ChannelFeed.GetMultipleFeedPosts(testUser, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestChannelFeedService_CreateAndDeleteFeedPost(t *testing.T) {
	setupWithAccess(t)
	content := fmt.Sprintf("testing %v", time.Now())
	res, _, err := client.ChannelFeed.CreateFeedPost(testUser, content, nil)
	if err != nil {
		t.Error("failed create feed")
		t.Fatal(err)
	}
	if res.Post.Body != content {
		t.Fatal("posted invalied content feed")
	}
	d, _, err := client.ChannelFeed.DeleteFeedPost(testUser, res.Post.ID)
	if err != nil {
		t.Error("failed delete feed")
		t.Fatal(err)
	}
	if !d.Deleted {
		t.Fatal("feed is not deleted")
	}
	if d.ID != res.Post.ID {
		t.Fatal("deleted other feed")
	}
}
