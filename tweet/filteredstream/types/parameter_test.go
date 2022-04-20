package types_test

import (
	"testing"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/tweet/filteredstream/types"
	"github.com/stretchr/testify/assert"
)

func Test_FilteredStreamRulesGet_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.ListRulesInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_FilteredStreamRulesGet_ResolveEndpoint(t *testing.T) {
	const endpointBase = "test/endpoint"

	cases := []struct {
		name   string
		params *types.ListRulesInput
		expect string
	}{
		{
			name:   "has no parameter",
			params: &types.ListRulesInput{},
			expect: endpointBase,
		},
		{
			name: "with ids",
			params: &types.ListRulesInput{
				IDs: []string{"rid1", "rid2"},
			},
			expect: endpointBase + "?ids=rid1%2Crid2",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpointBase)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_FilteredStreamRulesGet_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.ListRulesInput
	}{
		{
			name:   "empty params",
			params: &types.ListRulesInput{},
		},
		{
			name:   "some params",
			params: &types.ListRulesInput{IDs: []string{"rid1", "rid2"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}

func Test_SearchStreamInput_SetAccessToken(t *testing.T) {
	cases := []struct {
		name   string
		token  string
		expect string
	}{
		{
			name:   "normal",
			token:  "test-token",
			expect: "test-token",
		},
		{
			name:   "empty",
			token:  "",
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			p := &types.SearchStreamInput{}
			p.SetAccessToken(c.token)
			assert.Equal(tt, c.expect, p.AccessToken())
		})
	}
}

func Test_SearchStreamInput_ResolveEndpoint(t *testing.T) {
	const endpoint = "test/endpoint/"
	cases := []struct {
		name   string
		params *types.SearchStreamInput
		expect string
	}{
		{
			name:   "ok",
			params: &types.SearchStreamInput{},
			expect: endpoint,
		},
		{
			name: "with valid backfill_minutes",
			params: &types.SearchStreamInput{
				BackfillMinutes: 3,
			},
			expect: endpoint + "?backfill_minutes=3",
		},
		{
			name: "with invalid backfill_minutes",
			params: &types.SearchStreamInput{
				BackfillMinutes: -1,
			},
			expect: endpoint,
		},
		{
			name: "with expansions",
			params: &types.SearchStreamInput{
				Expansions: fields.ExpansionList{"ex1", "ex2"},
			},
			expect: endpoint + "?expansions=ex1%2Cex2",
		},
		{
			name: "with media.fields",
			params: &types.SearchStreamInput{
				MediaFields: fields.MediaFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?media.fields=tf1%2Ctf2",
		},
		{
			name: "with place.fields",
			params: &types.SearchStreamInput{
				PlaceFields: fields.PlaceFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?place.fields=tf1%2Ctf2",
		},
		{
			name: "with poll.fields",
			params: &types.SearchStreamInput{
				PollFields: fields.PollFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?poll.fields=tf1%2Ctf2",
		},
		{
			name: "with tweets.fields",
			params: &types.SearchStreamInput{
				TweetFields: fields.TweetFieldList{"tf1", "tf2"},
			},
			expect: endpoint + "?tweet.fields=tf1%2Ctf2",
		},
		{
			name: "with users.fields",
			params: &types.SearchStreamInput{
				UserFields: fields.UserFieldList{"uf1", "uf2"},
			},
			expect: endpoint + "?user.fields=uf1%2Cuf2",
		},
		{
			name: "all query parameters",
			params: &types.SearchStreamInput{
				BackfillMinutes: 3,
				Expansions:      fields.ExpansionList{"ex"},
				MediaFields:     fields.MediaFieldList{"mf"},
				PlaceFields:     fields.PlaceFieldList{"plf"},
				PollFields:      fields.PollFieldList{"pof"},
				UserFields:      fields.UserFieldList{"uf"},
				TweetFields:     fields.TweetFieldList{"tf"},
			},
			expect: endpoint + "?backfill_minutes=3&expansions=ex&media.fields=mf&place.fields=plf&poll.fields=pof&tweet.fields=tf&user.fields=uf",
		},
		{
			name:   "nil",
			params: nil,
			expect: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ep := c.params.ResolveEndpoint(endpoint)
			assert.Equal(tt, c.expect, ep)
		})
	}
}

func Test_SearchStreamInput_Body(t *testing.T) {
	cases := []struct {
		name   string
		params *types.SearchStreamInput
	}{
		{
			name:   "empty params",
			params: &types.SearchStreamInput{},
		},
		{
			name:   "some params",
			params: &types.SearchStreamInput{Expansions: fields.ExpansionList{"ex"}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			r, err := c.params.Body()
			assert.NoError(tt, err)
			assert.Nil(tt, r)
		})
	}
}
