package core

const (
	GetTokenRequestUrl      = "https://oapi.dingtalk.com/gettoken"
	ListUserIdUrl           = "https://oapi.dingtalk.com/user/getDeptMember"
	GetUserInfoUrl          = "https://oapi.dingtalk.com/user/get"
	SendCropConversationUrl = "https://oapi.dingtalk.com/topapi/message/corpconversation/asyncsend_v2"
	SendRobotMessageUrl     = "https://oapi.dingtalk.com/robot/send"
)

type BaseRequest struct {
	AccessToken string `json:"access_token"`
}

type BaseResponse struct {
	ErrorCode int    `json:"errcode"`
	ErrorMsg  string `json:"errmsg"`
}

func (b BaseResponse) Success() bool {
	return b.ErrorCode == 0
}

// 获取企业凭证 https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=/gettoken
type GetTokenRequest struct {
	AppKey    string
	AppSecret string
}
type GetTokenResponse struct {
	BaseResponse
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
}

// 查询部门userid列表  https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=/user/listid
type ListUserIdRequest struct {
	BaseRequest
	DeptId uint32
}
type ListUserIdResponse struct {
	BaseResponse
	UserIds []string `json:"userids"`
}

// 获取部门用户 https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=/user/simplelist
type ListSimpleUserRequest struct {
	BaseRequest
	DeptId uint32 `json:"dept_id"`
	Offset uint32 `json:"offset"`
	Size   uint32 `json:"size"`
	Order  string `json:"order"`
}
type SimpleUserData struct {
	Id   string `json:"userid"`
	Name string `json:"name"`
}
type ListSimpleUserResponse struct {
	BaseResponse
	HasMore  bool `json:"hasMore"`
	UserList []SimpleUserData
}

//  获取用户详情 https://open-dev.dingtalk.com/apiExplorer#/?devType=org&api=/user/get
type GetUserRequest struct {
	BaseRequest
	UserId string
}
type GetUserResponse struct {
	BaseResponse
	UnionId string `json:"unionid"`
	UserId  string `json:"userid"`
	Mobile  string `json:"mobile"`
	Email   string `json:"email"`
	Name    string `json:"name"`
}

type SendCropConversationRequest struct {
	BaseRequest
	AgentId    uint32      `json:"agent_id"`
	UserIdList []string    `json:"userid_list"`
	DeptIdList []string    `json:"dept_id_list"`
	ToAllUser  bool        `json:"to_all_user"`
	Message    interface{} `json:"message"`
}

type SendCropConversationResponse struct {
	BaseResponse
	TaskId int `json:"task_id"`
}

type SendRobotMessageRequest struct {
	Token   string
	Secret  string
	Message interface{}
}
type SendRobotMessageResponse struct {
	BaseResponse
}
