package errors

import (
	"errors"
	"fmt"
	"os"
)

// Catch imprime el error
func Catch(err error, _exit ...bool) {
	if err != nil {
		fmt.Println(err.Error())
		if len(_exit) > 0 {
			os.Exit(0)
		}
	}
}

//New crea  un error
func New(_err string) error {
	return errors.New(_err)
}
