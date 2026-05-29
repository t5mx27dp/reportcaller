package reportcaller

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type Option func(*ReportCaller)

func WithLevels(levels []logrus.Level) Option {
	return func(r *ReportCaller) {
		r.levels = levels
	}
}

func WithLocationBuilder(fn LocationBuilder) Option {
	return func(r *ReportCaller) {
		r.locationBuilder = fn
	}
}

func WithField(field string) Option {
	return func(r *ReportCaller) {
		r.field = field
	}
}

func WithRootPath(path string) Option {
	return func(r *ReportCaller) {
		r.rootPath = strings.TrimRight(path, "/") + "/"
	}
}
