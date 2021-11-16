package service

type StarInvoking struct {
	BaseURL string
	Method  string
	Type    string
	ID      string

	Status  int
	Message string

	TxBody interface{}
	RxBody interface{}
}

type StarAPI interface {
	GetURI() string
	GetType() string

	Execute(invoking *StarInvoking) error
	Post(tx, rx interface{}) error
	Get(id string, rx interface{}) error
	Put(id string, tx, rx interface{}) error
	Delete(id string) error
}

type StarAgentService interface {
	GetAPI(selector string) (StarAPI, error)
}
