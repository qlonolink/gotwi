package fields

type PollField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/poll
const (
	PollFieldID              PollField = "id"
	PollFieldOptions         PollField = "options"
	PollFieldDurationMinutes PollField = "duration_minutes"
	PollFieldEndDatetime     PollField = "end_datetime"
	PollFieldVotingStatus    PollField = "voting_status"
)

func (f PollField) String() string {
	return string(f)
}

type PollFieldList []PollField

func (fl PollFieldList) FieldsName() string {
	return "poll.fields"
}

func (fl PollFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
