package errors

// 错误码
const (
	// 通用错误码
	CodeGeneral = 10000
	// 参数错误
	CodeInvalidParam = 11000
	// 数据库错误
	CodeDB = 12000
	// 暂未支持功能
	CodeNotSupport = 13000

	// 认证相关
	CodeAuth = 20000
	// 验证参数无效
	CodeAuthInvalidParam = 20001
	// 密码错误
	CodeAuthWrongPassword = 20002
	// 账号已经存在
	CodeAuthAccountExist = 20003
	// 账号不存在
	CodeAccountNotExist = 20004
	// 账号已经认证
	CodeAccountAuthenticationExist = 20005
	// 账号未认证
	CodeAccountUnAuthenticationExist = 20006
	// 账号未登录
	CodeAccountNotLogin = 20007
	// 账号类型错误
	CodeAccountWrongType = 20008
	// 认证token相关
	CodeAuthToken = 21000
	// 未找到token
	CodeAuthTokenNotFound = 21001
	// token已过期
	CodeAuthTokenExpired = 21002
	// token无效
	CodeAuthTokenInvalid = 21003
	// 应用相关
	CodeAuthApplication = 22000
	// 未找到应用
	CodeAuthApplicationNotFound = 22001
	// 应用已禁用
	CodeAuthApplicationInactive = 22002
	// 未找到code
	CodeAuthCodeNotFound = 23002
	// 消息已存在
	CodeAuthMessageFound = 24001
	// 消息不存在
	CodeAuthMessageNotFound = 24002
)
