package audio

import (
	"testing"
	"time"
)

func TestSilkToPcm(t *testing.T) {
	now := time.Now()
	if err := SilkToPcm("data", "1913fb7d482114ec5ad.aud.silk"); err != nil {
		t.Error(err)
	}
	t.Log(time.Since(now))
}

func TestDecoder(t *testing.T) {
	now := time.Now()
	if err := Decoder("data", "1913fb7d482114ec5ad.aud.silk"); err != nil {
		t.Error(err)
	}
	t.Log(time.Since(now))
}
