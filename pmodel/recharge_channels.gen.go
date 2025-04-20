package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameRechargeChannel = "recharge_channels"

// RechargeChannel 映射自表 <recharge_channels>
type RechargeChannel struct {
	ID                      int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                         // 主键
	RechargeChannelName     string  `gorm:"column:recharge_channel_name;not null;comment:充值通道名称" json:"recharge_channel_name"`                    // 充值通道名称
	RechargeChannelType     int8    `gorm:"column:recharge_channel_type;not null;comment:充值通道类型" json:"recharge_channel_type"`                    // 充值通道类型
	RechargeChannelCategory int8    `gorm:"column:recharge_channel_category;default:1;comment:充值通道分类" json:"recharge_channel_category"`           // 充值通道分类
	RechargeMerchantID      int64   `gorm:"column:recharge_merchant_id;not null;comment:充值商户配置id" json:"recharge_merchant_id"`                    // 充值商户配置id
	RechargeCategoryID      int64   `gorm:"column:recharge_category_id;not null;comment:充值大类id" json:"recharge_category_id"`                      // 充值大类id
	CurrencyCode            string  `gorm:"column:currency_code;size:255;comment:币种" json:"currency_code"`                                        // 币种
	RechargeChannelCode     string  `gorm:"column:recharge_channel_code;size:255;comment:充值通道编码" json:"recharge_channel_code"`                    // 充值通道编码
	UserLevels              string  `gorm:"column:user_levels;type:jsonb;default:'{}';comment:会员层级" json:"user_levels"`                           // 会员层级
	PlatformTypes           string  `gorm:"column:platform_types;type:jsonb;default:'{}';comment:充值平台类型" json:"platform_types"`                   // 充值平台类型
	MinLimitAmount          int64   `gorm:"column:min_limit_amount;not null;comment:单笔最小限额" json:"min_limit_amount"`                              // 单笔最小限额
	MaxLimitAmount          float64 `gorm:"column:max_limit_amount;type:numeric;not null;comment:单笔最大限额" json:"max_limit_amount"`                 // 单笔最大限额
	RequireRealName         int8    `gorm:"column:require_real_name;not null;comment:是否需要输入姓名" json:"require_real_name"`                          // 是否需要输入姓名
	Status                  int8    `gorm:"column:status;not null;comment:通道启用状态" json:"status"`                                                  // 通道启用状态
	FeePercent              float64 `gorm:"column:fee_percent;type:numeric;not null;comment:手续费（百分比）" json:"fee_percent"`                         // 手续费（百分比）
	CornerMarker            string  `gorm:"column:corner_marker;size:255;comment:角标" json:"corner_marker"`                                        // 角标
	FeeReductionConfig      string  `gorm:"column:fee_reduction_config;type:jsonb;default:'{}';comment:手续费减免配置" json:"fee_reduction_config"`      // 手续费减免配置
	RechargeGiftConfig      string  `gorm:"column:recharge_gift_config;type:jsonb;default:'{}';comment:赠送配置" json:"recharge_gift_config"`         // 赠送配置
	RecommendAmountConfig   string  `gorm:"column:recommend_amount_config;type:jsonb;default:'{}';comment:推荐金额配置" json:"recommend_amount_config"` // 推荐金额配置
	Hint                    string  `gorm:"column:hint;type:text;comment:提示" json:"hint"`                                                         // 提示
	Remark                  string  `gorm:"column:remark;type:text;comment:备注" json:"remark"`                                                     // 备注
	TopUpAt                 int64   `gorm:"column:top_up_at;comment:置顶时间" json:"top_up_at"`                                                       // 置顶时间
	ReceiveAccount          string  `gorm:"column:receive_account;size:255;comment:收款账号/地址" json:"receive_account"`                               // 收款账号/地址
	ReceiveName             string  `gorm:"column:receive_name;size:255;comment:收款人姓名" json:"receive_name"`                                       // 收款人姓名
	QrCode                  string  `gorm:"column:qr_code;type:text;comment:收款二维码" json:"qr_code"`                                                // 收款二维码
	UploadUtr               int64   `gorm:"column:upload_utr;comment:会员充值上传utr" json:"upload_utr"`                                                // 会员充值上传utr
	RandomDecimalStatus     int8    `gorm:"column:random_decimal_status;comment:是否开启随机小数" json:"random_decimal_status"`                           // 是否开启随机小数
	BankID                  int64   `gorm:"column:bank_id;comment:银行id" json:"bank_id"`                                                           // 银行id
	Ifsc                    string  `gorm:"column:ifsc;size:255;comment:银行ifsc" json:"ifsc"`                                                      // 银行ifsc
	Protocol                string  `gorm:"column:protocol;size:255;comment:协议类型" json:"protocol"`                                                // 协议类型
	TransferRemarkStatus    int8    `gorm:"column:transfer_remark_status;comment:转账备注状态" json:"transfer_remark_status"`                           // 转账备注状态
	RechargeWalletName      string  `gorm:"column:recharge_wallet_name;size:255;comment:三方钱包名称" json:"recharge_wallet_name"`                      // 三方钱包名称
	CustomerType            int8    `gorm:"column:customer_type;comment:客服类型" json:"customer_type"`                                               // 客服类型
	CustomerRechargeTypes   string  `gorm:"column:customer_recharge_types;type:jsonb;default:'{}';comment:充值类型" json:"customer_recharge_types"`   // 充值类型
	CustomerLinkType        int8    `gorm:"column:customer_link_type;comment:客服链接类型" json:"customer_link_type"`                                   // 客服链接类型
	CustomerLink            string  `gorm:"column:customer_link;size:255;comment:客服链接" json:"customer_link"`                                      // 客服链接
	CustomerCurrencyCodes   string  `gorm:"column:customer_currency_codes;type:jsonb;default:'{}';comment:货币" json:"customer_currency_codes"`     // 货币
	UsdtExchangeRate        int64   `gorm:"column:usdt_exchange_rate;comment:USDT兑换汇率（固定1000个指定货币兑换多少ustd的汇率）" json:"usdt_exchange_rate"`         // USDT兑换汇率（固定1000个指定货币兑换多少ustd的汇率）
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName RechargeChannel 表名
func (*RechargeChannel) TableName() string {
	return TableNameRechargeChannel
}
