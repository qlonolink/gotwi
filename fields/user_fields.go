package fields

type UserField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user
const (
	UserFieldCreatedAt       UserField = "created_at"
	UserFieldDescription     UserField = "description"
	UserFieldEntities        UserField = "entities"
	UserFieldID              UserField = "id"
	UserFieldLocation        UserField = "location"
	UserFieldName            UserField = "name"
	UserFieldPinnedTweetID   UserField = "pinned_tweet_id"
	UserFieldProfileImageUrl UserField = "profile_image_url"
	UserFieldProtected       UserField = "protected"
	UserFieldPublicMetrics   UserField = "public_metrics"
	UserFieldUrl             UserField = "url"
	UserFieldUsername        UserField = "username"
	UserFieldVerified        UserField = "verified"
	UserFieldVerifiedType    UserField = "verified_type"
	UserFieldWithheld        UserField = "withheld"

	UserFieldConnectionStatus UserField = "connection_status"
)

func (f UserField) String() string {
	return string(f)
}

type UserFieldList []UserField

func (fl UserFieldList) FieldsName() string {
	return "user.fields"
}

func (fl UserFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
