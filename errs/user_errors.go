package errs

import "github.com/QuantumShiftX/golib/xerr"

const (
	UserErrCode xerr.ErrCode = 10000

	UserUserNotFoundError      = iota + UserErrCode // 用户不存在
	UserCreateUserErrorCode                         // 创建用户错误
	UserUserSignInErrorCode                         // 用户登录错误
	UserUserFotGotPwdErrorCode                      // 重置密码错误
	UserUserValidateErrorCode                       // 用户验证错误
	UserUserAgreementErrorCode                      // 获取用户协议错误
	UserSendCaptchaErrorCode                        // 发送验证码错误
	UserPasswordMismatchError                       // 密码不匹配
	UserAccountDisabledError                        // 账户已禁用
	UserDuplicateUsernameError                      // 用户名已存在
	UserInvalidEmailError                           // 无效的邮箱格式
	UserTokenExpiredError                           // 用户令牌过期
	UserInvalidPhoneError                           // 无效的电话格式
	UserValidationTypeError                         // 验证类型错误
	UserCredentialsError                            // 用户或者密码信息有误
	UserNameExistsError                             // 用户名字已存在
	UserValidationMatchError                        // 验证匹配失败
)

var (
	ErrUserNotFound          = xerr.New(UserUserNotFoundError, "用户不存在")
	ErrUserCreate            = xerr.New(UserCreateUserErrorCode, "创建用户失败")
	ErrUserSignIn            = xerr.New(UserUserSignInErrorCode, "用户登录失败")
	ErrUserForgotPwd         = xerr.New(UserUserFotGotPwdErrorCode, "重置密码失败")
	ErrUserValidate          = xerr.New(UserUserValidateErrorCode, "用户验证失败")
	ErrUserAgreement         = xerr.New(UserUserAgreementErrorCode, "获取用户协议失败")
	ErrSendCaptcha           = xerr.New(UserSendCaptchaErrorCode, "发送验证码失败")
	ErrUserPasswordMismatch  = xerr.New(UserPasswordMismatchError, "密码不匹配")
	ErrUserAccountDisabled   = xerr.New(UserAccountDisabledError, "账户已被禁用")
	ErrUserDuplicateUsername = xerr.New(UserDuplicateUsernameError, "用户名已存在")
	ErrUserInvalidEmail      = xerr.New(UserInvalidEmailError, "无效的邮箱格式")
	ErrUserTokenExpired      = xerr.New(UserTokenExpiredError, "用户令牌已过期")
	ErrUserInvalidPhone      = xerr.New(UserInvalidPhoneError, "无效的电话格式")
	ErrUserValidationType    = xerr.New(UserValidationTypeError, "验证类型错误")
	ErrUserCredentials       = xerr.New(UserCredentialsError, "用户或密码信息有误")
	ErrUserNameExists        = xerr.New(UserNameExistsError, "用户名字已存在")
	ErrUserValidationMatch   = xerr.New(UserValidationMatchError, "验证匹配失败")
)
