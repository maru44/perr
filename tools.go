package perr

import "reflect"

// "IsPerror" is function reveal whether Perror or not.
// If Perror, return Perror and true.
// If not Perror, return nil and false.
func IsPerror(err error) (Perror, bool) {
	if perror, ok := err.(Perror); ok {
		return perror, ok
	}
	return nil, false
}

func IsNoError(err error) bool {
	if err == nil || reflect.ValueOf(err).IsNil() {
		return true
	}
	return false
}

// get level by as
func getErrLevel(as error) ErrLevel {
	var l ErrLevel
	switch as {
	case ErrInternalServerErrorWithUrgency, ErrNotExtendedWithUrgency, ErrBadGatewayWithUrgency,
		ErrServiceUnavailableWithUrgency, ErrGatewayTimeoutWithUrgency, ErrHTTPVersionNotSupportedWithUrgency,
		ErrVariantAlsoNegotiatesWithUrgency, ErrInsufficientStorageWithUrgency, ErrLoopDetectedWithUrgency,
		ErrNotExtendedWithUrgency, ErrNetworkAuthenticationRequiredWithUrgency:
		l = ErrLevelAlert
	case ErrInternalServerError, ErrNotImplemented, ErrBadGateway, ErrServiceUnavailable,
		ErrGatewayTimeout, ErrHTTPVersionNotSupported, ErrVariantAlsoNegotiates,
		ErrInsufficientStorage, ErrLoopDetected, ErrNotExtended, ErrNetworkAuthenticationRequired:
		l = ErrLevelInternal
	default:
		l = ErrLevelExternal
	}
	return l
}
