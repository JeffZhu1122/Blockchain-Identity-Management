package lib

// Account - admin users
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //no
	Balance   float64 `json:"balance"`   //no

	UName string `json:"uname"`
	Password   string `json:"password"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Photo string `json:"photo"`
	Branch string `json:"branch"`
	Level int `json:"level"`
}

// RealEstate - people info
type RealEstate struct {
	RealEstateID string  `json:"realEstateId"` //people id
	Proprietor   string  `json:"proprietor"`   //no
	Encumbrance  bool    `json:"encumbrance"`  //no
	TotalArea    float64 `json:"totalArea"`    //no
	LivingSpace  float64 `json:"livingSpace"`  //no

	Firstname string  `json:"firstname"`
	Lastname string  `json:"lastname"`
	DOB string  `json:"dob"`
	Gender string  `json:"gender"`
	Status string  `json:"status"`
	Photo string `json:"photo"`
}

// Selling - records
type Selling struct {
	ObjectOfSale  string  `json:"objectOfSale"`  //people id
	Seller        string  `json:"seller"`        //no
	Buyer         string  `json:"buyer"`         //no
	Price         float64 `json:"price"`         //no
	CreateTime    string  `json:"createTime"`    //no
	SalePeriod    int     `json:"salePeriod"`    //no
	SellingStatus string  `json:"sellingStatus"` //no

	AgentID        string  `json:"agentId"`
	BranchID         string  `json:"branchId"`
	RecordID  string  `json:"recordId"`
	From string  `json:"from"` 
	FromDate string  `json:"fromDate"` 
	To string  `json:"to"` 
	ToDate string  `json:"toDate"` 

}

// SellingStatusConstant 销售状态
var SellingStatusConstant = func() map[string]string {
	return map[string]string{
		"saleStart": "销售中", //正在销售状态,等待买家光顾
		"cancelled": "已取消", //被卖家取消销售或买家退款操作导致取消
		"expired":   "已过期", //销售期限到期
		"delivery":  "交付中", //买家买下并付款,处于等待卖家确认收款状态,如若卖家未能确认收款，买家可以取消并退款
		"done":      "完成",  //卖家确认接收资金，交易完成
	}
}

// SellingBuy - vaccination
type SellingBuy struct {
	Buyer      string  `json:"buyer"`      //参与销售人、买家(买家AccountId)
	CreateTime string  `json:"createTime"` //创建时间
	Selling    Selling `json:"selling"`    //销售对象
}

// Donating - branch info
type Donating struct {
	ObjectOfDonating string `json:"objectOfDonating"` //捐赠对象(正在捐赠的房地产RealEstateID)
	Donor            string `json:"donor"`            //捐赠人(捐赠人AccountId)
	Grantee          string `json:"grantee"`          //受赠人(受赠人AccountId)
	CreateTime       string `json:"createTime"`       //创建时间
	DonatingStatus   string `json:"donatingStatus"`   //捐赠状态
	
	BranchID string  `json:"branchId"`
	BName string  `json:"banme"` 
	Location string  `json:"location"` 
}

// DonatingStatusConstant 捐赠状态
var DonatingStatusConstant = func() map[string]string {
	return map[string]string{
		"donatingStart": "捐赠中", //捐赠人发起捐赠合约，等待受赠人确认受赠
		"cancelled":     "已取消", //捐赠人在受赠人确认受赠之前取消捐赠或受赠人取消接收受赠
		"done":          "完成",  //受赠人确认接收，交易完成
	}
}

// DonatingGrantee 
type DonatingGrantee struct {
	Grantee    string   `json:"grantee"`    //no
	CreateTime string   `json:"createTime"` //no
	Donating   Donating `json:"donating"`   //no
}

const (
	AccountKey         = "account-key"
	RealEstateKey      = "real-estate-key"
	SellingKey         = "selling-key"
	SellingBuyKey      = "selling-buy-key"
	DonatingKey        = "donating-key"
	DonatingGranteeKey = "donating-grantee-key"
)
