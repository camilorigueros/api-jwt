package logger

import (
	"log"
	"os"
)

// Logger es una estructura que permite personalizar el logging.
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// New crea una nueva instancia de Logger.
func New() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info registra un mensaje informativo.
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Error registra un mensaje de error.
func (l *Logger) Error(err error) {
	l.errorLogger.Println(err.Error())
}

// Log es una instancia global de Logger que puedes reutilizar en tu proyecto.
var Log = New()
