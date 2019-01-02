package g

type HttpBindReq struct {
	Body BindReqBody `json:"body"`
	Sign string      `json:"sign"`
}

type BindReqBody struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Data BindReq `json:"data"`
}

type BindReq struct {
	Dn       string `json:"dn"`
	Password string `json:"password"`
}

type HttpSearchReq struct {
	Body SearchReqBody `json:"body"`
	Sign string        `json:"sign"`
}

type SearchReqBody struct {
	From string    `json:"from"`
	To   string    `json:"to"`
	Data SearchReq `json:"data"`
}

type SearchReq struct {
	Attributes []string `json:"attributes"`
	Username   string   `json:"username"`
}
