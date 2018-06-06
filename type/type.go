//cat tweet.json |gojson

package main

type Foo []struct {
	Contributors interface{} `json:"contributors"`
	Coordinates  interface{} `json:"coordinates"`
	CreatedAt    string      `json:"created_at"`
	Entities     struct {
		Hashtags []interface{} `json:"hashtags"`
		Media    []struct {
			DisplayURL    string  `json:"display_url"`
			ExpandedURL   string  `json:"expanded_url"`
			ID            int64   `json:"id"`
			IDStr         string  `json:"id_str"`
			Indices       []int64 `json:"indices"`
			MediaURL      string  `json:"media_url"`
			MediaURLHTTPS string  `json:"media_url_https"`
			Sizes         struct {
				Large struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"large"`
				Medium struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"medium"`
				Small struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"small"`
				Thumb struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"thumb"`
			} `json:"sizes"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"media"`
		Symbols []interface{} `json:"symbols"`
		Urls    []struct {
			DisplayURL  string  `json:"display_url"`
			ExpandedURL string  `json:"expanded_url"`
			Indices     []int64 `json:"indices"`
			URL         string  `json:"url"`
		} `json:"urls"`
		UserMentions []interface{} `json:"user_mentions"`
	} `json:"entities"`
	ExtendedEntities struct {
		Media []struct {
			DisplayURL    string  `json:"display_url"`
			ExpandedURL   string  `json:"expanded_url"`
			ID            int64   `json:"id"`
			IDStr         string  `json:"id_str"`
			Indices       []int64 `json:"indices"`
			MediaURL      string  `json:"media_url"`
			MediaURLHTTPS string  `json:"media_url_https"`
			Sizes         struct {
				Large struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"large"`
				Medium struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"medium"`
				Small struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"small"`
				Thumb struct {
					H      int64  `json:"h"`
					Resize string `json:"resize"`
					W      int64  `json:"w"`
				} `json:"thumb"`
			} `json:"sizes"`
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"media"`
	} `json:"extended_entities"`
	FavoriteCount               int64       `json:"favorite_count"`
	Favorited                   bool        `json:"favorited"`
	Geo                         interface{} `json:"geo"`
	ID                          int64       `json:"id"`
	IDStr                       string      `json:"id_str"`
	InReplyToScreenName         interface{} `json:"in_reply_to_screen_name"`
	InReplyToStatusID           interface{} `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr        interface{} `json:"in_reply_to_status_id_str"`
	InReplyToUserID             interface{} `json:"in_reply_to_user_id"`
	InReplyToUserIDStr          interface{} `json:"in_reply_to_user_id_str"`
	IsQuoteStatus               bool        `json:"is_quote_status"`
	Lang                        string      `json:"lang"`
	Place                       interface{} `json:"place"`
	PossiblySensitive           bool        `json:"possibly_sensitive"`
	PossiblySensitiveAppealable bool        `json:"possibly_sensitive_appealable"`
	RetweetCount                int64       `json:"retweet_count"`
	Retweeted                   bool        `json:"retweeted"`
	Source                      string      `json:"source"`
	Text                        string      `json:"text"`
	Truncated                   bool        `json:"truncated"`
	User                        struct {
		ContributorsEnabled bool   `json:"contributors_enabled"`
		CreatedAt           string `json:"created_at"`
		DefaultProfile      bool   `json:"default_profile"`
		DefaultProfileImage bool   `json:"default_profile_image"`
		Description         string `json:"description"`
		Entities            struct {
			Description struct {
				Urls []interface{} `json:"urls"`
			} `json:"description"`
			URL struct {
				Urls []struct {
					DisplayURL  string  `json:"display_url"`
					ExpandedURL string  `json:"expanded_url"`
					Indices     []int64 `json:"indices"`
					URL         string  `json:"url"`
				} `json:"urls"`
			} `json:"url"`
		} `json:"entities"`
		FavouritesCount                int64       `json:"favourites_count"`
		FollowRequestSent              bool        `json:"follow_request_sent"`
		FollowersCount                 int64       `json:"followers_count"`
		Following                      bool        `json:"following"`
		FriendsCount                   int64       `json:"friends_count"`
		GeoEnabled                     bool        `json:"geo_enabled"`
		HasExtendedProfile             bool        `json:"has_extended_profile"`
		ID                             int64       `json:"id"`
		IDStr                          string      `json:"id_str"`
		IsTranslationEnabled           bool        `json:"is_translation_enabled"`
		IsTranslator                   bool        `json:"is_translator"`
		Lang                           string      `json:"lang"`
		ListedCount                    int64       `json:"listed_count"`
		Location                       string      `json:"location"`
		Name                           string      `json:"name"`
		Notifications                  bool        `json:"notifications"`
		ProfileBackgroundColor         string      `json:"profile_background_color"`
		ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
		ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
		ProfileBackgroundTile          bool        `json:"profile_background_tile"`
		ProfileBannerURL               string      `json:"profile_banner_url"`
		ProfileImageURL                string      `json:"profile_image_url"`
		ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
		ProfileLinkColor               string      `json:"profile_link_color"`
		ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
		ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
		ProfileTextColor               string      `json:"profile_text_color"`
		ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
		Protected                      bool        `json:"protected"`
		ScreenName                     string      `json:"screen_name"`
		StatusesCount                  int64       `json:"statuses_count"`
		TimeZone                       interface{} `json:"time_zone"`
		TranslatorType                 string      `json:"translator_type"`
		URL                            string      `json:"url"`
		UtcOffset                      interface{} `json:"utc_offset"`
		Verified                       bool        `json:"verified"`
	} `json:"user"`
}
