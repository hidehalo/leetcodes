package main

func generateTheString(n int) string {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		ret[i] = 'a'
	}
	if n%2 == 0 {
		ret[n-1] = 'b'
	}

	return string(ret)
}
