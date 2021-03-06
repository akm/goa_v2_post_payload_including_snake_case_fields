// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// user views
//
// Command:
// $ goa gen
// github.com/akm/goa_v2_post_payload_including_snake_case_fields/design

package views

import (
	goa "goa.design/goa"
)

// User is the viewed result type that is projected based on a view.
type User struct {
	// Type to project
	Projected *UserView
	// View to render
	View string
}

// UserView is a type that runs validations on a projected type.
type UserView struct {
	ID        *string
	FirstName *string
	LastName  *string
}

// ValidateUser runs the validations defined on the viewed result type User.
func ValidateUser(result *User) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUserView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateUserView runs the validations defined on UserView using the
// "default" view.
func ValidateUserView(result *UserView) (err error) {
	if result.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first_name", "result"))
	}
	if result.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_name", "result"))
	}
	return
}
