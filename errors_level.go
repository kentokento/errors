package errors

type level string

/* for logging */
const (
	levelPanic  level = "panic"
	levelEmerg  level = "emerg"
	levelAlert  level = "alert"
	levelCrit   level = "crit"
	levelError  level = "error"
	levelWarn   level = "warn"
	levelNotice level = "notice"
	levelInfo   level = "info"
	levelDebug  level = "debug"
)

func (e *appError) Panic() AppError {
	e.level = levelPanic
	return e
}
func (e *appError) Crit() AppError {
	e.level = levelCrit
	return e
}
func (e *appError) Warn() AppError {
	e.level = levelWarn
	return e
}
func (e *appError) Info() AppError {
	e.level = levelInfo
	return e
}

func (e *appError) IsPanic() bool { return e.checkLevel(levelPanic) }
func (e *appError) IsCrit() bool  { return e.checkLevel(levelCrit) }
func (e *appError) IsWarn() bool  { return e.checkLevel(levelWarn) }
func (e *appError) IsInfo() bool  { return e.checkLevel(levelInfo) }

func (e *appError) checkLevel(lv level) bool {
	if e.level != `` {
		return e.level == lv
	}
	next := AsAppError(e.next)
	if next != nil {
		return next.checkLevel(lv)
	}
	// Default level is critical
	return lv == levelCrit
}
