package kavenegar

import (
	"net/url"
)

//MakeTTS ...
func (call *CallService) MakeTTS(sender string, receptor string, message string, params *CallParam) (Message, error) {
	v := structToURLValues(params)
	v.Set("sender", sender)
	v.Set("receptor", receptor)
	v.Set("message", message)
	return call.CreateMakeTTS(v)
}

//CreateMakeTTS ...
func (call *CallService) CreateMakeTTS(v url.Values) (Message, error) {
	u := call.client.EndPoint("call", "maketts")
	m := new(MessageResult)
	err := call.client.Execute(u.String(), v, m)
	if	m.Entries==nil{
		return Message{},err
	}
	return m.Entries[0], err
}
