package iTransact

type iTransactResponse struct {
	GatewayInterface struct {
		TransactionResponse struct {
			TransactionResult struct {
				Status                string `xml:"Status"`
				ErrorCategory         string `xml:"ErrorCategory"`
				ErrorMessage          string `xml:"ErrorMessage"`
				WarningMessage        string `xml:"WarningMessage"`
				AuthAmount            string `xml:"AuthAmount"`
				AuthCode              string `xml:"AuthCode"`
				AVSCategory           string `xml:"AVSCategory"`
				AVSResponse           string `xml:"AVSResponse"`
				Balance               string `xml:"Balance"`
				CardLevel             string `xml:"CardLevel"`
				CardName              string `xml:"CardName"`
				CVV2Response          string `xml:"CVV2Response"`
				PurchaseCardLevel     string `xml:"PurchaseCardLevel"`
				RefundTransactionType string `xml:"RefundTransactionType"`
				TimeStamp             string `xml:"TimeStamp"`
				TestMode              string `xml:"TestMode"`
				Total                 string `xml:"Total"`
				XID                   string `xml:"XID"`
				CustomerData          struct {
					BillingAddress struct {
						Address1  string `xml:"Address1"`
						City      string `xml:"City"`
						FirstName string `xml:"FirstName"`
						LastName  string `xml:"LastName"`
						State     string `xml:"State"`
						Zip       string `xml:"Zip"`
						Country   string `xml:"Country"`
						Phone     string `xml:"Phone"`
					} `xml:"BillingAddress"`
					ShippingAddress struct {
						Address1  string `xml:"Address1"`
						City      string `xml:"City"`
						FirstName string `xml:"FirstName"`
						LastName  string `xml:"LastName"`
						State     string `xml:"State"`
						Zip       string `xml:"Zip"`
						Country   string `xml:"Country"`
						Phone     string `xml:"Phone"`
					} `xml:"ShippingAddress"`
				} `xml:"CustomerData"`
			} `xml:"TransactionResult"`
		} `xml:"TransactionResponse"`
	} `xml:"GatewayInterface"`
}

type RunBatchCloseResponse struct {
	GatewayInterface struct {
		BatchCloseResponse struct {
			Status        string `xml:"Status"`
			ErrorCategory string `xml:"ErrorCategory"`
			ErrorMessage  string `xml:"ErrorMessage"`
			TimeStamp     string `xml:"TimeStamp"`
			TestMode      string `xml:"TestMode"`
			BatchList     struct {
				Batch []struct {
					BatchNumber  string `xml:"BatchNumber"`
					CreditAmount string `xml:"CreditAmount"`
					CreditCount  string `xml:"CreditCount"`
					NetAmount    string `xml:"NetAmount"`
					NetCount     string `xml:"NetCount"`
					SaleAmount   string `xml:"SaleAmount"`
					SaleCount    string `xml:"SaleCount"`
					VoidAmount   string `xml:"VoidAmount"`
					VoidCount    string `xml:"VoidCount"`
				} `xml:"Batch"`
			} `xml:"BatchList"`
		} `xml:"BatchCloseResponse"`
	} `xml:"GatewayInterface"`
}

type CreditTransaction struct {
	Total        string       `xml:"Total"`
	CustomerData CustomerData `xml:"CustomerData"`
	AccountInfo  CardAccount  `xml:"AccountInfo"`
}

type PostAuthTransaction struct {
	OperationXID string `xml:"OperationXID"`
	Total        string `xml:"Total,omitempty"`
}

type TranCredTransaction struct {
	OperationXID string `xml:"OperationXID"`
	Total        string `xml:"Total,omitempty"`
}

type TranForceTransaction struct {
	OperationXID string `xml:"OperationXID"`
	ApprovalCode string `xml:"ApprovalCode"`
	Total        string `xml:"Total,omitempty"`
}

type TranRetryTransaction struct {
	OperationXID string `xml:"OperationXID"`
	Total        string `xml:"Total,omitempty"`
	Description  string `xml:"Description,omitempty"`
}

type VoidTransaction struct {
	OperationXID string `xml:"OperationXID"`
}

type TranRefundTransaction struct {
	OperationXID string `xml:"OperationXID"`
}

type BatchClose struct {
	TransactionControl TransactionControl `xml:"TransactionControl"`
}

type TransactionControl struct {
	TestMode          string `xml:"TestMode"`
	SendMerchantEmail string `xml:"SendMerchantEmail"`
}

type AuthCreds struct {
	APICredentials APICredentials `xml:"APICredentials"`
}

