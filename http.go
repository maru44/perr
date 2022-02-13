package perr

import "errors"

var (
	/* success */

	SuccessOk                          = errors.New("OK")                            // 200
	SuccessCreated                     = errors.New("Created")                       // 201
	SuccessAccepted                    = errors.New("Accepted")                      // 202
	SuccessNonAuthoritativeInformation = errors.New("Non-Authoritative Information") // 203
	SuccessNoContent                   = errors.New("No Content")                    // 204
	SuccessResetContent                = errors.New("Reset Content")                 // 205
	SuccessPartialContent              = errors.New("Pertial Content")               // 206

	/* redirection */

	RedirectionMultipleChoice    = errors.New("Multiple Choice")    // 300
	RedirectionMovedPermanently  = errors.New("Moved Permanently")  // 301
	RedirectionFound             = errors.New("Found")              // 302
	RedirectionSeeOther          = errors.New("See Other")          // 303
	RedirectionNotModified       = errors.New("Not Modified")       // 304
	RedirectionTemporaryRedirect = errors.New("Temporary Redirect") // 307
	RedirectionPermanentRedirect = errors.New("Permanent Redirect") // 308

	/* client error */

	ErrBadRequest                  = errors.New("Bad Request")                   // 400
	ErrUserDoesNotExist            = errors.New("User Does Not Exist")           // 400
	ErrInvalidToken                = errors.New("Invalid Token")                 // 400
	ErrPasswordNotMatch            = errors.New("Password Not Match")            // 400
	ErrExpired                     = errors.New("Expired")                       // 400
	ErrUnauthorized                = errors.New("Unauthorized")                  // 401
	ErrPaymentRequired             = errors.New("Payment Required")              // 402
	ErrForbidden                   = errors.New("Forbidden")                     // 403
	ErrNotFound                    = errors.New("Not Found")                     // 404
	ErrMethodNotAllowed            = errors.New("Method Not Allowed")            // 405
	ErrNotAcceptable               = errors.New("Not Acceptable")                // 406
	ErrProxyAuthenticationRequired = errors.New("Proxy Authentication Required") // 408
	ErrConflict                    = errors.New("Conflict")                      // 409
	ErrGone                        = errors.New("Gone")                          // 410
	ErrLengthRequired              = errors.New("Length Required")               // 411
	ErrPreconditionField           = errors.New("Precondition Failed")           // 412
	ErrPayloadTooLarge             = errors.New("Payload Too Large")             // 413
	ErrURITooLong                  = errors.New("URI Too Large")                 // 414
	ErrUnsupportedMediaType        = errors.New("Unsupported Media Type")        // 415
	ErrRangeNotSatisfiable         = errors.New("Range Not Satisfiable")         // 416
	ErrExpectationFailed           = errors.New("Expectation Failed")            // 417
	ErrIAmATeaPot                  = errors.New("I'm a teapot")                  // 418
	ErrCorsError                   = errors.New("Cors error")
	ErrMisdirectedRequest          = errors.New("Misdirected Request")             // 421
	ErrUnprocessableEntity         = errors.New("Unprocessable Entity")            // 422
	ErrLocked                      = errors.New("Locked")                          // 423
	ErrFailedDependency            = errors.New("Failed Dependency")               // 424
	ErrTooEarly                    = errors.New("TooEarly")                        // 425
	ErrUpgradeRequire              = errors.New("Upgrade Required")                // 426
	ErrPreconditionRequired        = errors.New("Precondition Required")           // 428
	ErrTooManyRequests             = errors.New("Too Many Requests")               // 429
	ErrRequestHeaderFieldsTooLarge = errors.New("Request Header Fields Too Large") // 431
	ErrUnavailableForLegalReasons  = errors.New("Unavailable For Legal Reasons")   // 451

	/* server error */

	ErrInternalServerError           = errors.New("Internal Server Error")           // 500
	ErrNotImplemented                = errors.New("NotImplemented")                  // 501
	ErrBadGateway                    = errors.New("BadGateway")                      // 502
	ErrServiceUnavailable            = errors.New("Service Unavailable")             // 503
	ErrGatewayTimeout                = errors.New("Gateway Timeout")                 // 504
	ErrHTTPVersionNotSupported       = errors.New("HTTP Version Not Supported")      // 505
	ErrVariantAlsoNegotiates         = errors.New("Variant Also Negotiates")         // 506
	ErrInsufficientStorage           = errors.New("Insufficient Storage")            // 507
	ErrLoopDetected                  = errors.New("Loop Detected")                   // 508
	ErrNotExtended                   = errors.New("Not Extended")                    // 510
	ErrNetworkAuthenticationRequired = errors.New("Network Authentication Required") // 511

	ErrInternalServerErrorWithUrgency           = errors.New("Internal Server Error")           // 500
	ErrNotImplementedWithUrgency                = errors.New("NotImplemented")                  // 501
	ErrBadGatewayWithUrgency                    = errors.New("BadGateway")                      // 502
	ErrServiceUnavailableWithUrgency            = errors.New("Service Unavailable")             // 503
	ErrGatewayTimeoutWithUrgency                = errors.New("Gateway Timeout")                 // 504
	ErrHTTPVersionNotSupportedWithUrgency       = errors.New("HTTP Version Not Supported")      // 505
	ErrVariantAlsoNegotiatesWithUrgency         = errors.New("Variant Also Negotiates")         // 506
	ErrInsufficientStorageWithUrgency           = errors.New("Insufficient Storage")            // 507
	ErrLoopDetectedWithUrgency                  = errors.New("Loop Detected")                   // 508
	ErrNotExtendedWithUrgency                   = errors.New("Not Extended")                    // 510
	ErrNetworkAuthenticationRequiredWithUrgency = errors.New("Network Authentication Required") // 511
)
