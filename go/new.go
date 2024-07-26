package main

type Param map[string]interface{}

type Show struct {
	Param Param
}

func main() {
	s := Show{}
	s.Param["RMB"] = 10000

}
