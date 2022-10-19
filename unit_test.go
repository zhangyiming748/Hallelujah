package Hallelujah

import (
	"testing"
)

func TestUnit(t *testing.T) {
	Hallelujah()
}

func TestApifox(t *testing.T) {
	var img Img

	img = Hallelujah()
	t.Logf("%+v\n", img)
}
