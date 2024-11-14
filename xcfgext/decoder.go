package xcfgext

import (
	"github.com/BurntSushi/toml"
	"github.com/xanygo/anygo/xcfg"
	"github.com/xanygo/anygo/xcodec"
	"gopkg.in/yaml.v3"
)

type decoders struct {
	Decoder xcodec.Decoder
	Ext     string
}

// defaultParsers 所有默认的 parser，
// 当传入配置文件名不包含后置的时候，会使用此顺序依次查找
var parsers = []decoders{
	{Ext: ".yml", Decoder: xcodec.DecodeFunc(yaml.Unmarshal)},
	{Ext: ".toml", Decoder: xcodec.DecodeFunc(toml.Unmarshal)},
	{Ext: ".yaml", Decoder: xcodec.DecodeFunc(yaml.Unmarshal)},
}

func init() {
	for _, p := range parsers {
		_ = xcfg.WithDecoder(p.Ext, p.Decoder)
	}
}

func Init() {
}
