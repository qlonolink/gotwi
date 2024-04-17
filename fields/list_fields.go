package fields

type ListField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/lists
const (
	ListFieldID            ListField = "id"
	ListFieldName          ListField = "name"
	ListFieldCreatedAt     ListField = "created_at"
	ListFieldDescription   ListField = "description"
	ListFieldFollowerCount ListField = "follower_count"
	ListFieldMemberCount   ListField = "member_count"
	ListFieldPrivate       ListField = "private"
	ListFieldOwnerID       ListField = "owner_id"
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
