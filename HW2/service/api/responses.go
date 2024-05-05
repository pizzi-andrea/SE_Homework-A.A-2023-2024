package api

import "net/http"

var (
	UnauthorizedError http.Response = http.Response{

		Status:     "Access token is missing or invalid",
		StatusCode: http.StatusUnauthorized,
	}

	UnauthorizedToken http.Response = http.Response{
		Status:     "the submitted token does not have the necessary permissions to access this resource",
		StatusCode: http.StatusForbidden,
	}

	ServerError http.Response = http.Response{
		Status:     "server internal error, for more information see log file",
		StatusCode: http.StatusInternalServerError,
	}

	BadRequest http.Response = http.Response{
		Status:     "Bad request, syntax errors",
		StatusCode: http.StatusBadRequest,
	}
)
