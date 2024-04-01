package errors

type ErrCode uint

const (
	ErrUnknown                      ErrCode = 0
	ErrBadRequest                   ErrCode = 400
	ErrUnauthorized                 ErrCode = 401
	ErrForbidden                    ErrCode = 403
	ErrNotFound                     ErrCode = 404
	ErrMethodNotAllowed             ErrCode = 405
	ErrNotAcceptable                ErrCode = 406
	ErrProxyAuthRequired            ErrCode = 407
	ErrRequestTimeout               ErrCode = 408
	ErrConflict                     ErrCode = 409
	ErrGone                         ErrCode = 410
	ErrLengthRequired               ErrCode = 411
	ErrPreconditionFailed           ErrCode = 412
	ErrRequestEntityTooLarge        ErrCode = 413
	ErrRequestURITooLong            ErrCode = 414
	ErrUnsupportedMediaType         ErrCode = 415
	ErrRequestedRangeNotSatisfiable ErrCode = 416
	ErrTooManyRequests              ErrCode = 429
	ErrUnexpected                   ErrCode = 500
	ErrServiceUnavailable           ErrCode = 503
	ErrGatewayTimeout               ErrCode = 504
)
