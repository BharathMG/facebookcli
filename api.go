package main

import (
	fb "github.com/huandu/facebook"
)

// GET scoped information with valid session
func FbGet(path string, data interface{}) error {
	res, _ := GetSession().Get(path, nil)
	return res.Decode(data)
}

// GET Public information.
func FbPublicGet(path string, data interface{}) error {
	res, _ := fb.Get(path, fb.Params{"field": "bio"})
	return res.Decode(data)
}

func FbPagingGet(path string) ([]fb.Result, error) {
	res, _ := GetSession().Get(path, nil)
	// create a paging structure.
	paging, err := res.Paging(session)

	// get current results.
	results := paging.Data()
	return results, err
}
