package logger

/*
	Scrumble Logger Config object
*/
type Config struct {
	// directory for *.log files
	FileDirectory string

	// log filename
	FileName      string

	// custom time/date format: zapcore.TimeEncoderOfLayout
	// TimeEncoderOfLayout returns TimeEncoder which serializes a time.Time using
	// https://github.com/uber-go/zap/pull/629
	TimeFormat    string

	// adds color to log level
	ColorLevel     bool

	// is a compressed flag for file log
	Compress       bool

	// max log file size in megabytes
	FileSize       int

	// file age in days
	FileAge        int

	// num of log files backups
	FileBackups    int
}
