package errors

var (
	InvalidParameter = newBadRequest("invalid_parameter", "パラメータが不正です。")
	SystemDefault    = newInternalServerError("sytem_default", "エラーが発生しました。")
)
