package util

import "log"

//LogErrOrResp logs if error or print if resp
func LogErrOrResp(resp interface{}, err error) {
	if err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}
