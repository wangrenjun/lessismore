// Code generated by "stringer -type ReturnCode -linecomment -output reasons.go"; DO NOT EDIT.

package codes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RC_OK-0]
	_ = x[RC_INVALID_ARGS-1]
	_ = x[RC_MALFORMED_MESSAGE-2]
	_ = x[RC_MSG_TOOL_LARGE-3]
	_ = x[RC_MESSAGE_UNDEFINED-4]
	_ = x[RC_SERVER_ERROR-5]
	_ = x[RC_SERVICE_DOWN-6]
	_ = x[RC_RESOURCE_NOT_FOUND-7]
	_ = x[RC_REQUEST_TIMEOUT-8]
	_ = x[RC_TOKEN_EXPIRED-9]
	_ = x[RC_TOKEN_REQUIRED-10]
	_ = x[RC_UID_REQUIRED-11]
	_ = x[RC_TOKEN_MISMATCH-12]
	_ = x[RC_TOKEN_NOT_EXIST-13]
	_ = x[RC_USER_NOT_FOUND-14]
	_ = x[RC_ROOM_NOT_FOUND-15]
	_ = x[RC_TABLE_NOT_FOUND-16]
	_ = x[RC_MAXIMUM-17]
}

const _ReturnCode_name = "OKInvalid argumentsMalformed messageMessage too largeMessage undefinedServer errorService is downResource not foundRequest timeoutToken expiredToken requiredUID requiredToken mismatchedToken not existUser does not foundRoom does not foundTable does not foundMaximum"

var _ReturnCode_index = [...]uint16{0, 2, 19, 36, 53, 70, 82, 97, 115, 130, 143, 157, 169, 185, 200, 219, 238, 258, 265}

func (i ReturnCode) String() string {
	if i < 0 || i >= ReturnCode(len(_ReturnCode_index)-1) {
		return "ReturnCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ReturnCode_name[_ReturnCode_index[i]:_ReturnCode_index[i+1]]
}