type APICredentials struct {
	Username         string `xml:"Username"`
	TargetGateway    string `xml:"TargetGateway"`
	PayloadSignature string `xml:"PayloadSignature"`
}

type AuthTransaction struct {
	CustomerData CustomerData `xml:"CustomerData"`
	OrderItems   *Items       `xml:"OrderItems,omitempty"`
	Total        string       `xml:"Total,omitempty"`
	Description  string       `xml:"Description,omitempty"`
	AccountInfo  CardAccount  `xml:"AccountInfo"`
}

type CustomerData struct {
	Email           string  `xml:"Email"`
	BillingAddress  Address `xml:"BillingAddress"`
	ShippingAddress Address `xml:"ShippingAddress,omitempty"`
	CustId          string  `xml:"CustId,omitempty"`
}

type Items struct {
	Items []Item `xml:"Item,omitempty"`
}

type Item struct {
	Description string `xml:"Description,omitempty"`
	Cost        string `xml:"Cost,omitempty"`
	Qty         string `xml:"Qty,omitempty"`
}

type Address struct {
	Address1  string `xml:"Address1"`
	Address2  string `xml:"Address2,omitempty"`
	FirstName string `xml:"FirstName"`
	LastName  string `xml:"LastName"`
	City      string `xml:"City"`
	State     string `xml:"State"`
	Zip       string `xml:"Zip"`
	Country   string `xml:"Country"`
	Phone     string `xml:"Phone"`
}

type CardAccount struct {
	CardAccount CreditCard `xml:"CardAccount"`
}

type CreditCard struct {
	AccountNumber   string `xml:"AccountNumber"`
	ExpirationMonth string `xml:"ExpirationMonth"`
	ExpirationYear  string `xml:"ExpirationYear"`
	CVVNumber       string `xml:"CVVNumber,omitempty"`
}

type TestModeBlock struct {
	SendCustomerEmail string `xml:"SendCustomerEmail"`
	SendMerchantEmail string `xml:"SendMerchantEmail"`
	TestMode          string `xml:"TestMode"`
}

type RecureUpdateRequest struct {
	GatewayInterface struct {
		APICredentials APICredentials `xml:"APICredentials"`
		RecurUpdate    RecurUpdate    `xml:"RecurUpdate"`
	} `xml:"GatewayInterface"`
}

type RecurUpdate struct {
	OperationXID string        `xml:"OperationXID"`
	RemReps      *string       `xml:"RemReps,omitempty"`
	Recipe       *string       `xml:"Recipe,omitempty"`
	CustomerData *CustomerData `xml:"CustomerData,omitempty"`
	OrderItems   *[]Item       `xml:"OrderItems,omitempty"`
	Total        string        `xml:"Total,omitempty"`
	Description  string        `xml:"Description,omitempty"`
}

type OrderItems struct {
	Item []Item `xml:"Item,omitempty"`
}

type RecurUpdateReponse struct {
	GatewayInterface struct {
		RecurUpdate SendRecurUpdate `xml:"RecurUpdateResponse"`
	} `xml:"GatewayInterface"`
}

type SendRecurUpdate struct {
	Status        string `xml:"Status"`
	ErrorCategory string `xml:"ErrorCategory"`
	ErrorMessage  string `xml:"ErrorMessage"`
	TimeStamp     string `xml:"TimeStamp"`
	TestMode      string `xml:"TestMode"`
	RecurDetails  struct {
		RemReps    string `xml:"RemReps"`
		RecipeName string `xml:"RecipeName"`
		RecurTotal string `xml:"RecurTotal"`
	} `xml:"RecurDetails"`
}

type RecurDetails struct {
	GatewayInterface struct {
		RecurDetailsResponse struct {
			Status        string `xml:"Status"`
			ErrorCategory string `xml:"ErrorCategory"`
			ErrorMessage  string `xml:"ErrorMessage"`
			TimeStamp     string `xml:"TimeStamp"`
			TestMode      string `xml:"TestMode"`
			RecurDetails  struct {
				CardExpired             string `xml:"CardExpired"`
				CardExpiresWithinThirty string `xml:"CardExpiresWithinThirty"`
				CardLastFour            string `xml:"CardLastFour"`
				CardName                string `xml:"CardName"`
				RemReps                 string `xml:"RemReps"`
				RecipeName              string `xml:"RecipeName"`
				RecurTotal              string `xml:"RecurTotal"`
			} `xml:"RecurDetails"`
		} `xml:"RecurDetailsResponse"`
	} `xml:"GatewayInterface"`
}
