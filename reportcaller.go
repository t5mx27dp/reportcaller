package reportcaller

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

type LocationBuilder func(filePath string, line int) string

type ReportCaller struct {
	levels []logrus.Level

	locationBuilder LocationBuilder

	field    string
	rootPath string
}

func New(opts ...Option) *ReportCaller {
	r := &ReportCaller{
		levels: []logrus.Level{logrus.ErrorLevel},
		field:  "Location",
	}

	r.locationBuilder = func(filePath string, line int) string {
		return strings.ReplaceAll(filePath, r.rootPath+"/", "") + ":" + strconv.Itoa(line)
	}

	r.rootPath, _ = os.Getwd()

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func (r *ReportCaller) Levels() []logrus.Level {
	return r.levels
}

func (r *ReportCaller) Fire(entry *logrus.Entry) error {
	caller := getCaller()
	if caller == nil {
		return nil
	}

	entry.Data[r.field] = r.locationBuilder(caller.File, caller.Line)
	return nil
}
