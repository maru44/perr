package perr

import "github.com/pkg/errors"

var (
	/* success */

	OK                          = errors.New("OK")                            // 200
	Created                     = errors.New("Created")                       // 201
	Accepted                    = errors.New("Accepted")                      // 202
	NonAuthoritativeInformation = errors.New("Non-Authoritative Information") // 203
	NoContent                   = errors.New("No Content")                    // 204
	ResetContent                = errors.New("Reset Content")                 // 205
	PartialContent              = errors.New("Pertial Content")               // 206

	/* redirection */

	MultipleChoice    = errors.New("Multiple Choice")    // 300
	MovedPermanently  = errors.New("Moved Permanently")  // 301
	Found             = errors.New("Found")              // 302
	SeeOther          = errors.New("See Other")          // 303
	NotModified       = errors.New("Not Modified")       // 304
	TemporaryRedirect = errors.New("Temporary Redirect") // 307
	PermanentRedirect = errors.New("Permanent Redirect") // 308

	/* client error */

	BadRequest                  = errors.New("Bad Request")                     // 400
	UserDoesNotExist            = errors.New("User Does Not Exist")             // 400
	InvalidToken                = errors.New("Invalid Token")                   // 400
	PasswordNotMatch            = errors.New("Password Not Match")              // 400
	Expired                     = errors.New("Expired")                         // 400
	Unauthorized                = errors.New("Unauthorized")                    // 401
	PaymentRequired             = errors.New("Payment Required")                // 402
	Forbidden                   = errors.New("Forbidden")                       // 403
	NotFound                    = errors.New("Not Found")                       // 404
	MethodNotAllowed            = errors.New("Method Not Allowed")              // 405
	NotAcceptable               = errors.New("Not Acceptable")                  // 406
	ProxyAuthenticationRequired = errors.New("Proxy Authentication Required")   // 408
	Conflict                    = errors.New("Conflict")                        // 409
	Gone                        = errors.New("Gone")                            // 410
	LengthRequired              = errors.New("Length Required")                 // 411
	PreconditionField           = errors.New("Precondition Failed")             // 412
	PayloadTooLarge             = errors.New("Payload Too Large")               // 413
	URITooLong                  = errors.New("URI Too Large")                   // 414
	UnsupportedMediaType        = errors.New("Unsupported Media Type")          // 415
	RangeNotSatisfiable         = errors.New("Range Not Satisfiable")           // 416
	ExpectationFailed           = errors.New("Expectation Failed")              // 417
	IAmATeaPot                  = errors.New("I'm a teapot")                    // 418
	MisdirectedRequest          = errors.New("Misdirected Request")             // 421
	UnprocessableEntity         = errors.New("Unprocessable Entity")            // 422
	Locked                      = errors.New("Locked")                          // 423
	FailedDependency            = errors.New("Failed Dependency")               // 424
	TooEarly                    = errors.New("TooEarly")                        // 425
	UpgradeRequire              = errors.New("Upgrade Required")                // 426
	PreconditionRequired        = errors.New("Precondition Required")           // 428
	TooManyRequests             = errors.New("Too Many Requests")               // 429
	RequestHeaderFieldsTooLarge = errors.New("Request Header Fields Too Large") // 431
	UnavailableForLegalReasons  = errors.New("Unavailable For Legal Reasons")   // 451

	/* server error */

	IntervalServerError           = errors.New("Internal Server Error")           // 500
	NotImplemented                = errors.New("NotImplemented")                  // 501
	BadGateway                    = errors.New("BadGateway")                      // 502
	ServiceUnavailable            = errors.New("Service Unavailable")             // 503
	GatewayTimeout                = errors.New("Gateway Timeout")                 // 504
	HTTPVersionNotSupported       = errors.New("HTTP Version Not Supported")      // 505
	VariantAlsoNegotiates         = errors.New("Variant Also Negotiates")         // 506
	InsufficientStorage           = errors.New("Insufficient Storage")            // 507
	LoopDetected                  = errors.New("Loop Detected")                   // 508
	NotExtended                   = errors.New("Not Extended")                    // 510
	NetworkAuthenticationRequired = errors.New("Network Authentication Required") // 511
)

func (e Err) Output() error {
	if e.textOutput != "" {
		return errors.New(e.textOutput)
	}

	return e.error()
}

func (e Err) error() error {
	switch e.Flag {
	case FlagBadRequest:
		return BadRequest
	case FlagUserDoesNotExist:
		return UserDoesNotExist
	case FlagInvalidToken:
		return InvalidToken
	case FlagPasswordNotMatch:
		return PasswordNotMatch
	case FlagExpired:
		return Expired
	case FlagUnauthorized:
		return Unauthorized
	case FlagPaymentRequired:
		return PaymentRequired
	case FlagForbidden:
		return Forbidden
	case FlagNotFound:
		return NotFound
	case FlagMethodNotAllowed:
		return MethodNotAllowed
	case FlagNotAcceptable:
		return NotAcceptable
	case FlagProxyAuthenticationRequired:
		return NetworkAuthenticationRequired
	case FlagConflict:
		return Conflict
	case FlagGone:
		return Gone
	case FlagLengthRequired:
		return LengthRequired
	case FlagPreconditionField:
		return PreconditionField
	case FlagPayloadTooLarge:
		return PayloadTooLarge
	case FlagURITooLong:
		return URITooLong
	case FlagUnsupportedMediaType:
		return UnsupportedMediaType
	case FlagRangeNotSatisfiable:
		return RangeNotSatisfiable
	case FlagExpectationFailed:
		return ExpectationFailed
	case FlagIAmTeaPot:
		return IAmATeaPot
	case FlagMisdirectedRequest:
		return MisdirectedRequest
	case FlagUnprocessableEntity:
		return UnprocessableEntity
	case FlagLocked:
		return Locked
	case FlagFailedDependency:
		return FailedDependency
	case FlagTooEarly:
		return TooEarly
	case FlagUpgradeRequire:
		return UpgradeRequire
	case FlagPreconditionRequired:
		return PreconditionRequired
	case FlagTooManyRequests:
		return TooManyRequests
	case FlagRequestHeaderFieldsTooLarge:
		return RequestHeaderFieldsTooLarge
	case FlagUnavailableLegalReasons:
		return UnavailableForLegalReasons

	case FlagInternalServerError:
		return IntervalServerError
	case FlagBadGateway:
		return BadGateway
	case FlagServiceUnavailable:
		return ServiceUnavailable
	case FlagGatewayTimeout:
		return GatewayTimeout
	case FlagHTTPVersionNotSupported:
		return HTTPVersionNotSupported
	case FlagVariantAlsoNegotiates:
		return VariantAlsoNegotiates
	case FlagInsufficientStorage:
		return InsufficientStorage
	case FlagLoopDetected:
		return LoopDetected
	case FlagNotExtended:
		return NotExtended
	case FlagNetworkAuthenticationRequired:
		return NetworkAuthenticationRequired

	case FlagInternalServerErrorWithUrgency:
		return IntervalServerError
	default:
		return IntervalServerError
	}
}
