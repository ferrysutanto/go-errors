package errors

type ErrCode uint

const (
	ErrorUnknown                      ErrCode = 0
	ErrNotFound                       ErrCode = 404
	ErrBadRequest                     ErrCode = 400
	ErrUnauthorized                   ErrCode = 401
	ErrForbidden                      ErrCode = 403
	ErrorConflict                     ErrCode = 409
	ErrorInternalServer               ErrCode = 500
	ErrorServiceUnavailable           ErrCode = 503
	ErrorGatewayTimeout               ErrCode = 504
	ErrorTooManyRequests              ErrCode = 429
	ErrorBadRequest                   ErrCode = 400
	ErrorUnauthorized                 ErrCode = 401
	ErrorForbidden                    ErrCode = 403
	ErrorNotFound                     ErrCode = 404
	ErrorMethodNotAllowed             ErrCode = 405
	ErrorNotAcceptable                ErrCode = 406
	ErrorProxyAuthRequired            ErrCode = 407
	ErrorRequestTimeout               ErrCode = 408
	ErrorLengthRequired               ErrCode = 411
	ErrorPreconditionFailed           ErrCode = 412
	ErrorRequestEntityTooLarge        ErrCode = 413
	ErrorRequestURITooLong            ErrCode = 414
	ErrorUnsupportedMediaType         ErrCode = 415
	ErrorRequestedRangeNotSatisfiable ErrCode = 416
	ErrUnexpected                     ErrCode = 500
)
