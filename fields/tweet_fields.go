package fields

type TweetField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/tweet
const (
	TweetFieldID                  TweetField = "id"
	TweetFieldText                TweetField = "text"
	TweetFieldEditHistoryTweetIDs TweetField = "edit_history_tweet_ids"
	TweetFieldAttachments         TweetField = "attachments"
	TweetFieldAuthorID            TweetField = "author_id"
	TweetFieldContextAnnotations  TweetField = "context_annotations"
	TweetFieldConversationID      TweetField = "conversation_id"
	TweetFieldCreatedAt           TweetField = "created_at"
	TweetFieldEditControls        TweetField = "edit_controls"
	TweetFieldEntities            TweetField = "entities"
	TweetFieldInReplyToUserID     TweetField = "in_reply_to_user_id"
	TweetFieldLang                TweetField = "lang"
	TweetFieldNonPublicMetrics    TweetField = "non_public_metrics"
	TweetFieldOrganicMetrics      TweetField = "organic_metrics"
	TweetFieldPossiblySensitive   TweetField = "possibly_sensitive"
	TweetFieldPromotedMetrics     TweetField = "promoted_metrics"
	TweetFieldPublicMetrics       TweetField = "public_metrics"
	TweetFieldReferencedTweets    TweetField = "referenced_tweets"
	TweetFieldReplySettings       TweetField = "reply_settings"
	TweetFieldWithheld            TweetField = "withheld"
	TweetFieldGeo                 TweetField = "geo"
	TweetFieldSource              TweetField = "source"
)

func (f TweetField) String() string {
	return string(f)
}

type TweetFieldList []TweetField

func (fl TweetFieldList) FieldsName() string {
	return "tweet.fields"
}

func (fl TweetFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}

func (fl TweetFieldList) HasContextAnnotations() bool {
	for _, v := range fl {
		if v == TweetFieldContextAnnotations {
			return true
		}
	}
	return false
}
