package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Create a new cluster configuration
	cluster := gocql.NewCluster("localhost:9042") // Replace with your Cassandra cluster address
	cluster.Keyspace = ""                         // Replace with your keyspace name
	cluster.Consistency = gocql.Quorum            // Set the desired consistency level

	// Create a session to connect to Cassandra
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	// // Create a keyspace and table (if they don't exist)
	// if err := session.Query(`
	//       CREATE KEYSPACE IF NOT EXISTS mykeyspace
	//       WITH replication = {
	//           'class': 'SimpleStrategy',
	//           'replication_factor': 1
	//       }`).Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	if err := session.Query(`
        CREATE TABLE IF NOT EXISTS mytable (
            id UUID PRIMARY KEY,
            name TEXT
        )`).Exec(); err != nil {
		log.Fatal(err)
	}

	// Insert data into the table
	id := gocql.TimeUUID()
	name := "John Doe"

	if err := session.Query(`
        INSERT INTO mytable (id, name)
        VALUES (?, ?)`, id, name).Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted data: ID=%s, Name=%s\n", id.String(), name)

	// Retrieve data from the table
	var retrievedID gocql.UUID
	var retrievedName string

	if err := session.Query(`
        SELECT id, name
        FROM mytable
        WHERE id = ?`, id).Scan(&retrievedID, &retrievedName); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved data: ID=%s, Name=%s\n", retrievedID.String(), retrievedName)
}
