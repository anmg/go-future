package client

func Sum(ops...int) int {
	ret := 0
	for _,op := range ops {
		ret += op
	}
	return ret
}
