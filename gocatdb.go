package gocatdb

import (
"bytes"
"database/sql"
"log"
"reflect"
"fmt"
)

type Catdb struct{
    Connected int
    db *sql.DB
    tablename string
    sql string
    Columns []string
}

func (this *Catdb) BindDb(db *sql.DB) {
    this.db = db
    this.Connected = 1
}

func (this *Catdb) Table(tablename string)(*Catdb) {
    this.tablename = tablename
    return this
}

func (this *Catdb) Insert(data interface{}) (*Catdb){
    paramType := reflect.TypeOf(data)
    // fmt.Println(paramType)
    // fmt.Println(paramType.Kind())
    switch paramType.Kind(){
        case reflect.Slice:
        fmt.Println("not support")
        case reflect.Map:
        if paramType.Key().Kind()==reflect.String && paramType.Elem().Kind()==reflect.Interface {
            this.insert_map(data.(map[string]interface{}))

        }
    }







    return this
}




func GetSqliteType(i interface{}) string{
     switch i.(type){
        case int:
            return "integer"
        case float64:
            return "float"
        case string:
            return "text"
        case bool:
            return "integer"
        case []byte:
            return "blob"
        // case time.Time:
            // return "datetime"
        default:
            return "text"
    }
}

func (this *Catdb) Create(data map[string]interface{}) (*Catdb){
    var sql bytes.Buffer

    sql.WriteString("create table ")
    sql.WriteString(this.tablename)
    sql.WriteString(" ( ")
    first := true
    for columnName ,value := range data{
        if first==false {
            sql.WriteString(" , ")
        }
        sql.WriteString("`"+columnName+"`")
        sql.WriteString(" ")
        sql.WriteString(GetSqliteType(value))
        first = false
    }
    sql.WriteString(" ); ")

    this.sql = sql.String()
    return this
}




func (this *Catdb) Sql() (string){
    return this.sql
}

func (this *Catdb) Execute() (sql.Result){
    stmt,err := this.db.Prepare(this.sql)
    if err!=nil {
        log.Println(err)
    }
    res,err := stmt.Exec()
    return res
    // return this.sql
}


