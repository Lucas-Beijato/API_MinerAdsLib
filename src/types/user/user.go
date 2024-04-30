package usertype

import reqkiwifywhtype "ApiExtention.com/src/types/req_kiwify_wh"

type User struct {
	Data_User       *reqkiwifywhtype.Req_Kiwify_Wh_Type `bson:"Data_User"`
	Token           string                              `bson:"token"`
	Subscription_ID string                              `json:"subscription_id"`
}
