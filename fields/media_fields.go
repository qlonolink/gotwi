package fields

type MediaField string

// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/media
const (
	MediaFieldMediaKey         MediaField = "media_key"
	MediaFieldType             MediaField = "type"
	MediaFieldUrl              MediaField = "url"
	MediaFieldDurationMs       MediaField = "duration_ms"
	MediaFieldHeight           MediaField = "height"
	MediaFieldNonPublicMetrics MediaField = "non_public_metrics"
	MediaFieldOrganicMetrics   MediaField = "organic_metrics"
	MediaFieldPreviewImageUrl  MediaField = "preview_image_url"
	MediaFieldPromotedMetrics  MediaField = "promoted_metrics"
	MediaFieldPublicMetrics    MediaField = "public_metrics"
	MediaFieldWidth            MediaField = "width"
	MediaFieldAltText          MediaField = "alt_text"
	MediaFieldVariants         MediaField = "variants"
)

func (f MediaField) String() string {
	return string(f)
}

type MediaFieldList []MediaField

func (fl MediaFieldList) FieldsName() string {
	return "media.fields"
}

func (fl MediaFieldList) Values() []string {
	if fl == nil {
		return []string{}
	}

	s := []string{}
	for _, f := range fl {
		s = append(s, f.String())
	}

	return s
}
