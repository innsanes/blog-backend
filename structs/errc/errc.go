package errc

import (
	g "blog-backend/global"
	"go.uber.org/zap"
)

func Handle(desc string, err error, fields ...zap.Field) error {
	if err != nil {
		g.Log.WithOptions(zap.AddCallerSkip(1)).Error(desc, append(fields, zap.Error(err))...)
	}
	return err
}
