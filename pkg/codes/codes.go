//go:generate stringer -type ReturnCode -linecomment -output reasons.go
package codes

import "encoding/json"

type ReturnCode int

func (rc ReturnCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code   int
		Reason string
	}{
		Code:   int(rc),
		Reason: rc.String(),
	})
}

const (
	RC_OK                 ReturnCode = iota // OK
	RC_INVALID_ARGS                         // Invalid arguments
	RC_MALFORMED_MESSAGE                    // Malformed message
	RC_MSG_TOOL_LARGE                       // Message too large
	RC_MESSAGE_UNDEFINED                    // Message undefined
	RC_SERVER_ERROR                         // Server error
	RC_SERVICE_DOWN                         // Service is down
	RC_RESOURCE_NOT_FOUND                   // Resource not found
	RC_REQUEST_TIMEOUT                      // Request timeout

	RC_TOKEN_EXPIRED   // Token expired
	RC_TOKEN_REQUIRED  // Token required
	RC_UID_REQUIRED    // UID required
	RC_TOKEN_MISMATCH  // Token mismatched
	RC_TOKEN_NOT_EXIST // Token not exist
	RC_USER_NOT_FOUND  // User does not found
	RC_ROOM_NOT_FOUND  // Room does not found
	RC_TABLE_NOT_FOUND // Table does not found

	/////////////////////////////////////////////
	RC_MAXIMUM // Maximum
)
