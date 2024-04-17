package fields

type Expansion string

// https://developer.twitter.com/en/docs/twitter-api/expansions
const (
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionEditHistoryTweetIDs        Expansion = "edit_history_tweet_ids"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionEntitiesMentionsUsername   Expansion = "entities.mentions.username"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
	ExpansionPinnedTweetID              Expansion = "pinned_tweet_id"
	ExpansionSenderID                   Expansion = "sender_id"
	ExpansionParticipantIDs             Expansion = "participant_ids"
	ExpansionInvitedUserIDs             Expansion = "invited_user_ids"
	ExpansionSpeakerIDs                 Expansion = "speaker_ids"
	ExpansionCreatorID                  Expansion = "creator_id"
	ExpansionHostIDs                    Expansion = "host_ids"
	ExpansionTopicsIDs                  Expansion = "topics_ids"
	ExpansionOwnerID                    Expansion = "owner_id"
)

func (e Expansion) String() string {
	return string(e)
}

type ExpansionList []Expansion

func (el ExpansionList) FieldsName() string {
	return "expansions"
}

func (el ExpansionList) Values() []string {
	if el == nil {
		return []string{}
	}

	s := []string{}
	for _, e := range el {
		s = append(s, e.String())
	}

	return s
}
