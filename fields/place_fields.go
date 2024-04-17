package fields

type PlaceField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/place
const (
	PlaceFieldFullName        PlaceField = "full_name"
	PlaceFieldID              PlaceField = "id"
	PlaceFieldContainedWithin PlaceField = "contained_within"
	PlaceFieldCountry         PlaceField = "country"
	PlaceFieldCountryCode     PlaceField = "country_code"
	PlaceFieldGeo             PlaceField = "geo"
	PlaceFieldName            PlaceField = "name"
	PlaceFieldPlaceType       PlaceField = "place_type"
)

func (f PlaceField) String() string {
	return string(f)
}

type PlaceFieldList []PlaceField

func (fl PlaceFieldList) FieldsName() string {
	return "place.fields"
}

func (fl PlaceFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
