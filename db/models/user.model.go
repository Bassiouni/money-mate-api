package models

import (
	"github.com/Bassiouni/money-mate-api/contracts"
	"github.com/Bassiouni/money-mate-api/db"
	"log"
)

func All() ([]contracts.User, error) {

	rows, err := db.Instance.Query("SELECT * FROM users")

	if err != nil {
		log.Println("Couldn't fetch all users")
		return nil, err
	}

	defer rows.Close()

	var dest = []contracts.User{}

	for rows.Next() {
		var r contracts.User

		err = rows.Scan(&r.ID, &r.FirstName, &r.LastName, &r.Email, &r.Password, &r.RoleType)

		if err != nil {
			log.Fatalf("Scan: %v\n", err)
		}

		dest = append(dest, r)
	}

	return dest, nil
}

func FindByEmail(email string) (contracts.User, error) {
	var dest contracts.User

	err := db.Instance.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&dest.ID, &dest.FirstName, &dest.LastName, &dest.Email, &dest.Password, &dest.RoleType)

	if err != nil {
		log.Println("Couldn't fetch the user")
		return contracts.User{}, err
	}

	return dest, nil
}

func FindByID(id int) (contracts.User, error) {
	var dest contracts.User

	err := db.Instance.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&dest.ID, &dest.FirstName, &dest.LastName, &dest.Email, &dest.Password, &dest.RoleType)

	if err != nil {
		log.Println("Couldn't fetch the user")
		return contracts.User{}, err
	}

	return dest, nil
}

func CreateNewUser(u *contracts.User) (contracts.User, error) {
	var r contracts.User

	err := db.Instance.QueryRow(
		"INSERT INTO users (firstname, lastname, email, password, role_type) VALUES ($1, $2, $3, $4, $5) RETURNING *",
		u.FirstName, u.LastName, u.Email, u.Password, u.RoleType,
	).Scan(&r.ID, &r.FirstName, &r.LastName, &r.Email, &r.Password, &r.RoleType)

	if err != nil {
		log.Printf("Couldn't create new user: %v\n", err)
		return contracts.User{}, err
	}

	return r, nil
}
