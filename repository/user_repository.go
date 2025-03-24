package repository

import (
	"log"
	"payroll_calculator/models"
)

func (r *Repository) ListAllUsers() {
	query := "SELECT * FROM users"

	rows, err := r.rawSql.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	// Es una palabra clave en Go que se utiliza para posponer la ejecución de
	// una función (o método) hasta que la función que la contiene termine su
	// ejecución, ya sea de manera normal o debido a un error. Esto es útil para
	// asegurar que ciertas operaciones de limpieza o liberación de recursos se
	// realicen siempre, incluso si ocurre un error durante la ejecución de
	// la función.
	defer rows.Close()

	// Declarando una lista de estructuras User
	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

}
