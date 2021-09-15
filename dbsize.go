package eosrpc

import "github.com/TheBoringDude/go-urljoin"

const DBSIZE_API = "db_size"

func (d *DBSizeAPI) Get() (DBSizeGetResponse, error) {
	var r = DBSizeGetResponse{}
	err := request(d.Client, urljoin.UrlJoin(d.ApiUrl, "get"), nil, &r)

	return r, err
}
