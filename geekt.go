package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type family struct {
	id, generation                        int
	name, gender, father, mother, partner string
}

//////////////////////////////////

var db *sql.DB
var err error

///////////////////////////////////////////////////
/* function to find brother of the given person */
//////////////////////////////////////////////////

func find_brother(person_name string) (err error) {

	var person_father string

	err = db.QueryRow("select father from family where name = ?", person_name).Scan(&person_father)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where father = ? and name != ? gender = 'M'", person_father, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_brother string
		err = rows.Scan(&person_brother)

		fmt.Println("Brother of ", person_name, " is ", person_brother)
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find sister of the given person */
//////////////////////////////////////////////////

func find_sister(person_name string) (err error) {

	var person_father string

	err = db.QueryRow("select father from family where name = ?", person_name).Scan(&person_father)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where father = ? and name != ? and gender = 'F'", person_father, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_sister string
		err = rows.Scan(&person_sister)

		fmt.Println("Sister of ", person_name, " is ", person_sister)

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find father of the given person */
//////////////////////////////////////////////////

func find_father(person_name string) (err error) {

	var person_father string

	err = db.QueryRow("select father from family where name = ?", person_name).Scan(&person_father)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Father of ", person_name, " is ", person_father)

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find mother of the given person */
//////////////////////////////////////////////////

func find_mother(person_name string) (err error) {

	var person_mother string

	err = db.QueryRow("select mother from family where name = ?", person_name).Scan(&person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Mother of ", person_name, " is ", person_mother)

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find son of the given person */
//////////////////////////////////////////////////

func find_son(person_name string) (err error) {

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and gender = 'M'", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_son string
		err = rows.Scan(&person_son)

		fmt.Println("Son of ", person_name, " is ", person_son)
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find daughter of the given person */
//////////////////////////////////////////////////

func find_daughter(person_name string) (err error) {

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and gender = 'F'", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_daughter string
		err = rows.Scan(&person_daughter)

		fmt.Println("Daughter of ", person_name, " is ", person_daughter)
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find children of the given person */
//////////////////////////////////////////////////

func find_children(person_name string) (err error) {

	rows, err := db.Query("select name from family where father = ? or mother = ?", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_children string
		err = rows.Scan(&person_children)

		fmt.Println("Children of ", person_name, " is ", person_children)
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find grandfather of the given person */
//////////////////////////////////////////////////

func find_grandfather(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather string

	err = db.QueryRow("select father from family where name = ? or name = ?", person_father, person_mother).Scan(&person_grandfather)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Grandfather of ", person_name, " is ", person_grandfather)

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find grandmother of the given person */
//////////////////////////////////////////////////

func find_grandmother(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather string

	err = db.QueryRow("select mother from family where name = ? or name = ?", person_father, person_mother).Scan(&person_grandfather)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Grandmother of ", person_name, " is ", person_grandfather)

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find grandson of the given person */
//////////////////////////////////////////////////

func find_grandson(person_name string) (err error) {

	row, err := db.Query("select name from family where father = ? or mother = ?", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer row.Close()

	for row.Next() {

		var person_children string
		err = row.Scan(&person_children)

		rows, err := db.Query("select name from family where (father = ? or mother = ?) and gender = 'M'", person_children, person_children)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_grandson string
			err = rows.Scan(&person_grandson)

			fmt.Println("Grandson of ", person_name, " is ", person_grandson)

		}
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find granddaughter of the given person */
//////////////////////////////////////////////////

func find_granddaughter(person_name string) (err error) {

	row, err := db.Query("select name from family where father = ? or mother = ?", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer row.Close()

	for row.Next() {

		var person_children string
		err = row.Scan(&person_children)

		rows, err := db.Query("select name from family where (father = ? or mother = ?) and gender = 'F'", person_children, person_children)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_granddaughter string
			err = rows.Scan(&person_granddaughter)

			fmt.Println("Granddaughter of ", person_name, " is ", person_granddaughter)

		}
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find grandchildren of the given person */
//////////////////////////////////////////////////

func find_grandchildren(person_name string) (err error) {

	row, err := db.Query("select name from family where father = ? or mother = ?", person_name, person_name)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer row.Close()

	for row.Next() {

		var person_children string
		err = row.Scan(&person_children)

		//fmt.Println("Children of ", person_name, " is ", person_children)
		rows, err := db.Query("select name from family where father = ? or mother = ?", person_children, person_children)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_grandchildren string
			err = rows.Scan(&person_grandchildren)

			fmt.Println("Grandchildren of ", person_name, " is ", person_grandchildren)

		}
	}

	return err
}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find cousins of the given person */
//////////////////////////////////////////////////

func find_cousins(person_name string) (err error) {

	var person_father string
	var person_generation int

	err = db.QueryRow("select father, generation from family where name = ?", person_name).Scan(&person_father, &person_generation)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where father != ? and generation = ?", person_father, person_generation)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_cousin string
		err = rows.Scan(&person_cousin)

		fmt.Println("Cousin of ", person_name, " is ", person_cousin)

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find paternal uncle of the given person */
//////////////////////////////////////////////////

func find_paternal_uncle(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather, person_grandmother string

	err = db.QueryRow("select father, mother from family where name = ? or name =?", person_father, person_mother).Scan(&person_grandfather, &person_grandmother)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'M'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_paternal_uncle string
		err = rows.Scan(&person_paternal_uncle)

		fmt.Println("Paternal uncle of ", person_name, " is ", person_paternal_uncle)

	}

	rows1, err := db.Query("select partner from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'F'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows1.Close()

	for rows1.Next() {

		var person_paternal_uncle string
		err = rows1.Scan(&person_paternal_uncle)

		if person_paternal_uncle != "" {
			fmt.Println("Paternal uncle of ", person_name, " is ", person_paternal_uncle)

		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find paternal aunt of the given person */
//////////////////////////////////////////////////

func find_paternal_aunt(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather, person_grandmother string

	err = db.QueryRow("select father, mother from family where name = ? or name = ?", person_father, person_mother).Scan(&person_grandfather, &person_grandmother)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'F'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_paternal_aunt string
		err = rows.Scan(&person_paternal_aunt)

		fmt.Println("Paternal aunt of ", person_name, " is ", person_paternal_aunt)

	}

	rows1, err := db.Query("select partner from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'M'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows1.Close()

	for rows1.Next() {

		var person_paternal_aunt string
		err = rows1.Scan(&person_paternal_aunt)

		if person_paternal_aunt != "" {
			fmt.Println("Paternal aunt of ", person_name, " is ", person_paternal_aunt)

		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find maternal uncle of the given person */
//////////////////////////////////////////////////

func find_maternal_uncle(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather, person_grandmother string

	err = db.QueryRow("select father, mother from family where name = ? or name = ?", person_father, person_mother).Scan(&person_grandfather, &person_grandmother)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'M'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_maternal_uncle string
		err = rows.Scan(&person_maternal_uncle)

		fmt.Println("Paternal uncle of ", person_name, " is ", person_maternal_uncle)

	}

	rows1, err := db.Query("select partner from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'F'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows1.Close()

	for rows1.Next() {

		var person_maternal_uncle string
		err = rows1.Scan(&person_maternal_uncle)

		if person_maternal_uncle != "" {
			fmt.Println("Maternal uncle of ", person_name, " is ", person_maternal_uncle)

		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find maternal aunt of the given person */
//////////////////////////////////////////////////

func find_maternal_aunt(person_name string) (err error) {

	var person_father, person_mother string

	err = db.QueryRow("select father, mother from family where name = ?", person_name).Scan(&person_father, &person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	var person_grandfather, person_grandmother string

	err = db.QueryRow("select father, mother from family where name = ? or name = ?", person_father, person_mother).Scan(&person_grandfather, &person_grandmother)

	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query("select name from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'F'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()

	for rows.Next() {

		var person_maternal_aunt string
		err = rows.Scan(&person_maternal_aunt)

		fmt.Println("Maternal aunt of ", person_name, " is ", person_maternal_aunt)

	}

	rows1, err := db.Query("select partner from family where (father = ? or mother = ?) and (name != ? and name != ?) and gender = 'M'", person_grandfather, person_grandmother, person_father, person_mother)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows1.Close()

	for rows1.Next() {

		var person_maternal_aunt string
		err = rows1.Scan(&person_maternal_aunt)

		if person_maternal_aunt != "" {
			fmt.Println("Maternal aunt of ", person_name, " is ", person_maternal_aunt)

		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find brother in law of the given person */
//////////////////////////////////////////////////

func find_brother_law(person_name string) (err error) {

	var person_partner, person_father string

	err = db.QueryRow("select name from family where partner = ?", person_name).Scan(&person_partner)

	if person_partner != "" {
		err = db.QueryRow("select father from family where name = ?", person_partner).Scan(&person_father)

		rows, err := db.Query("select name from family where father = ? and name != ? and gender = 'M'", person_father, person_partner)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_brother_law string
			err = rows.Scan(&person_brother_law)
			if person_brother_law != "" {
				fmt.Println("Brother in law of ", person_name, " is ", person_brother_law)
			}
		}

	} else {
		err = db.QueryRow("select father from family where name = ?", person_name).Scan(&person_father)

		rows, err := db.Query("select partner from family where father = ? and name != ? and gender = 'F'", person_father, person_name)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_brother_law string
			err = rows.Scan(&person_brother_law)
			if person_brother_law != "" {
				fmt.Println("Brother in law of ", person_name, " is ", person_brother_law)
			}
		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* function to find sister in law of the given person */
//////////////////////////////////////////////////

func find_sister_law(person_name string) (err error) {

	var person_partner, person_father string

	err = db.QueryRow("select name from family where partner = ?", person_name).Scan(&person_partner)

	if person_partner != "" {
		err = db.QueryRow("select father from family where name = ?", person_partner).Scan(&person_father)

		rows, err := db.Query("select name from family where father = ? and name != ? and gender = 'F'", person_father, person_partner)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_sister_law string
			err = rows.Scan(&person_sister_law)
			if person_sister_law != "" {
				fmt.Println("Sister in law of ", person_name, " is ", person_sister_law)
			}
		}

	} else {
		err = db.QueryRow("select father from family where name = ?", person_name).Scan(&person_father)

		rows, err := db.Query("select partner from family where father = ? and name != ? and gender = 'M'", person_father, person_name)

		if err != nil {
			fmt.Println(err.Error())
		}

		defer rows.Close()

		for rows.Next() {

			var person_sister_law string
			err = rows.Scan(&person_sister_law)
			if person_sister_law != "" {
				fmt.Println("Sister in law of ", person_name, " is ", person_sister_law)
			}
		}

	}

	return err

}

///////////////////////////////////////////////////

///////////////////////////////////////////////////
/* this is whre the main function starts. */
//////////////////////////////////////////////////

func main() {
	db, err = sql.Open("mysql", "test123:Keshav@14@tcp(127.0.0.1:3306)/test")
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	/* Taking user input */
	//////////////////////////////////////////////////

	scanner := bufio.NewScanner(os.Stdin)
	var person_name, relation string

	fmt.Println("\n Enter name of the person (first letter should be capital): ")
	scanner.Scan()
	person_name = scanner.Text()

	fmt.Println("\n Enter the relation (every letter should be small): ")
	scanner.Scan()
	relation = scanner.Text()

	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	/* Condition starts here */
	//////////////////////////////////////////////////

	if relation == "brother" {
		find_brother(person_name)

	} else if relation == "sister" {
		find_sister(person_name)

	} else if relation == "father" {
		find_father(person_name)

	} else if relation == "mother" {
		find_mother(person_name)

	} else if relation == "son" {
		find_son(person_name)

	} else if relation == "daughter" {
		find_daughter(person_name)

	} else if relation == "children" {
		find_children(person_name)

	} else if relation == "grandfather" {
		find_grandfather(person_name)

	} else if relation == "grandmother" {
		find_grandmother(person_name)

	} else if relation == "grandson" {
		find_grandson(person_name)

	} else if relation == "granddaughter" {
		find_granddaughter(person_name)

	} else if relation == "grandchildren" {
		find_grandchildren(person_name)

	} else if relation == "cousins" {
		find_cousins(person_name)

	} else if relation == "paternal uncle" {
		find_paternal_uncle(person_name)

	} else if relation == "paternal aunt" {
		find_paternal_aunt(person_name)

	} else if relation == "maternal uncle" {
		find_maternal_uncle(person_name)

	} else if relation == "maternal aunt" {
		find_maternal_aunt(person_name)

	} else if relation == "brother in law" {
		find_brother_law(person_name)

	} else if relation == "sister in law" {
		find_sister_law(person_name)

	} else {
		fmt.Println("not found")

	}

	///////////////////////////////////////////////////

	///////////////////////////////////////////////////
	/* End of the program */
	//////////////////////////////////////////////////

}
