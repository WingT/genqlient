// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package test

import (
	"github.com/Khan/genqlient/graphql"
)

// EmptyInterfaceResponse is returned by EmptyInterface on success.
type EmptyInterfaceResponse struct {
	GetJunk        interface{}                             `json:"getJunk"`
	GetComplexJunk []map[string]*[]*map[string]interface{} `json:"getComplexJunk"`
}

// GetGetJunk returns EmptyInterfaceResponse.GetJunk, and is useful for accessing the field via an interface.
func (v *EmptyInterfaceResponse) GetGetJunk() (val interface{}) {
	if v == nil {
		return
	}
	return v.GetJunk
}

// GetGetComplexJunk returns EmptyInterfaceResponse.GetComplexJunk, and is useful for accessing the field via an interface.
func (v *EmptyInterfaceResponse) GetGetComplexJunk() (val []map[string]*[]*map[string]interface{}) {
	if v == nil {
		return
	}
	return v.GetComplexJunk
}

// The query executed by EmptyInterface.
const EmptyInterface_Operation = `
query EmptyInterface {
	getJunk
	getComplexJunk
}
`

func EmptyInterface(
	client_ graphql.Client,
) (data_ *EmptyInterfaceResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "EmptyInterface",
		Query:  EmptyInterface_Operation,
	}

	data_ = &EmptyInterfaceResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		nil,
		req_,
		resp_,
	)

	return data_, err_
}

