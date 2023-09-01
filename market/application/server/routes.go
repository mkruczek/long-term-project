package server

import (
	"fmt"
	"market/market/application/handlers"
)

func (svr *Server) Routes() {

	svr.GET(fmt.Sprintf("%s/upload/xtb", svr.UrlPath), handlers.XtbUpload(svr.services.XtbProvider))
}
