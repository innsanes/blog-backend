package errc

import (
	g "blog-backend/global"
)

func Handle(desc string, err error) error {
	if err != nil {
		g.Log.Error(desc + " err:" + err.Error())
	}
	return err
}
