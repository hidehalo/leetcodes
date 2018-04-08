package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	ret := make([]string, 0, 100)
	d := make(map[string]int)
	for _, cp := range cpdomains {
		tmp := strings.Split(cp, " ")
		cnt, domain := tmp[0], tmp[1]
		sub := strings.Split(domain, ".")
		crt, _ := strconv.Atoi(cnt)
		if _, ok := d[domain]; ok {
			d[domain] += crt
		} else {
			d[domain] = crt
		}
		size := len(sub)
		for i := size - 1; i > 0; i-- {
			subd := strings.Join(sub[size-i:size], ".")
			if _, ok := d[subd]; ok {
				d[subd] += crt
			} else {
				d[subd] = crt
			}
		}
	}
	for d, c := range d {
		ret = append(ret, fmt.Sprintf("%d %s", c, d))
	}

	return ret
}

func main() {
	param := []string{"900 google.mail.com", "50 www.yahoo.com", "1 intel.mail.com", "5 wiki.org", "8 google.wiki.org", "900 google.mail.com"}
	str := subdomainVisits(param)
	fmt.Println(str)
}
