package main

import (
	"bytes"
)

// PAYPALISHIRING n=4
// P   I   N
// A  LS  IG
// Y A H R
// P   I
// chunk1: PAY
// chunk2: PAL
// chunk3: ISH
// chunk4: IRI
// chunk5: NG
func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	} else if numRows == 1 {
		return s
	}
	size := len(s)
	if numRows >= size {
		return s
	}
	chunks := size / (numRows - 1)
	if size%(numRows-1) != 0 {
		chunks++
	}
	var buf []byte
	var tailBuf []byte
	strBuilder := bytes.NewBuffer(buf)
	tail := bytes.NewBuffer(tailBuf)
	for i := 0; i < numRows-1; i++ {
		for j := 0; j < chunks; j++ {
			if i == 0 {
				if j%2 == 0 {
					if j*(numRows-1) < size {
						strBuilder.WriteByte(s[j*(numRows-1)])
					}
				} else {
					if j*(numRows-1) < size {
						tail.WriteByte(s[j*(numRows-1)])
					}
				}
			} else {
				if j%2 == 0 {
					if j*(numRows-1)+i < size {
						strBuilder.WriteByte(s[j*(numRows-1)+i])
					}
				} else {
					if (j+1)*(numRows-1)-i < size {
						strBuilder.WriteByte(s[(j+1)*(numRows-1)-i])
					}
				}
			}
		}
	}
	strBuilder.Write(tail.Bytes())

	return strBuilder.String()
}
