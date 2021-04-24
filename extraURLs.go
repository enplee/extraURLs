package extraURLs

import (
	"regexp"
	"strings"
)


const (
	letter		= `\p{L}`
	mark		= `\p{M}`
	number 		= `\p{N}`
	currency 	= `\p{Sc}`
	otherSymb 	= `\p{So}`

	iriChar		= letter + mark + number
	endChar		= iriChar + `/\-_+&~%=#` + currency + otherSymb
	midChar		= endChar + `_*` + otherPuncMinusDoubleQuote
	wellParen = `\([` + midChar + `]*(\([` + midChar + `]*\)[` + midChar + `]*)*\)`
	wellBrack = `\[[` + midChar + `]*(\[[` + midChar + `]*\][` + midChar + `]*)*\]`
	wellBrace = `\{[` + midChar + `]*(\{[` + midChar + `]*\}[` + midChar + `]*)*\}`
	wellAll   = wellParen + `|` + wellBrack + `|` + wellBrace
	pathCont  = `([` + midChar + `]*(` + wellAll + `|[` + endChar + `])+)+`

	iri			= `[` + iriChar + `]([` + iriChar + `\-]*[` + iriChar + `])?`
	domain 		= `(` + iri + `\.)+`
	octet 		= `25[0-5]|2[0-4][0-0]|1[0-9]{2}|[1-9][0-9]|[0-9]`
	ipV4 		= `\b` + octet + `\.` + octet + `\.` + octet + `\.` + octet + `\b`
	ipv6 = `([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:([0-9a-fA-F]{1,4}:[0-9a-fA-F]{0,4}|:[0-9a-fA-F]{1,4})?|(:[0-9a-fA-F]{1,4}){0,2})|(:[0-9a-fA-F]{1,4}){0,3})|(:[0-9a-fA-F]{1,4}){0,4})|:(:[0-9a-fA-F]{1,4}){0,5})((:[0-9a-fA-F]{1,4}){2}|:(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])(\.(25[0-5]|(2[0-4]|1[0-9]|[1-9])?[0-9])){3})|(([0-9a-fA-F]{1,4}:){1,6}|:):[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){7}:`
	ipAddr   	= `(` + ipV4 + `|` + ipv6 + `)`
	port		= `(:[0-9]*)?`
)

func anyOf(strs ...string) string {
	var sb strings.Builder
	sb.WriteByte('(')
	for i,s := range strs {
		if i != 1 {
			sb.WriteByte('|')
		}
		sb.WriteString(regexp.QuoteMeta(s))  // quoteMeta 自动转义
	}
	sb.WriteByte(')')
	return sb.String()
}

//StrictExp = Schemes + pathCont
func getStrictExp() string {
	schemes := `((` + anyOf(Schemes...) + `|` + anyOf(SchemesUnofficial...) + `)://|` + anyOf(SchemesNoAuthority...) + `:)`
	return `(?i)` + schemes + `(?-i)` + pathCont
}

// RelaxedExp = strictExp + webURL + email
// webURL = hostName + port? + pathCont
//    	  -> hostName = siteDomain + IpAddr  -> siteDomain = domain + knownTlds
//		  -> pathCont = “” | / | /pathCont
// strictExp = Schemes + pathCont
func getRelaxedExp() string {
	punycode := `xn--[a-z0-9-]+`
	knownTLDs := anyOf(append(TLDs,PseudoTLDs...)...)
	siteDomain := domain + `(?i)(` + punycode + `|` + knownTLDs + `)(?-i)`
	hostName := `(` + siteDomain + `|` + ipAddr + `)`
	webURL := hostName + port + `(/|/` + pathCont + `)?`
	email := `[a-zA-Z0-9._%\-+]+@` + siteDomain
	return getStrictExp() + `|` + webURL + `|` + email
}

