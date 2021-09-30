package perr

// "IsPerror" is function reveal whether Perror or not.
// If Perror, return Perror and true.
// If not Perror, return nil and false.
func IsPerror(err error) (Perror, bool) {
	if perror, ok := err.(Perror); ok {
		return perror, ok
	}
	return nil, false
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
