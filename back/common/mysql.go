package common

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	ConnectStr   string //user:password@tcp(ip:port)/database
	MaxOpenConns int
	MaxIdleConns int
}

//
var (
	MysqlDbHandle *sql.DB
	ProgramClose  bool
)

// MysqlDbHandle is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a DB.
// Your routine must first initialize the log
func MysqlInit(mysqlconf MysqlConfig) (err error) {

	//TODO optimzation log

	if MysqlDbHandle != nil {
		MysqlDbHandle.Close()
		MysqlDbHandle = nil
	}
	//	fmt.Println("Mysql mysqlconf 	!", mysqlconf)
	MysqlDbHandle = &sql.DB{}
	//MysqlDbHandle, err = sql.Open("mysql", "user:password@tcp(ip:port)/database")
	MysqlDbHandle, err = sql.Open("mysql", mysqlconf.ConnectStr)
	if err != nil {
		fmt.Println("Mysql Open error 	!", err)
		return
	}
	MysqlDbHandle.SetMaxOpenConns(mysqlconf.MaxOpenConns)
	MysqlDbHandle.SetMaxIdleConns(mysqlconf.MaxIdleConns)
	err = MysqlDbHandle.Ping()
	if err != nil {
		fmt.Println("Mysql ping error 	!", err)
		return
	}
	return
}

// Close closes the database, releasing any open resources.
func UnInit() {
	if MysqlDbHandle != nil {
		err := MysqlDbHandle.Close()
		if err != nil {
			fmt.Println("Mysql close error 	!", err)
		}
		MysqlDbHandle = nil
	}
}

func GetOneValueFromDB(selectSql string, args ...interface{}) (valueStr string, err error) {

	// get total record
	stm, err := MysqlDbHandle.Prepare(selectSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()
	rows, err := stm.Query(args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var nilSTRING sql.NullString
	for rows.Next() {
		if err = rows.Scan(&nilSTRING); err != nil {
			fmt.Println("GetOneRecordFromDB sql: ", selectSql, "err:", err)
		}

	}

	if nilSTRING.Valid {
		valueStr = nilSTRING.String
	} else {

	}
	return
}

func GetOneIntegerFromDB(selectSql string, args ...interface{}) (valueInt int64, err error) {

	valueStr, err := GetOneValueFromDB(selectSql, args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(valueStr) == 0 {
		return 0, nil
	}
	valueInt, err = strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func GetTwoValueFromDB(selectSql string, args ...interface{}) (OneStr string, TwoStr string, err error) {

	// get total record
	stm, err := MysqlDbHandle.Prepare(selectSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()
	rows, err := stm.Query(args...)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var nilSTRINGOne sql.NullString
	var nilSTRINGTwo sql.NullString
	for rows.Next() {
		if err = rows.Scan(&nilSTRINGOne, &nilSTRINGTwo); err != nil {
			fmt.Println("GetOneRecordFromDB sql: ", selectSql, "err:", err)
		}

	}

	if nilSTRINGOne.Valid {
		OneStr = nilSTRINGOne.String
	} else {

	}
	if nilSTRINGTwo.Valid {
		TwoStr = nilSTRINGTwo.String
	} else {

	}
	return
}

func ExcSqlFromDB(selectSql string, args ...interface{}) (lastInsertid int64, rowsAffected int64, err error) {

	// get total record
	stm, err := MysqlDbHandle.Prepare(selectSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()

	ret, err := stm.Exec(args...)
	if err != nil {
		fmt.Println(err)
		showProcesslist()
		return
	}
	lastInsertid, err = ret.LastInsertId()
	if nil == err {
		//fmt.Println("ExcSqlFromDB LastInsertId:", lastInsertid)
	}
	if rowsAffected, err = ret.RowsAffected(); nil == err {
		//fmt.Println("ExcSqlFromDB RowsAffected:", rowsAffected)
	}
	return
}

func GetOneSliceFromDB(selectSql string, args ...interface{}) (vArr []string, err error) {
	//start := time.Now()

	//gained ==0 ,the bet transaction failed/loser
	stm, err := MysqlDbHandle.Prepare(selectSql)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()

	rows, err := stm.Query(args...)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {

		var ValueStr sql.NullString
		if err = rows.Scan(&ValueStr); err != nil {
			fmt.Println(err)
			return
		}

		if ValueStr.Valid {
			vArr = append(vArr, ValueStr.String)
		}

	}

	//end := time.Now()
	//Log.Infoln("GetOneSliceFromDB total time:", end.Sub(start).Seconds())
	return
}

func showProcesslist() {
	// get total record
	stm, err := MysqlDbHandle.Prepare("show full processlist")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()
	var id string
	var user string
	var host string
	var db string
	var command string
	var time string
	var state string
	var info string

	rows, err := stm.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("show processlit sql : id ,user,host,db,command ,time,state ,info ")
	for rows.Next() {
		if err = rows.Scan(&id, &user, &host, &db, &command, &time, &state, &info); err != nil {

		}
		fmt.Println(id, user, host, db, command, time, state, info)
	}

	return
}
func showLocklist() {
	// get total record
	// var MysqlDbHandle *sql.DB
	// MysqlDbHandle = &sql.DB{}
	// //MysqlDbHandle, err = sql.Open("mysql", "user:password@tcp(ip:port)/database")
	// MysqlDbHandle, err := sql.Open("mysql", "root:scry@tcp(127.0.0.1:3306)/information_schema")
	// if err != nil {
	// 	fmt.Println("Mysql Open error 	!", err)
	// 	return
	// }
	// MysqlDbHandle.SetMaxOpenConns(10)
	// MysqlDbHandle.SetMaxIdleConns(10)
	// defer MysqlDbHandle.Close()

	stm, err := MysqlDbHandle.Prepare("select trx_id,trx_state,trx_started,trx_query from information_schema.innodb_trx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stm.Close()
	var trxId string
	var trxState string
	var trxStarted string
	var trxQuery string

	rows, err := stm.Query()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("show processlit sql : id ,user,host,db,command ,time,state ,info ")
	for rows.Next() {
		if err = rows.Scan(&trxId, &trxState, &trxStarted, &trxQuery); err != nil {

		}
		fmt.Println(trxId, "\n", trxState, "\n", trxStarted, "\n", trxQuery)
	}

	return
}
func MysqlEndCommit(mysqlTx *sql.Tx) (err error) {
	err = mysqlTx.Commit()
	fmt.Println("sqlEndCommit ")
	if err != nil {
		fmt.Println("!!!!!!!!!!!", err)

		showProcesslist()
		showLocklist()
		return err
	}
	return err

}

func MysqlEndRollBack(mysqlTx *sql.Tx, dbCommit *bool) (err error) {
	if *dbCommit {
		return nil
	}
	err = mysqlTx.Rollback()
	fmt.Println("MysqlEndRollBack ")
	if err != nil {
		fmt.Println("!!!!!!!!!!!", err)
		showProcesslist()
		showLocklist()
		return err
	}
	return err

}
