package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/stoewer/go-strcase"
	//
)

var csv_path = "A:\\Gitrepo\\Unlight-Zwei\\data\\csv"

var db_path = "127.0.0.1:3306"
var db_name = "unlight_old_db"
var db_user = "root:"
var db *sql.DB

var ctx context.Context

func main() {
	var err error
	// myd := mysql.New()
	db, err = sql.Open("mysql",  db_user+ "@tcp(" +  db_path+")/" + db_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("db", db)
	defer db.Close()
	
	log.Println("start up importer")
	var pathList []string
	err = filepath.Walk(csv_path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				pathList = append(pathList, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	var pipeList []*PipelineStmt
	for _, path := range pathList {
		tmp := strings.ReplaceAll(path, csv_path, "")
		p, q := CSV_importTo_MySql(&path, &tmp)
		if q != nil {
			fmt.Println("build error: please do manually")
			fmt.Println(q)
		} else {
			pipeList = append(pipeList, p...)
		}
	}

	err = WithTransaction(db, func(tx Transaction) error {
		_, err := RunPipeline(tx, pipeList...)
		return err
	})

	handleError(err)
	log.Println("Done.")

}

func CSV_importTo_MySql(path *string, name_path *string) ([]*PipelineStmt, error) {
	log.Println("in ,", *name_path)

	var pipe []*PipelineStmt

	tableName := strings.Split(
		strings.Replace(*name_path, "\\", "", 1),
		"\\")

	// fmt.Println("tableName:", tableName)

	f, err := os.Open(*path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return nil, err
	}

	header := r[0]
	val := r[1:]
	var snake = strcase.SnakeCase(strings.Replace(tableName[1], ".csv", "", 1))
	fmt.Println("tableName:", snake+`_`+tableName[0])
	create_schema := "CREATE TABLE IF NOT EXISTS `" + snake + "_" + tableName[0] + "` ("

	time_format := "2011-09-14 05:31:00 UTC"

	var para_list []string

	for k, v := range val[0] {
		varpara := ""

		if header[k] == "id" {
			varpara = " `" + header[k] + "` INT(6) UNSIGNED PRIMARY KEY "
		} else if _, err := time.Parse(time_format, v); err == nil {
			varpara = " `" + header[k] + "` VARCHAR(255) "
		} else if _, err := strconv.Atoi(v); err == nil {
			varpara = " `" + header[k] + "` INT(10) "
		} else {
			varpara = " `" + header[k] + "` VARCHAR(255) "
		}
		para_list = append(para_list, varpara)
	}

	create_schema += strings.Join(para_list, ",") + " );"

	pipe = append(pipe, NewPipelineStmt(create_schema))

	mysql.RegisterLocalFile(strings.ReplaceAll(*path, `\`, `/`))

	InsertCase := "LOAD DATA LOCAL INFILE '" + strings.ReplaceAll(*path, `\`, `/`) + "' INTO TABLE  `" + snake + "_" + tableName[0] + "`" +
		`FIELDS TERMINATED BY ','` + `ENCLOSED BY '"' ` + `LINES TERMINATED BY '\n' IGNORE 1 ROWS;	`

	pipe = append(pipe, NewPipelineStmt(InsertCase))

	return pipe, nil
}

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TxFn func(Transaction) error

func WithTransaction(db *sql.DB, fn TxFn) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

type PipelineStmt struct {
	query string
	args  []interface{}
}

func NewPipelineStmt(query string, args ...interface{}) *PipelineStmt {
	return &PipelineStmt{query, args}
}

func (ps *PipelineStmt) Exec(tx Transaction, lastInsertId int64) (sql.Result, error) {
	query := strings.Replace(ps.query, "{LAST_INS_ID}", strconv.Itoa(int(lastInsertId)), -1)
	return tx.Exec(query, ps.args...)
}

func RunPipeline(tx Transaction, stmts ...*PipelineStmt) (sql.Result, error) {
	var res sql.Result
	var err error
	var lastInsId int64

	for k, ps := range stmts {
		fmt.Println("ex:", k)
		res, err = ps.Exec(tx, lastInsId)
		if err != nil {
			fmt.Println("query?:", ps.query)
			return nil, err
		}

		lastInsId, err = res.LastInsertId()
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
