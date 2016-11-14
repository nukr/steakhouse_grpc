package main

import (
	"fmt"
	"log"

	pb "github.com/nukr/steakhouse_grpc/pb"
)

func queryCuisineByID(id string) (*pb.Cuisine, error) {
	fmt.Println(id)
	q := db.QueryRow(`
  SELECT
    id,
    title,
    description,
    price,
    is_deleted
  FROM cuisine
  WHERE id = $1
  `, id)
	cuisine := &pb.Cuisine{}
	err := q.Scan(
		&cuisine.Id,
		&cuisine.Title,
		&cuisine.Description,
		&cuisine.Price,
		&cuisine.IsDeleted)
	if err != nil {
		log.Printf("scan error %v", err)
	}
	return cuisine, nil
}

func eachCuisines(fn func(*pb.Cuisine)) {
	rows, err := db.Query(`
  SELECT id, title, description, price, is_deleted
  FROM cuisine
  `)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		cuisine := &pb.Cuisine{}
		err := rows.Scan(
			&cuisine.Id,
			&cuisine.Title,
			&cuisine.Description,
			&cuisine.Price,
			&cuisine.IsDeleted)
		if err != nil {
			log.Println(err)
		}
		fn(cuisine)
	}
}

func createCuisine(cuisine *pb.Cuisine) error {
	_, err := db.Exec(`
  INSERT INTO cuisine (id, title, description, price)
  VALUES ($1, $2, $3, $4)
  `, cuisine.Id, cuisine.Title, cuisine.Description, cuisine.Price)
	if err != nil {
		return err
	}
	return nil
}

func deleteCuisine(id string) error {
	_, err := db.Exec(`
  DELETE FROM cuisine
  WHERE id = $1
  `, id)
	if err != nil {
		return err
	}
	return nil
}

func updateCuisine(cuisine *pb.Cuisine) error {
	_, err := db.Exec(`
  UPDATE cuisine
  SET title = $1, description = $2, price = $3
  WHERE id = $4
  `, cuisine.Title, cuisine.Description, cuisine.Price, cuisine.Id)
	return err
}
