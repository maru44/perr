package perr

// whether is perror or not
func IsPerror(err error) bool {
	_, ok := err.(Perror)
	return ok
}

// get level by as
func getErrLevel(as error) ErrLevel {
	var l ErrLevel
	switch as {
	case InternalServerErrorWithUrgency, NotExtendedWithUrgency, BadGatewayWithUrgency,
		ServiceUnavailableWithUrgency, GatewayTimeoutWithUrgency, HTTPVersionNotSupportedWithUrgency,
		VariantAlsoNegotiatesWithUrgency, InsufficientStorageWithUrgency, LoopDetectedWithUrgency,
		NotExtendedWithUrgency, NetworkAuthenticationRequiredWithUrgency:
		l = ErrLevelAlert
	case InternalServerError, NotImplemented, BadGateway, ServiceUnavailable,
		GatewayTimeout, HTTPVersionNotSupported, VariantAlsoNegotiates,
		InsufficientStorage, LoopDetected, NotExtended, NetworkAuthenticationRequired:
		l = ErrLevelInternal
	default:
		l = ErrLevelExternal
	}
	return l
}
