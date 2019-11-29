// Copyright 2019 Reaction Engineering International. All rights reserved.
// Use of this source code is governed by the MIT license in the file LICENSE.txt.

package passwords

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/reaction-eng/restlib/configuration"
	"github.com/reaction-eng/restlib/email"
)

/**
Define a struct for Repo for use with users
*/
type ResetRepoSql struct {
	//Hold on to the sql databased
	db *sql.DB

	//Also store the table name
	tableName string

	//We need the emailer
	emailer               email.Emailer
	resetEmailConfig      PasswordResetConfig
	activationEmailConfig PasswordResetConfig

	//Store the required statements to reduce comput time
	addRequestStatement *sql.Stmt
	getRequestStatement *sql.Stmt
	rmRequestStatement  *sql.Stmt
}

/**
Store the type of token
*/
type tokenType int

const (
	activation tokenType = 1
	reset      tokenType = 2
)

//Provide a method to make a new UserRepoSql
func NewRepoMySql(db *sql.DB, tableName string, emailer email.Emailer, configuration configuration.Configuration) *ResetRepoSql {

	//Build a reset and activation config
	resetEmailConfig := PasswordResetConfig{}
	activationEmailConfig := PasswordResetConfig{}

	//Pull from the config
	configuration.GetStruct("password_reset", &resetEmailConfig)
	configuration.GetStruct("user_activation", &activationEmailConfig)

	//Define a new repo
	newRepo := ResetRepoSql{
		db:                    db,
		tableName:             tableName,
		emailer:               emailer,
		resetEmailConfig:      resetEmailConfig,
		activationEmailConfig: activationEmailConfig,
	}

	//Create the table if it is not already there
	//Create a table
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "(id int NOT NULL AUTO_INCREMENT, userId int, email TEXT, token TEXT, issued DATE, type INT, PRIMARY KEY (id) )")
	if err != nil {
		log.Fatal(err)
	}

	//Add request data to table
	addRequest, err := db.Prepare("INSERT INTO " + tableName + "(userId,email, token, issued, type) VALUES (?, ?, ?, ?, ?)")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.addRequestStatement = addRequest

	//pull the request from the table
	getRequest, err := db.Prepare("SELECT * FROM " + tableName + " where userId = ? AND token = ? AND type = ?")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.getRequestStatement = getRequest

	//pull the request from the table
	rmRequest, err := db.Prepare("delete FROM " + tableName + " where id = ? limit 1")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.rmRequestStatement = rmRequest

	//Return a point
	return &newRepo

}

//Provide a method to make a new UserRepoSql
func NewRepoPostgresSql(db *sql.DB, tableName string, emailer email.Emailer, configuration configuration.Configuration) *ResetRepoSql {
	//Build a reset and activation config
	resetEmailConfig := PasswordResetConfig{}
	activationEmailConfig := PasswordResetConfig{}

	//Pull from the config
	configuration.GetStruct("password_reset", &resetEmailConfig)
	configuration.GetStruct("user_activation", &activationEmailConfig)

	//Define a new repo
	newRepo := ResetRepoSql{
		db:                    db,
		tableName:             tableName,
		emailer:               emailer,
		resetEmailConfig:      resetEmailConfig,
		activationEmailConfig: activationEmailConfig,
	}

	//Create the table if it is not already there
	//Create a table
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "(id SERIAL PRIMARY KEY, userId int NOT NULL, email TEXT NOT NULL, token TEXT NOT NULL,issued DATE NOT NULL, type int NOT NULL)")

	if err != nil {
		log.Fatal(err)
	}

	//Add request data to table
	addRequest, err := db.Prepare("INSERT INTO " + tableName + "(userId,email, token, issued, type) VALUES ($1, $2, $3, $4, $5)")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.addRequestStatement = addRequest

	//pull the request from the table
	getRequest, err := db.Prepare("SELECT * FROM " + tableName + " where userId = $1 AND token = $2 AND type = $3")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.getRequestStatement = getRequest

	//pull the request from the table
	rmRequest, err := db.Prepare("delete FROM " + tableName + " where id = $1")
	//Check for error
	if err != nil {
		log.Fatal(err)
	}
	//Store it
	newRepo.rmRequestStatement = rmRequest

	//Return a point
	return &newRepo

}

