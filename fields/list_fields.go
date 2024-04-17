package fields

type ListField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/lists
const (
	ListFieldCreatedAt     ListField = "created_at"
	ListFieldFollowerCount ListField = "follower_count"
	ListFieldMemberCount   ListField = "member_count"
	ListFieldPrivate       ListField = "private"
	ListFieldDescription   ListField = "description"
	ListFieldOwnerID       ListField = "owner_id"

	ListFieldID   ListField = "id"
	ListFieldName ListField = "name"
)

func (f ListField) String() string {
	return string(f)
}

type ListFieldList []ListField

func (fl ListFieldList) FieldsName() string {
	return "list.fields"
}

func (fl ListFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
