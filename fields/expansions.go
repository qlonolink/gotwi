package fields

type Expansion string

// https://developer.twitter.com/en/docs/twitter-api/expansions
const (
	ExpansionPinnedTweetID              Expansion = "pinned_tweet_id"
	ExpansionAttachmentsPollIDs         Expansion = "attachments.poll_ids"
	ExpansionAttachmentsMediaKeys       Expansion = "attachments.media_keys"
	ExpansionAuthorID                   Expansion = "author_id"
	ExpansionEntitiesMentionsUsername   Expansion = "entities.mentions.username"
	ExpansionGeoPlaceID                 Expansion = "geo.place_id"
	ExpansionInReplyToUserID            Expansion = "in_reply_to_user_id"
	ExpansionReferencedTweetsID         Expansion = "referenced_tweets.id"
	ExpansionReferencedTweetsIDAuthorID Expansion = "referenced_tweets.id.author_id"
	ExpansionInvitedUserIDs             Expansion = "invited_user_ids"
	ExpansionSpeakerIDs                 Expansion = "speaker_ids"
	ExpansionCreatorID                  Expansion = "creator_id"
	ExpansionHostIDs                    Expansion = "host_ids"

	ExpansionEditHistoryTweetIDs Expansion = "edit_history_tweet_ids"
	ExpansionSenderID            Expansion = "sender_id"
	ExpansionParticipantIDs      Expansion = "participant_ids"
	ExpansionTopicsIDs           Expansion = "topics_ids"
	ExpansionOwnerID             Expansion = "owner_id"
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
