package e

func Generate(s string, d interface{}) e{
	result := elist[s]
	result.Data = d
	return result
}

