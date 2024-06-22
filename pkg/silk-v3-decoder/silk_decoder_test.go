package silkv3decoder

import "testing"

func TestDecoder(t *testing.T) {
	if err := Decoder("./", "1913fb7d482114ec5ad.aud.silk", "./"); err != nil {
		t.Log(err)
	}
}
