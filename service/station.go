package service

import (
	"wtws-server/common"
	common_struct "wtws-server/common/common-struct"
	wtws_mysql "wtws-server/models/wtws-mysql"
)

func GetUserStation(userId int) common_struct.ResponseStruct {
	station := wtws_mysql.GetUserStation(userId)
	return common.ResponseStatus(0, "", station)
}
