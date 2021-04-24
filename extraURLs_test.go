package extraURLs

import (
	"regexp"
	"testing"
)



func TestRexg(t *testing.T)  {
	re := regexp.MustCompile(domain)
	s := re.FindString("w1221-s-s-.3.s.s")
	t.Log(s)
}