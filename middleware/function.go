package middleware

import (
	"encoding/json"
	"fmt"
	"kriya_Test/utilities/db"
)

func checkRoleByID(id string) bool {
	var (
		role Role
		data DataJSONB
	)
	db := db.Connect()
	defer db.Close()

	query := `select "data"  from roles
			where id = $1`

	err := db.Get(&role, query, id)
	if err != nil {
		fmt.Println(err)
		return false
	}

	json.Unmarshal([]byte(*role.Data), &data)

	fmt.Println(*data.Role_name == "admin")
	return *data.Role_name == "admin"

}
