package types

import "github.com/qlonolink/gotwi/resources"

type ListOutput struct {
	Data     []resources.Space `json:"data"`
	Includes struct {
		Users []resources.User `json:"users"`
	} `json:"includes"`
	Errors []resources.PartialError `json:"errors"`
}

func (r *ListOutput) HasPartialError() bool {
	return !(r.Errors == nil || len(r.Errors) == 0)
}
