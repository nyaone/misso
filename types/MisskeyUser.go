package types

import "time"

type MisskeyUser struct {
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	Username       string      `json:"username"`
	Host           interface{} `json:"host"`
	AvatarUrl      string      `json:"avatarUrl"`
	AvatarBlurhash string      `json:"avatarBlurhash"`
	IsBot          bool        `json:"isBot"`
	IsCat          bool        `json:"isCat"`
	OnlineStatus   string      `json:"onlineStatus"`
	Url            interface{} `json:"url"`
	Uri            interface{} `json:"uri"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	LastFetchedAt  interface{} `json:"lastFetchedAt"`
	BannerUrl      string      `json:"bannerUrl"`
	BannerBlurhash string      `json:"bannerBlurhash"`
	IsLocked       bool        `json:"isLocked"`
	IsSilenced     bool        `json:"isSilenced"`
	IsSuspended    bool        `json:"isSuspended"`
	Description    string      `json:"description"`
	Location       string      `json:"location"`
	Birthday       string      `json:"birthday"`
	Lang           string      `json:"lang"`
	Fields         []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"fields"`
	FollowersCount int      `json:"followersCount"`
	FollowingCount int      `json:"followingCount"`
	NotesCount     int      `json:"notesCount"`
	PinnedNoteIds  []string `json:"pinnedNoteIds"`
	PinnedNotes    []struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UserId    string    `json:"userId"`
		User      struct {
			Id             string      `json:"id"`
			Name           string      `json:"name"`
			Username       string      `json:"username"`
			Host           interface{} `json:"host"`
			AvatarUrl      string      `json:"avatarUrl"`
			AvatarBlurhash string      `json:"avatarBlurhash"`
			IsBot          bool        `json:"isBot"`
			IsCat          bool        `json:"isCat"`
			OnlineStatus   string      `json:"onlineStatus"`
		} `json:"user"`
		Text         string         `json:"text"`
		Cw           *string        `json:"cw"`
		Visibility   string         `json:"visibility"`
		LocalOnly    bool           `json:"localOnly"`
		RenoteCount  int            `json:"renoteCount"`
		RepliesCount int            `json:"repliesCount"`
		Reactions    map[string]int `json:"reactions"`
		FileIds      []string       `json:"fileIds"`
		Files        []struct {
			Id          string    `json:"id"`
			CreatedAt   time.Time `json:"createdAt"`
			Name        string    `json:"name"`
			Type        string    `json:"type"`
			Md5         string    `json:"md5"`
			Size        int       `json:"size"`
			IsSensitive bool      `json:"isSensitive"`
			Blurhash    string    `json:"blurhash"`
			Properties  struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"properties"`
			Url          string      `json:"url"`
			ThumbnailUrl string      `json:"thumbnailUrl"`
			Comment      interface{} `json:"comment"`
			FolderId     interface{} `json:"folderId"`
			Folder       interface{} `json:"folder"`
			UserId       interface{} `json:"userId"`
			User         interface{} `json:"user"`
		} `json:"files"`
		ReplyId    interface{} `json:"replyId"`
		RenoteId   interface{} `json:"renoteId"`
		MyReaction string      `json:"myReaction,omitempty"`
	} `json:"pinnedNotes"`
	PinnedPageId string `json:"pinnedPageId"`
	PinnedPage   struct {
		Id        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		UserId    string    `json:"userId"`
		User      struct {
			Id             string      `json:"id"`
			Name           string      `json:"name"`
			Username       string      `json:"username"`
			Host           interface{} `json:"host"`
			AvatarUrl      string      `json:"avatarUrl"`
			AvatarBlurhash string      `json:"avatarBlurhash"`
			IsBot          bool        `json:"isBot"`
			IsCat          bool        `json:"isCat"`
			OnlineStatus   string      `json:"onlineStatus"`
		} `json:"user"`
		Content []struct {
			Id    string `json:"id"`
			Text  string `json:"text,omitempty"`
			Type  string `json:"type"`
			Title string `json:"title,omitempty"`
		} `json:"content"`
		Variables           []interface{} `json:"variables"`
		Title               string        `json:"title"`
		Name                string        `json:"name"`
		Summary             string        `json:"summary"`
		HideTitleWhenPinned bool          `json:"hideTitleWhenPinned"`
		AlignCenter         bool          `json:"alignCenter"`
		Font                string        `json:"font"`
		Script              string        `json:"script"`
		EyeCatchingImageId  string        `json:"eyeCatchingImageId"`
		EyeCatchingImage    struct {
			Id          string    `json:"id"`
			CreatedAt   time.Time `json:"createdAt"`
			Name        string    `json:"name"`
			Type        string    `json:"type"`
			Md5         string    `json:"md5"`
			Size        int       `json:"size"`
			IsSensitive bool      `json:"isSensitive"`
			Blurhash    string    `json:"blurhash"`
			Properties  struct {
				Width  int `json:"width"`
				Height int `json:"height"`
			} `json:"properties"`
			Url          string      `json:"url"`
			ThumbnailUrl string      `json:"thumbnailUrl"`
			Comment      interface{} `json:"comment"`
			FolderId     interface{} `json:"folderId"`
			Folder       interface{} `json:"folder"`
			UserId       interface{} `json:"userId"`
			User         interface{} `json:"user"`
		} `json:"eyeCatchingImage"`
		AttachedFiles []interface{} `json:"attachedFiles"`
		LikedCount    int           `json:"likedCount"`
		IsLiked       bool          `json:"isLiked"`
	} `json:"pinnedPage"`
	PublicReactions      bool   `json:"publicReactions"`
	FfVisibility         string `json:"ffVisibility"`
	TwoFactorEnabled     bool   `json:"twoFactorEnabled"`
	UsePasswordLessLogin bool   `json:"usePasswordLessLogin"`
	SecurityKeys         bool   `json:"securityKeys"`
	Roles                []struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		Color           string `json:"color"`
		Description     string `json:"description"`
		IsModerator     bool   `json:"isModerator"`
		IsAdministrator bool   `json:"isAdministrator"`
	} `json:"roles"`
	AvatarId                        string `json:"avatarId"`
	BannerId                        string `json:"bannerId"`
	IsModerator                     bool   `json:"isModerator"`
	IsAdmin                         bool   `json:"isAdmin"`
	InjectFeaturedNote              bool   `json:"injectFeaturedNote"`
	ReceiveAnnouncementEmail        bool   `json:"receiveAnnouncementEmail"`
	AlwaysMarkNsfw                  bool   `json:"alwaysMarkNsfw"`
	AutoSensitive                   bool   `json:"autoSensitive"`
	CarefulBot                      bool   `json:"carefulBot"`
	AutoAcceptFollowed              bool   `json:"autoAcceptFollowed"`
	NoCrawle                        bool   `json:"noCrawle"`
	IsExplorable                    bool   `json:"isExplorable"`
	IsDeleted                       bool   `json:"isDeleted"`
	HideOnlineStatus                bool   `json:"hideOnlineStatus"`
	HasUnreadSpecifiedNotes         bool   `json:"hasUnreadSpecifiedNotes"`
	HasUnreadMentions               bool   `json:"hasUnreadMentions"`
	HasUnreadAnnouncement           bool   `json:"hasUnreadAnnouncement"`
	HasUnreadAntenna                bool   `json:"hasUnreadAntenna"`
	HasUnreadChannel                bool   `json:"hasUnreadChannel"`
	HasUnreadMessagingMessage       bool   `json:"hasUnreadMessagingMessage"`
	HasUnreadNotification           bool   `json:"hasUnreadNotification"`
	HasPendingReceivedFollowRequest bool   `json:"hasPendingReceivedFollowRequest"`
	Integrations                    struct {
	} `json:"integrations"`
	MutedWords              [][]string `json:"mutedWords"`
	MutedInstances          []string   `json:"mutedInstances"`
	MutingNotificationTypes []string   `json:"mutingNotificationTypes"`
	EmailNotificationTypes  []string   `json:"emailNotificationTypes"`
	ShowTimelineReplies     bool       `json:"showTimelineReplies"`
}
