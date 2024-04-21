package db_test

import dbactionsfuncs "ApiExtention.com/src/services/db_service/funcs/actions"

// Exemple Code of MongoDB Connection
func Test_Connect() {

	// // Criação da conexão
	// client, err := db_funcs.Gen_ServerAPIClient()
	// if err != nil {
	// 	fmt.Printf("Error (Create a Connection to the DB): %s", err)
	// }

	// // Teste de conexão com o banco
	// act_ping := db_funcs.Ping_Connection_DataBase(client)
	// if act_ping != nil {
	// 	fmt.Printf("Error (Ping DB): %s", act_ping)
	// } else {
	// 	fmt.Println("Active connection to the DB")
	// }

	// // Fechamento da conexão
	// defer func() {
	// 	if err := db_funcs.Close_DataBase_Connection(client); err != nil {
	// 		fmt.Printf("Error (Close Connection): %s", err)
	// 	} else {
	// 		fmt.Println("Connection closed to the DB")
	// 	}
	// }()

	dbactionsfuncs.Query_Token("aaa")
}
