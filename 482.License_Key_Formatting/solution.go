package main

import (
	"bytes"
	"log"
)

func licenseKeyFormatting(S string, K int) string {
	origin := bytes.NewBuffer(make([]byte, 0, len(S)))
	ret := bytes.NewBuffer(make([]byte, 0, len(S)*2))
	for i := 0; i < len(S); i++ {
		if S[i] == '-' {
			continue
		}
		origin.WriteByte(toUpper(S[i]))
	}
	// fmt.Println(origin.String())
	firstGroupLen := origin.Len() % K
	var groupNum int
	if firstGroupLen == 0 {
		groupNum = origin.Len() / K
		firstGroupLen = K
	} else {
		groupNum = origin.Len()/K + 1
	}
	// fmt.Println("Group Num=", groupNum)
	for i := 0; i < groupNum; i++ {
		var needCopy int
		if i == 0 {
			needCopy = firstGroupLen
		} else {
			needCopy = K
			ret.WriteByte('-')
		}
		// fmt.Println("Group #", i+1, "Need Copy=", needCopy)
		for j := 0; j < needCopy; j++ {
			b, err := origin.ReadByte()
			if err != nil {
				log.Fatalln(err)
			}
			err = ret.WriteByte(b)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	return ret.String()
}

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - byte(32)
	}
	return b
}
