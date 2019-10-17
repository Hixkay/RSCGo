package db

import (
	"bitbucket.org/zlacki/rscgo/pkg/server/config"
	"bitbucket.org/zlacki/rscgo/pkg/server/log"

	// Necessary for sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

//ObjectDefinition This represents a single definition for a single object in the game.
type ObjectDefinition struct {
	ID            int
	Name          string
	Commands      []string
	Description   string
	Type          int
	Width, Height int
	Length        int
}

//ItemDefinition This represents a single definition for a single item in the game.
type ItemDefinition struct {
	ID          int
	Name        string
	Description string
	Command     string
	BasePrice   int
	Stackable   bool
	Quest       bool
	Members     bool
}

//BoundaryDefinition This represents a single definition for a single boundary object in the game.
type BoundaryDefinition struct {
	ID          int
	Name        string
	Commands    []string
	Description string
}

//Objects This holds the defining characteristics for all of the game's scene objects, ordered by ID.
var Objects []ObjectDefinition

//Items This holds the defining characteristics for all of the game's scene objects, ordered by ID.
var Items []ItemDefinition

//Boundarys This holds the defining characteristics for all of the game's boundary scene objects, ordered by ID.
var Boundarys []BoundaryDefinition

//LoadObjectDefinitions Loads game object data into memory for quick access.
func LoadObjectDefinitions() error {
	database := Open(config.WorldDB())
	defer database.Close()
	rows, err := database.Query("SELECT id, name, description, command_one, command_two, type, width, height, ground_item_var FROM `game_objects`")
	defer rows.Close()
	if err != nil {
		log.Error.Println("Couldn't load SQLite3 database:", err)
		return err
	}
	for rows.Next() {
		nextDef := ObjectDefinition{Commands: make([]string, 2)}
		rows.Scan(&nextDef.ID, &nextDef.Name, &nextDef.Description, &nextDef.Commands[0], &nextDef.Commands[1], &nextDef.Type, &nextDef.Width, &nextDef.Height, &nextDef.Length)
		Objects = append(Objects, nextDef)
	}
	return nil
}

//LoadBoundaryDefinitions Loads game boundary object data into memory for quick access.
func LoadBoundaryDefinitions() {
	database := Open(config.WorldDB())
	defer database.Close()
	// TODO: Seem to be missing a lot of door data.
	rows, err := database.Query("SELECT id, name, description, command_one, command_two FROM `doors`")
	defer rows.Close()
	if err != nil {
		log.Error.Println("Couldn't load SQLite3 database:", err)
		return
	}
	for rows.Next() {
		nextDef := BoundaryDefinition{Commands: make([]string, 2)}
		rows.Scan(&nextDef.ID, &nextDef.Name, &nextDef.Description, &nextDef.Commands[0], &nextDef.Commands[1])
		Boundarys = append(Boundarys, nextDef)
	}
}

//LoadItemDefinitions Loads game item data into memory for quick access.
func LoadItemDefinitions() {
	database := Open(config.WorldDB())
	defer database.Close()
	rows, err := database.Query("SELECT id, name, description, command, base_price, stackable, special, members FROM `items` ORDER BY id")
	defer rows.Close()
	if err != nil {
		log.Error.Println("Couldn't load SQLite3 database:", err)
		return
	}
	for rows.Next() {
		nextDef := ItemDefinition{}
		rows.Scan(&nextDef.ID, &nextDef.Name, &nextDef.Description, &nextDef.Command, &nextDef.BasePrice, &nextDef.Stackable, &nextDef.Quest, &nextDef.Members)
		Items = append(Items, nextDef)
	}
}