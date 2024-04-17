package fields

type UserField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/user
const (
	UserFieldID               UserField = "id"
	UserFieldName             UserField = "name"
	UserFieldUsername         UserField = "username"
	UserFieldConnectionStatus UserField = "connection_status"
	UserFieldCreatedAt        UserField = "created_at"
	UserFieldDescription      UserField = "description"
	UserFieldEntities         UserField = "entities"
	UserFieldLocation         UserField = "location"
	UserFieldPinnedTweetID    UserField = "pinned_tweet_id"
	UserFieldProfileImageUrl  UserField = "profile_image_url"
	UserFieldProtected        UserField = "protected"
	UserFieldPublicMetrics    UserField = "public_metrics"
	UserFieldUrl              UserField = "url"
	UserFieldVerified         UserField = "verified"
	UserFieldVerifiedType     UserField = "verified_type"
	UserFieldWithheld         UserField = "withheld"
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
