package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"app/pkg/fs"

	"github.com/sirupsen/logrus"
)

const (
	infoOut  = "info"
	warnOut  = "warn"
	errorOut = "error"
)

type logrusLogger struct {
	outputFiles map[string]*os.File

	logger *logrus.Logger
	multi  *MultiWriter
}

func NewLogrusLogger(runMode string) (LoggerService, error) {
	ls := &logrusLogger{}

	if err := ls.initLogDirs(); err != nil {
		return nil, fmt.Errorf("log: failed to check and create log folders: %v", err)
	}

	var (
		fileName             = time.Now().Format("02-Jan-2006") + ".log"
		flag                 = os.O_CREATE | os.O_WRONLY | os.O_APPEND
		perm     os.FileMode = 0666
	)

	logLevels := [3]string{infoOut, warnOut, errorOut}
	ls.outputFiles = make(map[string]*os.File, len(logLevels))
	for _, level := range logLevels {
		file, err := os.OpenFile(fmt.Sprintf("logs/%s/%s", level, fileName), flag, perm)
		if err != nil {
			return nil, fmt.Errorf("log: failed to open logs/%s/%v file: %v", level, fileName, err)
		}

		ls.outputFiles[level] = file
	}

	ls.multi = &MultiWriter{
		Stdout: os.Stdout,
	}

	logger := logrus.New()
	logger.SetLevel(logrus.Level(getLevelByAppRunMode(runMode)))
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		PadLevelText:    true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetOutput(ls.multi)

	ls.logger = logger

	return ls, nil
}

func (l *logrusLogger) Info(args ...interface{}) {
	l.multi.File = l.outputFiles[infoOut]
	l.logger.Info(args...)
}

func (l *logrusLogger) Warn(args ...interface{}) {
	l.multi.File = l.outputFiles[warnOut]
	l.logger.Warn(args...)
}

func (l *logrusLogger) Error(args ...interface{}) {
	l.multi.File = l.outputFiles[errorOut]
	l.logger.Error(args...)
}

func (l *logrusLogger) initLogDirs() error {
	// first create the root dir so that when creating
	// level dirs there are no run-queue errors
	if err := fs.CreateDirIfNotExists("logs"); err != nil {
		return fmt.Errorf("log: failed to create root dir for logs")
	}

	var wg sync.WaitGroup

	dirsNames := [...]string{infoOut, warnOut, errorOut}
	errorsCh := make(chan error, len(dirsNames))
	wg.Add(len(dirsNames))
	for _, dirName := range dirsNames {
		go func(dirName string) {
			defer wg.Done()
			if err := fs.CreateDirIfNotExists(fmt.Sprintf("logs/%s", dirName)); err != nil {
				errorsCh <- fmt.Errorf("log: failed to create %s dir: %v", dirsNames, err)
				return
			}
		}(dirName)
	}

	// separate goroutine for reading from the
	// errorsCh and processing them
	go func() {
		wg.Wait()
		close(errorsCh)
	}()

	for err := range errorsCh {
		if err != nil {
			return err
		}
	}

	return nil
}
