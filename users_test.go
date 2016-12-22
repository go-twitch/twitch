package twitch

import "testing"

func TestUsersService_GetUserByID(t *testing.T) {
	setup(t)
	user, _, err := client.Users.GetUserByID(testUser)
	if err != nil {
		t.Fatal(err)
	}
	if user.Name != "go_twitch_" {
		t.Fatal("invalied username")
	}
}

func TestUsersService_GetUser(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Users.GetUser()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsersService_GetUserEmotes(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Users.GetUserEmotes(testUser)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsersService_CheckUserSubscriptionByChannel(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Users.CheckUserSubscriptionByChannel(testUser, esl)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUsersService_CheckUserFollows(t *testing.T) {
	setupWithAccess(t)
	list := map[string]bool{
		"dota2ti":   false,
		"dota2ti_2": false,
		"dota2ti_3": false,
		"dota2ti_4": false,
	}
	follows, _, err := client.Users.GetUserFollows(testUser, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, follow := range follows.Follows {
		list[follow.Channel.Name] = true
	}
	for c, v := range list {
		if !v {
			t.Fatalf("why not follow %v", c)
		}
	}
}

func TestUsersService_CheckUserFollowsByChannel(t *testing.T) {
	setupWithAccess(t)
	dotamajor, _, err := client.Users.CheckUserFollowsByChannel(testUser, dota2ti)
	if err != nil {
		t.Fatal(err)
	}
	if dotamajor == nil {
		t.Fatal("why not follow dotamajor")
	}
	riotgames, _, err := client.Users.CheckUserFollowsByChannel(testUser, riotgames)
	if err != nil {
		t.Fatal(err)
	}
	if riotgames != nil {
		t.Fatal("why follow riotgames")
	}
}

func TestUsersService_FollowAndUnfollowChannel(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Users.FollowChannel(testUser, dotamajor, nil)
	if err != nil {
		t.Error("failed follow")
		t.Fatal(err)
	}
	_, err = client.Users.UnfollowChannel(testUser, dotamajor)
	if err != nil {
		t.Error("failed unfollow")
		t.Fatal(err)
	}
}

func TestUsersService_GetUserBlockList(t *testing.T) {
	setupWithAccess(t)
	list := map[string]bool{
		"esl_csgo":  false,
		"esl_dota2": false,
		"esl_lol":   false,
	}
	blocks, _, err := client.Users.GetUserBlockList(testUser, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, block := range blocks.Blocks {
		list[block.User.Name] = true
	}
	for c, v := range list {
		if !v {
			t.Fatalf("why not block %v", c)
		}
	}
}

func TestUsersService_BlockAndUnblockUser(t *testing.T) {
	setupWithAccess(t)
	_, _, err := client.Users.BlockUser(testUser, playparagon)
	if err != nil {
		t.Error("failed block")
		t.Fatal(err)
	}
	_, err = client.Users.UnblockUser(testUser, playparagon)
	if err != nil {
		t.Error("failed unblock")
		t.Fatal(err)
	}
}