/**
Look up the user and return if they were found
*/
func (repo *ResetRepoSql) IssueResetRequest(token string, userId int, emailAddress string) error {

	//Now add it to the database
	//Add the info
	//execute the statement//(userId,name,input,flow)- "(userId,email, token, issued)
	_, err := repo.addRequestStatement.Exec(userId, emailAddress, token, time.Now(), reset)

	//Make the email header
	header := email.HeaderInfo{
		Subject: repo.resetEmailConfig.Subject,
		To:      []string{emailAddress},
	}

	//Build a reset token
	resetInfo := PasswordResetInfo{
		Token: token,
		Email: emailAddress,
	}

	//Now email
	err = repo.emailer.SendTemplateFile(&header, repo.resetEmailConfig.Template, resetInfo, nil)

	//Return the user calcs
	return err
}

/**
Look up the user and return if they were found
*/
func (repo *ResetRepoSql) IssueActivationRequest(token string, userId int, emailAddress string) error {

	//Now add it to the database
	//Add the info
	//execute the statement//(userId,name,input,flow)- "(userId,email, token, issued)
	_, err := repo.addRequestStatement.Exec(userId, emailAddress, token, time.Now(), activation)

	//Make the email header
	header := email.HeaderInfo{
		Subject: repo.activationEmailConfig.Subject,
		To:      []string{emailAddress},
	}

	//Build a reset token
	resetInfo := PasswordResetInfo{
		Token: token,
		Email: emailAddress,
	}

	//Now email
	err = repo.emailer.SendTemplateFile(&header, repo.activationEmailConfig.Template, resetInfo, nil)

	//Return the user calcs
	return err
}

/**
Use the taken to validate
*/
func (repo *ResetRepoSql) CheckForResetToken(userId int, token string) (int, error) {

	//Get the id and errors
	id, err := repo.checkForToken(userId, token, reset)

	//If there is an error customize it
	if err != nil {
		err = errors.New("password_change_forbidden")
	}

	return id, err

}

/**
Use the taken to validate
*/
func (repo *ResetRepoSql) CheckForActivationToken(userId int, token string) (int, error) {

	//Get the id and errors
	id, err := repo.checkForToken(userId, token, activation)

	//If there is an error customize it
	if err != nil {
		err = errors.New("activation_forbidden")
	}

	return id, err

}

/**
Use the taken to validate
*/
func (repo *ResetRepoSql) checkForToken(userId int, token string, tkType tokenType) (int, error) {

	//Prepare to get values
	//id,  userId int, email TEXT, token TEXT, issued DATE,
	var id int
	var userIdDB int
	var emailDB string
	var tokenDB string
	var issued time.Time
	var tokenDb tokenType

	//Get the value
	err := repo.getRequestStatement.QueryRow(userId, token, tkType).Scan(&id, &userIdDB, &emailDB, &tokenDB, &issued, &tokenDb)

	//So it was correct, check the date
	//TODO: check the date

	//If there is an error, assume it can't be done
	if err != nil {
		return -1, errors.New("invalid_token")
	}

	//Make sure the user id and token match
	if userId != userIdDB || tokenDB != token {
		return -1, errors.New("invalid_token")
	}

	//Return the user calcs
	return id, nil
}

func (repo *ResetRepoSql) UseToken(id int) error {

	//Remove the token
	_, err := repo.rmRequestStatement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

/**
Clean up the database, nothing much to do
*/
func (repo *ResetRepoSql) CleanUp() {
	repo.getRequestStatement.Close()
	repo.addRequestStatement.Close()
	repo.rmRequestStatement.Close()

}

//func RepoDestroyCalc(id int) error {
//	for i, t := range usersList {
//		if t.id == id {
//			usersList = append(usersList[:i], usersList[i+1:]...)
//			return nil
//		}
//	}
//	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
//}
