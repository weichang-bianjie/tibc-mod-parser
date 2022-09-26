package tibc_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/tibc-mod-parser/codec"
	"github.com/kaifei-bianjie/tibc-mod-parser/modules/tibc"
)

type MsgClient struct {
	Tibc Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Tibc: tibc.NewClient(),
	}
}
