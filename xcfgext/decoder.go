package xcfgext

import (
	"github.com/BurntSushi/toml"
	"github.com/xanygo/anygo/xcfg"
	"github.com/xanygo/anygo/xenc"
	"github.com/xanygo/anygo/xenc/xcodec"
	"gopkg.in/yaml.v3"
)

type decoders struct {
	Decoder xcodec.Unmarshaler
	Ext     string
}

// defaultParsers 所有默认的 parser，
// 当传入配置文件名不包含后置的时候，会使用此顺序依次查找
var parsers = []decoders{
	{Ext: ".yml", Decoder: xenc.UnmarshalFunc(yaml.Unmarshal)},
	{Ext: ".toml", Decoder: xenc.UnmarshalFunc(toml.Unmarshal)},
	{Ext: ".yaml", Decoder: xenc.UnmarshalFunc(yaml.Unmarshal)},
}

func init() {
	for _, p := range parsers {
		_ = xcfg.WithDecoder(p.Ext, p.Decoder)
	}
}

func Init() {
}
