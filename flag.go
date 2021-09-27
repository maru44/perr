package perr

const (
	/* client error */

	FlagBadRequest uint32 = 1 + iota*5
	FlagUserDoesNotExist
	FlagInvalidToken
	FlagPasswordNotMatch
	FlagExpired

	FlagUnauthorized
	FlagPaymentRequired
	FlagForbidden
	FlagNotFound
	FlagMethodNotAllowed
	FlagNotAcceptable
	FlagProxyAuthenticationRequired
	FlagConflict
	FlagGone
	FlagLengthRequired
	FlagPreconditionField
	FlagPayloadTooLarge
	FlagURITooLong
	FlagUnsupportedMediaType
	FlagRangeNotSatisfiable
	FlagExpectationFailed
	FlagIAmTeaPot
	FlagMisdirectedRequest
	FlagUnprocessableEntity
	FlagLocked
	FlagFailedDependency
	FlagTooEarly
	FlagUpgradeRequire
	FlagPreconditionRequired
	FlagTooManyRequests
	FlagRequestHeaderFieldsTooLarge
	FlagUnavailableLegalReasons

	/* server error */

	FlagInternalServerError
	// FlagRDBError
	// FlagNoSQLError
	// FlagFileStorageError
	// FlagPaymentServiceError
	// FlagHTTPConnectionError

	FlagNotImplemented
	FlagBadGateway
	FlagServiceUnavailable
	FlagGatewayTimeout
	FlagHTTPVersionNotSupported
	FlagVariantAlsoNegotiates
	FlagInsufficientStorage
	FlagLoopDetected
	FlagNotExtended
	FlagNetworkAuthenticationRequired

	FlagInternalServerErrorWithUrgency
)
