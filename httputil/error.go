package httputil

func BadRequest(msg string) ResponseHTTP {
	return ResponseHTTP{false, nil, msg}
}

func InternalError(msg string) ResponseHTTP {
	return ResponseHTTP{false, nil, msg}
}
