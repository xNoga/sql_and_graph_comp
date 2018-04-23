package postgres

import "database/sql"

type person struct {
	id       int64
	birthday string
	name     string
	pageRank float64
	job      string
}

func getEndorse1st(db *sql.DB) (res []*person, err error) {
	rows, err := db.Query(`
		SELECT name,id FROM person
		WHERE id IN (
			SELECT id2 FROM relation
			WHERE relation.id1 IN (
				SELECT id FROM person LIMIT 1
			)
		);`)

	check(err)
	defer rows.Close()

	res, err = getDataFromRows(rows)
	return
}

func getEndorse2nd(db *sql.DB) (res []*person, err error) {
	rows, err := db.Query(`
		SELECT name,id FROM person
		WHERE id IN (
			SELECT id2 FROM relation
			WHERE relation.id1 IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT id2 FROM relation
					WHERE relation.id1 IN (
						SELECT id FROM person LIMIT 1
					)
				)
			)
		);
	`)
	check(err)
	defer rows.Close()

	res, err = getDataFromRows(rows)
	return
}

func getEndorse3rd(db *sql.DB) (res []*person, err error) {
	rows, err := db.Query(`
		SELECT name,id FROM person
		WHERE id IN (
			SELECT id2 FROM relation
			WHERE relation.id1 IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT id2 FROM relation
					WHERE relation.id1 IN (
						SELECT id FROM person
						WHERE id IN (
							SELECT id2 FROM relation
							WHERE relation.id1 IN (
								SELECT id FROM person LIMIT 1
							)
						)
					)
				)
			)
		);
	`)
	check(err)
	defer rows.Close()

	res, err = getDataFromRows(rows)
	return
}

func getEndorse4th(db *sql.DB) (res []*person, err error) {
	rows, err := db.Query(`
		SELECT name,id FROM person
		WHERE id IN (
			SELECT id2 FROM relation
			WHERE relation.id1 IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT id2 FROM relation
					WHERE relation.id1 IN (
						SELECT id FROM person
						WHERE id IN (
							SELECT id2 FROM relation
							WHERE relation.id1 IN (
								SELECT id FROM person
								WHERE id IN (
									SELECT id2 FROM relation
									WHERE relation.id1 IN (
										SELECT id FROM person LIMIT 1
									)
								)
							)
						)
					)
				)
			)
		);
	`)
	check(err)
	defer rows.Close()

	res, err = getDataFromRows(rows)
	return
}

func getEndorse5th(db *sql.DB) (res []*person, err error) {
	rows, err := db.Query(`
		SELECT name,id FROM person
		WHERE id IN (
			SELECT id2 FROM relation
			WHERE relation.id1 IN (
				SELECT id FROM person
				WHERE id IN (
					SELECT id2 FROM relation
					WHERE relation.id1 IN (
						SELECT id FROM person
						WHERE id IN (
							SELECT id2 FROM relation
							WHERE relation.id1 IN (
								SELECT id FROM person
								WHERE id IN (
									SELECT id2 FROM relation
									WHERE relation.id1 IN (
										SELECT id FROM person
										WHERE id IN (
											SELECT id2 FROM relation
											WHERE relation.id1 IN (
												SELECT id FROM person LIMIT 1
											)
										)
									)
								)
							)
						)
					)
				)
			)
		);
	`)
	check(err)
	defer rows.Close()

	res, err = getDataFromRows(rows)
	return
}

func getDataFromRows(rows *sql.Rows) (res []*person, err error) {
	for rows.Next() {
		p := &person{}
		err := rows.Scan(&p.name, &p.id)
		if err != nil {
			panic(err)
		}
		res = append(res, p)
	}

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
