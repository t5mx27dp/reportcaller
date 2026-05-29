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

	rootPath string
	field    string
}

var _ logrus.Hook = (*ReportCaller)(nil)

func New(opts ...Option) *ReportCaller {
	r := &ReportCaller{}

	for _, opt := range opts {
		opt(r)
	}

	if len(r.levels) == 0 {
		r.levels = []logrus.Level{logrus.ErrorLevel}
	}

	if r.locationBuilder == nil {
		r.locationBuilder = func(filePath string, line int) string {
			return strings.ReplaceAll(filePath, r.rootPath, "") + ":" + strconv.Itoa(line)
		}
	}

	if r.rootPath == "" {
		r.rootPath, _ = os.Getwd()
		r.rootPath = strings.TrimRight(r.rootPath, "/") + "/"
	}

	if r.field == "" {
		r.field = "Location"
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
