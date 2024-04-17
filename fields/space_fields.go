package fields

type SpaceField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/space
const (
	SpaceFieldID               SpaceField = "id"
	SpaceFieldState            SpaceField = "state"
	SpaceFieldCreatedAt        SpaceField = "created_at"
	SpaceFieldEndedAt          SpaceField = "ended_at"
	SpaceFieldHostIDs          SpaceField = "host_ids"
	SpaceFieldLang             SpaceField = "lang"
	SpaceFieldIsTicketed       SpaceField = "is_ticketed"
	SpaceFieldInvitedUserIDs   SpaceField = "invited_user_ids"
	SpaceFieldParticipantCount SpaceField = "participant_count"
	SpaceFieldSubscriberCount  SpaceField = "subscriber_count"
	SpaceFieldScheduledStart   SpaceField = "scheduled_start"
	SpaceFieldSpeakerIDs       SpaceField = "speaker_ids"
	SpaceFieldStartedAt        SpaceField = "started_at"
	SpaceFieldTitle            SpaceField = "title"
	SpaceFieldTopicIDs         SpaceField = "topic_ids"
	SpaceFieldUpdatedAt        SpaceField = "updated_at"
)

func (f SpaceField) String() string {
	return string(f)
}

type SpaceFieldList []SpaceField

func (fl SpaceFieldList) FieldsName() string {
	return "space.fields"
}

func (fl SpaceFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
