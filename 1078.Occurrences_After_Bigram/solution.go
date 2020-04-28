package main

type WordBuffer struct {
	len int
	buf []byte
}

func newWordBuffer() *WordBuffer {
	return &WordBuffer{
		0,
		make([]byte, 0),
	}
}

func (wb *WordBuffer) eq(s string) bool {
	if wb.len != len(s) {
		return false
	}
	for i := 0; i < wb.len; i++ {
		if wb.buf[i] != s[i] {
			return false
		}
	}
	return true
}

func (wb *WordBuffer) clone(awb *WordBuffer) {
	wb.clean()
	for i := 0; i < awb.len; i++ {
		wb.push(awb.buf[i])
	}
}

func (wb *WordBuffer) push(b byte) {
	if wb.len >= len(wb.buf) {
		wb.buf = append(wb.buf, b)
	} else {
		wb.buf[wb.len] = b
	}
	wb.len++
}

func (wb *WordBuffer) clean() {
	wb.len = 0
}

func (wb *WordBuffer) String() string {
	return string(wb.buf[:wb.len])
}

func findOcurrences(text string, first string, second string) []string {
	_1stWordBuf, _2ndWordBuf, _3rdWordBuf := newWordBuffer(), newWordBuffer(), newWordBuffer()
	ret := make([]string, 0)
	count := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' || i == len(text)-1 {
			count++
			if count == 3 {
				if i == len(text)-1 {
					_3rdWordBuf.push(text[i])
				}
				if _1stWordBuf.eq(first) && _2ndWordBuf.eq(second) {
					ret = append(ret, _3rdWordBuf.String())
					_1stWordBuf.clone(_3rdWordBuf)
					_2ndWordBuf.clean()
					_3rdWordBuf.clean()
					count = 1
				} else {
					_1stWordBuf.clone(_2ndWordBuf)
					_2ndWordBuf.clone(_3rdWordBuf)
					_3rdWordBuf.clean()
					count = 2
				}
			}
			continue
		}
		if count == 0 {
			_1stWordBuf.push(text[i])
		} else if count == 1 {
			_2ndWordBuf.push(text[i])
		} else if count == 2 {
			_3rdWordBuf.push(text[i])
		}
	}

	return ret
}
