package js

import (
	"bytes"
	"io"
	"io/ioutil"

	babel "github.com/jvatic/goja-babel"
)

var es2015TranformCfg = map[string]interface{}{
	"presets": []string{"es2015"},
	"plugins": []string{
		"transform-es2015-block-scoping",
	},
}

func Transform(src io.Reader) (io.Reader, error) {
	res, err := babel.Transform(src, es2015TranformCfg)
	if err != nil {
		return nil, err
	}
	code, err := ioutil.ReadAll(res)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(code), nil
}
