package moksha

type Authen struct {
	UserName string `json:"username"`
	Pass     string `json:"password"`
}

type AuthenResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	TcbsID    string `json:"tcbsid"`
	Token     string `json:"token"`
}

type Order struct {
	TcbsID      string      `json:"tcbsId"`
	CustodyCode string      `json:"custodyCode"`
	AccountNo   string      `json:"accountNo"`
	AccountType string      `json:"accountType"`
	Symbol      string      `json:"symbol"`
	Volume      int         `json:"volume"`
	StartDate   string      `json:"startDate"`
	EndDate     string      `json:"endDate"`
	ExecType    string      `json:"execType"`
	Conditions  []Condition `json:"conditions"`
}

type Condition struct {
	Code       string   `json:"code"`
	CostPrice  float64  `json:"costPrice"`
	Value      float64  `json:"value"`
	ValueType  string   `json:"valueType"`
	OrderPrice *float64 `json:"orderPrice"`
	Type       string   `json:"type"`
	PriceType  string   `json:"priceType"`
}

type OrderResponse struct {
	ID             int                `json:"id"`
	TcbsID         string             `json:"tcbsId"`
	CustodyCode    string             `json:"custodyCode"`
	AccountNo      string             `json:"accountNo"`
	AccountType    string             `json:"accountType"`
	Symbol         string             `json:"symbol"`
	Volume         int                `json:"volume"`
	Status         string             `json:"status"`
	OrderType      string             `json:"orderType"`
	StartDate      string             `json:"startDate"`
	EndDate        string             `json:"endDate"`
	ExecType       string             `json:"execType"`
	TriggerPrice   *float64           `json:"triggerPrice"`
	TriggerDate    *string            `json:"triggerDate"`
	FlexOrderPrice *string            `json:"flexOrderPrice"` //giá đặt lên flex
	MatchedVolume  *int64             `json:"matchedVolume"`
	CreatedOn      string             `json:"createdOn"`
	ModifiedOn     *string            `json:"modifiedOn"`
	Conditions     []ConditionRespose `json:"conditions"`
	FlexOrder      *FrOrder           `json:"flex_order"`
}

type ConditionRespose struct {
	Code       string   `json:"code"`
	Value      float64  `json:"value"`
	ValueType  string   `json:"valueType"`
	OrderPrice *float64 `json:"orderPrice"`
	Type       string   `json:"type"`
	PriceType  string   `json:"priceType"`
	CostPrice  float64  `json:"costPrice"`
	Matched    bool     `json:"matched"`
}

type FrOrder struct {
	Object        string  `json:"object"`
	OrderId       string  `json:"orderID"`
	OrderIdBO     string  `json:"orderID_BO"`
	AccountNo     string  `json:"accountNo"`
	ExecType      string  `json:"execType"`
	OrderQtty     float64 `json:"orderQtty"` //Khối lượng đặt
	ExecQtty      float64 `json:"execQtty"`  //Khối lượng khớp
	BRatio        float64 `json:"bRatio"`    //Tỷ lệ ký quỹ (bao gồm phí)
	CodeId        string  `json:"codeID"`
	Symbol        string  `json:"symbol"`
	PriceType     string  `json:"priceType"`
	TxTime        string  `json:"txtime"`       //giờ đặt lệnh
	TxDate        string  `json:"txdate"`       //ngày đặt lệnh
	ExpDate       string  `json:"expDate"`      //ngày hết hạn
	TimeType      string  `json:"timeType"`     //loại thời gian hiệu lực
	OrStatus      string  `json:"orStatus"`     //Trạng thái lệnh
	FeeAcr        float64 `json:"feeAcr"`       //phí
	LimitPrice    float64 `json:"limitPrice"`   //giá đặt(giá limit)
	CancelQtty    float64 `json:"cancelQtty"`   //Khối lượng huỷ
	Via           string  `json:"via"`          //kênh giao dịch
	QuotePrice    float64 `json:"quotePrice"`   //giá đặt
	MatchPrice    float64 `json:"matchPrice"`   //giá khớp
	TaxSellAmount float64 `json:"taxSellAmout"` //thuế bán
	TradePlace    string  `json:"tradePlace"`
	MathType      string  `json:"matchType"`  //Loại khớp
	IsDisposal    string  `json:"isDisposal"` //là lệnh bán xử lý?
	IsAmend       string  `json:"isAmend"`    //cho phép sửa?
	IsCancel      string  `json:"isCancel"`   //cho phép huỷ?
	UserName      string  `json:"userName"`
	OrsOrderId    string  `json:"orsOrderID"`
	SecType       string  `json:"sectype"` //Loại chứng khoán
	IsFOOrder     string  `json:"isFOOrder"`
	OdTimeStamp   string  `json:"odTimeStamp"` //thời gian sinh lệnh
}

type TCPrice struct {
	Ticker         string  `json:"ticker"`
	CurrentPrice   float64 `json:"currentPrice"`
	ReferencePrice float64 `json:"referencePrice"`
	FloorPrice     float64 `json:"floorPrice"`
	CeilingPrice   float64 `json:"ceilingPrice"`
}
