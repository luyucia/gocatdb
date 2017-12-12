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
    dialect Dialect
}

func (this *Catdb) BindDb(db *sql.DB,dialect string) {
    this.dialect = Dialect{Dbtype:dialect}
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
    // 根据不同类型判断执行操作
    switch paramType.Kind(){
        case reflect.Slice:
        if paramType.Elem().Kind() == reflect.Map{
            if paramType.Elem().Key().Kind()==reflect.String && paramType.Elem().Elem().Kind()==reflect.Interface {
                this.insert_map_slice(data.([]map[string]interface{}))
            }

        }
        case reflect.Map:
        if paramType.Key().Kind()==reflect.String && paramType.Elem().Kind()==reflect.Interface {
            this.insert_map(data.(map[string]interface{}))
        }

    }

    return this
}

func (this *Catdb) Update(data map[string]interface{},condition string) {
    var sql bytes.Buffer

    sql.WriteString("UPDATE ")
    sql.WriteString(this.tablename)
    sql.WriteString(" set ")
    // 遍历key放入slice 保持顺序
    vals := []interface{}{}
    first := true
    for column,val := range data{
        vals = append(vals,val)
        if first==false {
            sql.WriteString(",")
        }
        sql.WriteString(column+"=?")
        first = false
    }
    sql.WriteString(" where ")
    sql.WriteString(condition)

    this.sql = sql.String()
    fmt.Println(this.sql)
    this.Exec(vals...)
}


func (this *Catdb) Delete(condition string)(sql.Result) {
    var sql bytes.Buffer

    sql.WriteString("DELETE FROM ")
    sql.WriteString(this.tablename)
    sql.WriteString(" where ")
    sql.WriteString(condition)

    this.sql = sql.String()
    fmt.Println(this.sql)
    return this.Execute()
}



func (this *Catdb) Create(data map[string]interface{}) (sql.Result){
    var sql bytes.Buffer

    sql.WriteString("CREATE TABLE ")
    sql.WriteString(this.tablename)
    sql.WriteString(" ( ")
    first := true
    for columnName ,value := range data{
        if first==false {
            sql.WriteString(" , ")
        }
        sql.WriteString("`"+columnName+"`")
        sql.WriteString(" ")
        sql.WriteString(this.dialect.GetType(value))
        first = false
    }
    sql.WriteString(" ); ")

    this.sql = sql.String()
    return this.Execute()
    // return this
}


func (this *Catdb) Query(query string) ([]map[string]interface{}){
    rows, err := this.db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    // defer rows.Close()
    columns, _ := rows.Columns()
    length := len(columns)

    rs := []map[string]interface{}{}

    for rows.Next(){
        row := make([]string,length)
        columnPointers := make([]interface{}, length)
        for i:=0;i<length;i++ {
            columnPointers[i] = &row[i]
        }
        rows.Scan(columnPointers...)
        tmpmap := make(map[string]interface{})

        for i:=0 ; i<length ;i++{
            columnName := columns[i]
            columnValue := row[i]
            // fmt.Println(reflect.ValueOf(columnValue).Type())
            tmpmap[columnName] = columnValue

        }
        rs = append(rs,tmpmap)
    }
    return rs
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
    if err!=nil{
        log.Println(err)
    }
    return res
    // return this.sql
}

func (this *Catdb) Exec(args ...interface{}) {
    stmt,err := this.db.Prepare(this.sql)
    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    _,err = stmt.Exec(args...)

    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    stmt.Close()
}


