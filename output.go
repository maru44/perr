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

	/* client */
	BadRequest                  = errors.New("Bad Request")                   // 400
	UserDoesNotExists           = errors.New("User Does Not Exists")          // 400
	InvalidToken                = errors.New("Invalid Token")                 // 400
	PasswordNotMatch            = errors.New("Password Not Match")            // 400
	Expired                     = errors.New("Expired")                       // 400
	Unauthorized                = errors.New("Unauthorized")                  // 401
	PaymentRequired             = errors.New("Payment Required")              // 402
	Forbidden                   = errors.New("Forbidden")                     // 403
	NotFound                    = errors.New("Not Found")                     // 404
	MethodNotAllowed            = errors.New("Method Not Allowed")            // 405
	NotAcceptable               = errors.New("Not Acceptable")                // 406
	ProxyAuthenticationRequired = errors.New("Proxy Authentication Required") // 408
	Conflict                    = errors.New("Conflict")                      // 410
	Gone                        = errors.New("Gone")                          // 410

	IntervalServerError = errors.New("")
)
