package devafatekresult

import "encoding/json"

//Result Type
type ResultType struct {
	Result string
	Retval interface{}
}

//ToByte
func (res ResultType) ToByte() []byte {
	jData, _ := json.Marshal(res)
	return jData
}

//ToString
func (res ResultType) ToString() string {
	return string(res.ToByte())
}
