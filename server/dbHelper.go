package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

// TODO: set this to the actual default ballot
var tempCurrBal string

func setupDB() (*bolt.DB, error) {
	var err error
	db, err = bolt.Open("database/test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("BALLOT"))
		if err != nil {
			return fmt.Errorf("could not create ballot bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("PROMPT"))
		if err != nil {
			return fmt.Errorf("could not create prompt bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}

	fmt.Println("DB Setup Done")
	return db, nil
}

func dbSetCurrrentBallotByBallot(db *bolt.DB, bal ballotStruct) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Put([]byte("CURRENT_BALLOT"), []byte(bal.ID))
		if err != nil {
			return fmt.Errorf("could not set current ballot: %v", err)
		}
		return nil
	})

	fmt.Println("Set current ballot to ", bal.ID)
	return err
}

func dbSetCurrrentBallotByID(db *bolt.DB, id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("DB")).Put([]byte("CURRENT_BALLOT"), []byte(id))
		if err != nil {
			return fmt.Errorf("could not set current ballot: %v", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	fmt.Println("Set current ballot to ", id)
	return nil
}

func dbAddBallot(db *bolt.DB, bal ballotStruct) error {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(99999)
	fmt.Println("Adding ballot with id: ", id)

	// while populating set some random ballot as the current ballot
	tempCurrBal = strconv.Itoa(id)
	bal.ID = strconv.Itoa(id)
	return dbUpdateBallot(db, bal, []byte(strconv.Itoa(id)))
}

func dbUpdateBallot(db *bolt.DB, bal ballotStruct, id []byte) error {
	balBytes, err := json.Marshal(bal)
	if err != nil {
		return fmt.Errorf("could not marshal ballot json: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Put(id, balBytes)
		if err != nil {
			return fmt.Errorf("could not set ballot: %v", err)
		}
		return nil
	})
	return err
}

func dbGetBallotByID(db *bolt.DB, id string) (ballotStruct, error) {
	var bal ballotStruct
	err := db.View(func(tx *bolt.Tx) error {
		balBytes := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(id)))
		if err := json.Unmarshal(balBytes, &bal); err != nil {
			return fmt.Errorf("could not fetch ballot from db: %v", err)
		}
		return nil
	})
	if err != nil {
		return bal, fmt.Errorf("could not print ballots: %v", err)
	}
	return bal, nil
}

func dbGetCurrentBallot(db *bolt.DB) (ballotStruct, error) {
	var bal ballotStruct
	var id []byte
	err := db.View(func(tx *bolt.Tx) error {
		id = tx.Bucket([]byte("DB")).Get([]byte([]byte("CURRENT_BALLOT")))
		if id == nil {
			id = []byte(strconv.Itoa(16714))
		}
		fmt.Println("Got current id: ", id)

		balBytes := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(id)))
		if err := json.Unmarshal(balBytes, &bal); err != nil {
			return fmt.Errorf("could not convert returned ballot to type ballot: %v\n%v", err, bal)
		}
		return nil
	})
	if err != nil {
		return bal, fmt.Errorf("could not fetch current ballot from db: %v", err)
	}
	return bal, nil
}

func dbPrintBallots(db *bolt.DB) error {
	fmt.Println("Printing db now")
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not print ballots: %v", err)
	}
	fmt.Println("That was the db")
	return nil
}

func dbGetAllBallots(db *bolt.DB) ([]ballotStruct, error) {
	allBallots := []ballotStruct{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT"))
		b.ForEach(func(k, v []byte) error {
			var bal ballotStruct
			if err := json.Unmarshal(v, &bal); err != nil {
				return fmt.Errorf("could not fetch ballot from db: %v", err)
			}
			allBallots = append(allBallots, bal)
			return nil
		})
		return nil
	})
	if err != nil {
		return allBallots, fmt.Errorf("could not print ballots: %v", err)
	}
	return allBallots, nil
}

func populateDB(db *bolt.DB) {
	for _, bal := range ballots {
		dbAddBallot(db, bal)
	}
	dbSetCurrrentBallotByID(db, tempCurrBal)
	fmt.Println("Set current ballot to id: ", []byte(tempCurrBal))
}

func dbVote(db *bolt.DB, v voteStruct) (ballotStruct, error) {
	var bal ballotStruct
	err := db.View(func(tx *bolt.Tx) error {
		balBytes := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(v.BallotID)))

		if err := json.Unmarshal(balBytes, &bal); err != nil {
			return fmt.Errorf("could not fetch ballot from db: %v", err)
		}

		newPrompts := []promptStruct{}

		for _, p := range bal.Prompts {
			if p.ID == v.PromptID {
				p.Votes += 1
			}
			newPrompts = append(newPrompts, p)
		}
		bal.Prompts = newPrompts
		return nil
	})
	if err != nil {
		return bal, fmt.Errorf("could not update vote: %v", err)
	}
	dbUpdateBallot(db, bal, []byte(v.BallotID))
	fmt.Println(bal)
	return bal, nil
}

func dbPrintById(db *bolt.DB, id int) {
	err := db.View(func(tx *bolt.Tx) error {
		bal := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(strconv.Itoa(id))))
		fmt.Printf("Ballot: %s\n", bal)
		fmt.Println(reflect.TypeOf(bal))
		return nil
	})
	if err != nil {
		fmt.Printf("could not print ballots: %v", err)
	}
}
