package constants

import "errors"

// noinspection GoSnakeCaseUsage
const ENABLE_CUSTOM_ERRORS = true

// noinspection GoSnakeCaseUsage
const (
	// SYSTEM_ERRORS
	ERROR_MESSAGE_INITIALIZE                    = "ERROR_MESSAGE_INITIALIZE"
	ERROR_MESSAGE_NOT_IMPLEMENTED               = "ERROR_MESSAGE_NOT_IMPLEMENTED"
	ERROR_MESSAGE_OPERATION_FAILED              = "ERROR_MESSAGE_OPERATION_FAILED"
	ERROR_MESSAGE_OPERATION_NOT_SUPPORTED       = "ERROR_MESSAGE_OPERATION_NOT_SUPPORTED"
	ERROR_MESSAGE_UNRESOLVED_DEPENDENCIES       = "ERROR_MESSAGE_UNRESOLVED_DEPENDENCIES"
	ERROR_MESSAGE_SYSTEM_COMPONENT_NOT_FOUND    = "ERROR_MESSAGE_SYSTEM_COMPONENT_NOT_FOUND"
	ERROR_MESSAGE_DOCUMENT_NOT_FOUND            = "ERROR_MESSAGE_DOCUMENT_NOT_FOUND"
	ERROR_MESSAGE_SYSTEM_SCHEDULE_NOT_FOUND     = "ERROR_MESSAGE_SYSTEM_SCHEDULE_NOT_FOUND"
	ERROR_MESSAGE_IDENTITY_NOT_FOUND            = "ERROR_MESSAGE_IDENTITY_NOT_FOUND"
	ERROR_MESSAGE_ACCESS_CONTROL_NOT_FOUND      = "ERROR_MESSAGE_ACCESS_CONTROL_NOT_FOUND"
	ERROR_MESSAGE_REMOTE_ACTIVITY_NOT_FOUND     = "ERROR_MESSAGE_REMOTE_ACTIVITY_NOT_FOUND"
	ERROR_MESSAGE_CATEGORY_TYPE_NOT_FOUND       = "ERROR_MESSAGE_CATEGORY_TYPE_NOT_FOUND"
	ERROR_MESSAGE_CATEGORY_NOT_FOUND            = "ERROR_MESSAGE_CATEGORY_NOT_FOUND"
	ERROR_MESSAGE_USER_NOT_FOUND                = "ERROR_MESSAGE_USER_NOT_FOUND"
	ERROR_MESSAGE_ACTIVITY_PUB_OBJECT_NOT_FOUND = "ERROR_MESSAGE_ACTIVITY_PUB_OBJECT_NOT_FOUND"
	ERROR_MESSAGE_SPI_NOT_FOUND                 = "ERROR_MESSAGE_SPI_NOT_FOUND"
	ERROR_MESSAGE_UNKNOWN_DOCUMENT              = "ERROR_MESSAGE_UNKNOWN_DOCUMENT"
	ERROR_MESSAGE_UNKNOWN_SYSTEM_SCHEDULE       = "ERROR_MESSAGE_UNKNOWN_SYSTEM_SCHEDULE"
	ERROR_MESSAGE_UNKNOWN_IDENTITY              = "ERROR_MESSAGE_UNKNOWN_IDENTITY"
	ERROR_MESSAGE_UNKNOWN_ACCESS_CONTROL        = "ERROR_MESSAGE_UNKNOWN_ACCESS_CONTROL"
	ERROR_MESSAGE_UNKNOWN_REMOTE_ACTIVITY       = "ERROR_MESSAGE_UNKNOWN_REMOTE_ACTIVITY"
	ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE         = "ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE"
	ERROR_MESSAGE_UNKNOWN_CATEGORY              = "ERROR_MESSAGE_UNKNOWN_CATEGORY"
	ERROR_MESSAGE_UNKNOWN_USER                  = "ERROR_MESSAGE_UNKNOWN_USER"
	ERROR_MESSAGE_UNKNOWN_ACTIVITY_PUB_OBJECT   = "ERROR_MESSAGE_UNKNOWN_ACTIVITY_PUB_OBJECT"
	ERROR_MESSAGE_UNKNOWN_SPI                   = "ERROR_MESSAGE_UNKNOWN_SPI"
	ERROR_MESSAGE_INVALID_ID                    = "ERROR_MESSAGE_INVALID_ID"
	ERROR_MESSAGE_INVALID_PARAMETERS            = "ERROR_MESSAGE_INVALID_PARAMETERS"
	// CUSTOM_ERRORS
	ERROR_MESSAGE_DATA_INTEGRITY_VIOLATION             = "ERROR_MESSAGE_DATA_INTEGRITY_VIOLATION"
	ERROR_MESSAGE_INVALID_STATE                        = "ERROR_MESSAGE_INVALID_STATE"
	ERROR_MESSAGE_USER_NOT_REGISTERED                  = "ERROR_MESSAGE_USER_NOT_REGISTERED"
	ERROR_MESSAGE_USERNAME_OR_EMAIL_ALREADY_REGISTERED = "ERROR_MESSAGE_USERNAME_OR_EMAIL_ALREADY_REGISTERED"
	ERROR_MESSAGE_ACCOUNT_NOT_VERIFIED                 = "ERROR_MESSAGE_ACCOUNT_NOT_VERIFIED"
	ERROR_MESSAGE_ACCOUNT_BLOCKED                      = "ERROR_MESSAGE_ACCOUNT_BLOCKED"
	ERROR_MESSAGE_INVALID_TOKEN                        = "ERROR_MESSAGE_INVALID_TOKEN"
	ERROR_MESSAGE_INVALID_CONFIRMATION_CODE            = "ERROR_MESSAGE_INVALID_CONFIRMATION_CODE"
	ERROR_MESSAGE_PERMISSION_DENIED                    = "ERROR_MESSAGE_PERMISSION_DENIED"
	ERROR_MESSAGE_INVALID_PERSON_KIND                  = "ERROR_MESSAGE_INVALID_PERSON_KIND"
	ERROR_MESSAGE_INVALID_CREDENTIALS                  = "ERROR_MESSAGE_INVALID_CREDENTIALS"
)

