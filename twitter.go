package twitter

import (
	"net/http"

	"github.com/dghubble/sling"
)

const twitterAPI = "https://api.twitter.com/2/"

// Client is a Twitter client for making Twitter API requests.
type Client struct {
	sling *sling.Sling
	// Twitter API Services
	Accounts       *AccountService
	Blocks         *BlockService
	DirectMessages *DirectMessageService
	Favorites      *FavoriteService
	Followers      *FollowerService
	Friends        *FriendService
	Friendships    *FriendshipService
	Lists          *ListsService
	RateLimits     *RateLimitService
	Search         *SearchService
	PremiumSearch  *PremiumSearchService
	Statuses       *StatusService
	Streams        *StreamService
	Timelines      *TimelineService
	Trends         *TrendsService
	Users          *UserService
}

type StatusUpdateParam struct {
	Text                    string   	`url:"text,omitempty"`
	InReplyToStatusID         int64    `url:"in_reply_to_status_id,omitempty"`
	AutoPopulateReplyMetadata *bool    `url:"auto_populate_reply_metadata,omitempty"`
	ExcludeReplyUserIds       []int64  `url:"exclude_reply_user_ids,comma,omitempty"`
	AttachmentURL             string   `url:"attachment_url,omitempty"`
	MediaIds                  []int64  `url:"media_ids,omitempty,comma"`
	PossiblySensitive         *bool    `url:"possibly_sensitive,omitempty"`
	Lat                       *float64 `url:"lat,omitempty"`
	Long                      *float64 `url:"long,omitempty"`
	PlaceID                   string   `url:"place_id,omitempty"`
	DisplayCoordinates        *bool    `url:"display_coordinates,omitempty"`
	TrimUser                  *bool    `url:"trim_user,omitempty"`
	CardURI                   string   `url:"card_uri,omitempty"`
	// Deprecated
	TweetMode string `url:"tweet_mode,omitempty"`
}


// NewClient returns a new Client.
func NewClient(httpClient *http.Client) *Client {
	base := sling.New().Client(httpClient).Base(twitterAPI)

	return &Client{
		sling:          base,
		Accounts:       newAccountService(base.New()),
		Blocks:         newBlockService(base.New()),
		DirectMessages: newDirectMessageService(base.New()),
		Favorites:      newFavoriteService(base.New()),
		Followers:      newFollowerService(base.New()),
		Friends:        newFriendService(base.New()),
		Friendships:    newFriendshipService(base.New()),
		Lists:          newListService(base.New()),
		RateLimits:     newRateLimitService(base.New()),
		Search:         newSearchService(base.New()),
		PremiumSearch:  newPremiumSearchService(base.New()),
		Statuses:       newStatusService(base.New()),
		Streams:        newStreamService(httpClient, base.New()),
		Timelines:      newTimelineService(base.New()),
		Trends:         newTrendsService(base.New()),
		Users:          newUserService(base.New()),
	}
}

// Bool returns a new pointer to the given bool value.
func Bool(v bool) *bool {
	ptr := new(bool)
	*ptr = v
	return ptr
}

// Float returns a new pointer to the given float64 value.
func Float(v float64) *float64 {
	ptr := new(float64)
	*ptr = v
	return ptr
}
