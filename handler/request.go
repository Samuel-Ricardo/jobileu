package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("params %s of type %s is required", name, typ)
}
