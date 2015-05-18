package example

//
// Post -
//
// Multiple line
// documentations
// about this Post type
//
type Post struct {

	// Id documentation line
	Id int64 `db:"id" json:"id" foo:"bar"` // Id comment

	// UserId documentation line
	UserId int `db:"user_id" json:"user_id" foo:"-,omitempty"` // UserId comment

	// Title documentation line
	Title string `db:"title" json:"title" foo:"bar"` // Title comment

	// Body documentation line
	Body string `db:"body" json:"body" foo:"-,omitempty"` // Body comment

}

/**
 *
 * Comment -
 *
 * Multiple line
 * documentations
 * about this Comment type
 *
 */
type Comment struct {

	// Id documentation line
	Id int64 `db:"id" json:"id"` // Id comment

	// UserId documentation line
	UserId int `db:"user_id" json:"user_id"` // UserId comment

	// Title documentation line
	Title string `db:"title" json:"title"` // Title comment

	// Body documentation line
	Body string `db:"body" json:"body"` // Body comment

}

/**

  User -

  Multiple line
  documentations
  about this User type

*/
type User struct {

	// Id documentation line
	Id int `db:"id" json:"id"` // Id comment

	// Name documentation line
	Name string `db:"name" json:"name"` // Name comment

	// Email documentation line
	Email string `db:"email" json:"email"` // Email comment

	// Password documentation line
	Password string `db:"password" json:"-"` // Password comment

}