// noinspection GoSnakeCaseUsage,GoUnusedGlobalVariable
var (
	// SYSTEM_ERRORS
	ERROR_INITIALIZE                    = errors.New(ERROR_MESSAGE_INITIALIZE)
	ERROR_NOT_IMPLEMENTED               = errors.New(ERROR_MESSAGE_NOT_IMPLEMENTED)
	ERROR_OPERATION_FAILED              = errors.New(ERROR_MESSAGE_OPERATION_FAILED)
	ERROR_OPERATION_NOT_SUPPORTED       = errors.New(ERROR_MESSAGE_OPERATION_NOT_SUPPORTED)
	ERROR_UNRESOLVED_DEPENDENCIES       = errors.New(ERROR_MESSAGE_UNRESOLVED_DEPENDENCIES)
	ERROR_SYSTEM_COMPONENT_NOT_FOUND    = errors.New(ERROR_MESSAGE_SYSTEM_COMPONENT_NOT_FOUND)
	ERROR_DOCUMENT_NOT_FOUND            = errors.New(ERROR_MESSAGE_DOCUMENT_NOT_FOUND)
	ERROR_SYSTEM_SCHEDULE_NOT_FOUND     = errors.New(ERROR_MESSAGE_SYSTEM_SCHEDULE_NOT_FOUND)
	ERROR_IDENTITY_NOT_FOUND            = errors.New(ERROR_MESSAGE_IDENTITY_NOT_FOUND)
	ERROR_ACCESS_CONTROL_NOT_FOUND      = errors.New(ERROR_MESSAGE_ACCESS_CONTROL_NOT_FOUND)
	ERROR_REMOTE_ACTIVITY_NOT_FOUND     = errors.New(ERROR_MESSAGE_REMOTE_ACTIVITY_NOT_FOUND)
	ERROR_CATEGORY_TYPE_NOT_FOUND       = errors.New(ERROR_MESSAGE_CATEGORY_TYPE_NOT_FOUND)
	ERROR_CATEGORY_NOT_FOUND            = errors.New(ERROR_MESSAGE_CATEGORY_NOT_FOUND)
	ERROR_USER_NOT_FOUND                = errors.New(ERROR_MESSAGE_USER_NOT_FOUND)
	ERROR_ACTIVITY_PUB_OBJECT_NOT_FOUND = errors.New(ERROR_MESSAGE_ACTIVITY_PUB_OBJECT_NOT_FOUND)
	ERROR_SPI_NOT_FOUND                 = errors.New(ERROR_MESSAGE_SPI_NOT_FOUND)
	ERROR_UNKNOWN_DOCUMENT              = errors.New(ERROR_MESSAGE_UNKNOWN_DOCUMENT)
	ERROR_UNKNOWN_SYSTEM_SCHEDULE       = errors.New(ERROR_MESSAGE_UNKNOWN_SYSTEM_SCHEDULE)
	ERROR_UNKNOWN_IDENTITY              = errors.New(ERROR_MESSAGE_UNKNOWN_IDENTITY)
	ERROR_UNKNOWN_ACCESS_CONTROL        = errors.New(ERROR_MESSAGE_UNKNOWN_ACCESS_CONTROL)
	ERROR_UNKNOWN_REMOTE_ACTIVITY       = errors.New(ERROR_MESSAGE_UNKNOWN_REMOTE_ACTIVITY)
	ERROR_UNKNOWN_CATEGORY_TYPE         = errors.New(ERROR_MESSAGE_UNKNOWN_CATEGORY_TYPE)
	ERROR_UNKNOWN_CATEGORY              = errors.New(ERROR_MESSAGE_UNKNOWN_CATEGORY)
	ERROR_UNKNOWN_USER                  = errors.New(ERROR_MESSAGE_UNKNOWN_USER)
	ERROR_UNKNOWN_ACTIVITY_PUB_OBJECT   = errors.New(ERROR_MESSAGE_UNKNOWN_ACTIVITY_PUB_OBJECT)
	ERROR_UNKNOWN_SPI                   = errors.New(ERROR_MESSAGE_UNKNOWN_SPI)
	ERROR_INVALID_ID                    = errors.New(ERROR_MESSAGE_INVALID_ID)
	ERROR_INVALID_PARAMETERS            = errors.New(ERROR_MESSAGE_INVALID_PARAMETERS)
	// CUSTOM_ERRORS
	ERROR_DATA_INTEGRITY_VIOLATION             = errors.New(ERROR_MESSAGE_DATA_INTEGRITY_VIOLATION)
	ERROR_INVALID_STATE                        = errors.New(ERROR_MESSAGE_INVALID_STATE)
	ERROR_USER_NOT_REGISTERED                  = errors.New(ERROR_MESSAGE_USER_NOT_REGISTERED)
	ERROR_USERNAME_OR_EMAIL_ALREADY_REGISTERED = errors.New(ERROR_MESSAGE_USERNAME_OR_EMAIL_ALREADY_REGISTERED)
	ERROR_ACCOUNT_NOT_VERIFIED                 = errors.New(ERROR_MESSAGE_ACCOUNT_NOT_VERIFIED)
	ERROR_ACCOUNT_BLOCKED                      = errors.New(ERROR_MESSAGE_ACCOUNT_BLOCKED)
	ERROR_INVALID_TOKEN                        = errors.New(ERROR_MESSAGE_INVALID_TOKEN)
	ERROR_INVALID_CONFIRMATION_CODE            = errors.New(ERROR_MESSAGE_INVALID_CONFIRMATION_CODE)
	ERROR_PERMISSION_DENIED                    = errors.New(ERROR_MESSAGE_PERMISSION_DENIED)
	ERROR_INVALID_PERSON_KIND                  = errors.New(ERROR_MESSAGE_INVALID_PERSON_KIND)
	ERROR_INVALID_CREDENTIALS                  = errors.New(ERROR_MESSAGE_INVALID_CREDENTIALS)
)
