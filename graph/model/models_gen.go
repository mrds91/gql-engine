// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Customer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
	ContactNo string `json:"contact_no"`
}

type CustomerReq struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Gender    string `json:"gender"`
	ContactNo string `json:"contact_no"`
}
