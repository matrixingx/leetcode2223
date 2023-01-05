package twitter

import (
	"sort"
	"sync"
)

type IdConstructor struct {
	sync sync.Mutex
	idx int
}

var idc IdConstructor

func (this *IdConstructor) Incr() int {
	this.sync.Lock()
	this.idx++
	res := this.idx
	defer this.sync.Unlock()
	return res
}

type Tweet struct {
	Time    int
	TweetId int
}

type Twitter struct {
	UserFollowMap map[int][]int   // 记录用户关注的人
	UserTweetMap  map[int][]Tweet // 记录用户发的推
	
}

func Constructor() Twitter {
	idc = IdConstructor{
		idx: 0,
	}
	return Twitter{
		UserFollowMap: make(map[int][]int, 0),
		UserTweetMap:  make(map[int][]Tweet, 0),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.UserTweetMap[userId] = append(this.UserTweetMap[userId], Tweet{
		Time: idc.Incr(),
		TweetId: tweetId,
	})
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	var allTweets []Tweet
	allTweets = append(allTweets, this.UserTweetMap[userId]...)
	for _,v := range this.UserFollowMap[userId] {
		allTweets = append(allTweets, this.UserTweetMap[v]...)
	}
	sort.SliceStable(allTweets,func (i,j int) bool {
		return allTweets[i].Time > allTweets[j].Time
	})
	var res []int
	for i,v := range allTweets {
		if i >= 10 {
			break
		}
		res = append(res, v.TweetId)
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	followList := this.UserFollowMap[followerId]
	for _,v := range followList {
		if v == followeeId {
			return
		}
	}
	this.UserFollowMap[followerId] = append(this.UserFollowMap[followerId], followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId == followerId {
		return
	}
	followList := this.UserFollowMap[followerId]
	for i,v := range followList {
		if v == followeeId {
			this.UserFollowMap[followerId] = append(this.UserFollowMap[followerId][:i], this.UserFollowMap[followerId][i+1:]...)
		}
	}
}