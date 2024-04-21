package keysgenserviceTest

import (
	"fmt"

	tokengenfuncs "ApiExtention.com/src/services/token_gen_service/funcs"
)

func Gen_Key_Test() {

	var stringTest string = "var"

	tokenString, err := tokengenfuncs.Gen_Token(&stringTest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tokenString)
	}

	if err := tokengenfuncs.Validate_Token(tokenString); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("O token é válido")
	}
}
