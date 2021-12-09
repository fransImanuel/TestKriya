package user

import (
	"fmt"
	"math"

	"github.com/jmoiron/sqlx"
)

func getListUser(db *sqlx.DB, param GetUserListParam) ([]unMarshalUserData, error) {
	var (
		datas []unMarshalUserData
		data  UserData
	)
	query := `select "id","data" from users limit 5 offset $1`
	offset := math.Abs(float64(param.Page-1)) * 5
	result, err := db.Queryx(query, offset)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		err := result.StructScan(&data)
		if err != nil {
			return nil, err
		}
		json := data.UnMarshal()
		fmt.Println(json)
		datas = append(datas, unMarshalUserData{
			Username: json.Username,
			Email:    json.Email,
			Status:   json.Is_Active,
		})
	}

	return datas, nil
}

func getUser(db *sqlx.DB, param GetUserParam) (*unMarshalUserData, error) {
	var (
		data     UserData
		roleData UserData
	)
	query := `select u."data",u."role_id" 
		from users u
		join roles r on r.id  = u.role_id 
		where u.id = $1`
	user, err := db.Queryx(query, param.Id)
	if err != nil {
		return nil, err
	}

	for user.Next() {
		err := user.StructScan(&data)
		if err != nil {
			return nil, err
		}

		query := `select r."data"
			from roles r
			where id = $1`
		role, err := db.Queryx(query, data.Role_ID)
		if err != nil {
			return nil, err
		}
		for role.Next() {
			err = role.StructScan(&roleData)
			if err != nil {
				return nil, err
			}
		}

	}

	roleJson := roleData.UnMarshal()
	json := data.UnMarshal()

	return &unMarshalUserData{
		Id:        data.Id,
		Username:  json.Username,
		Email:     json.Email,
		Role_Name: roleJson.Role_Name,
	}, nil
}

func addUser(db *sqlx.DB, id, data, role string) error {
	query := `insert into users(id,data,role_id) 
		values($1, $2, $3)	
	`
	result, err := db.Exec(query, id, data, role)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Insert New User")
}

func UpdateUser(db *sqlx.DB, data, id string) error {
	query := `UPDATE users 
		SET data = $1
		WHERE id = $2;
	`
	result, err := db.Exec(query, data, id)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Update New User")
}

func DeleteUser(db *sqlx.DB, id string) error {
	query := `DELETE FROM users WHERE id = $1;
	`
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	if affected, _ := result.RowsAffected(); affected > 0 {
		return nil
	}
	return fmt.Errorf("error when Delete User")
}
