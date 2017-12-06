package gocatdb

import (
"bytes"
// "database/sql"
// "log"
// "reflect"
"fmt"
)


func (this *Catdb) parse_column(data map[string]interface{}) {

}

func (this *Catdb) insert_map_slice(data []map[string]interface{}){

    this.Columns = make([]string,0)
    for field,_ := range data[0]{
        this.Columns = append(this.Columns,field)
    }
    this.build_batch_insert(len(data))

    values := []interface{}{}
    for _,row := range data{
        for _,field := range this.Columns{
            values = append(values,row[field])
        }
    }

    stmt,err := this.db.Prepare(this.sql)
    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    _,err = stmt.Exec(values...)

    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    stmt.Close()

}

func (this *Catdb) insert_map(data map[string]interface{}){

    this.Columns = make([]string,0)
    for field,_ := range data{
        this.Columns = append(this.Columns,field)
    }
    this.build_single_insert()

    values := []interface{}{}
    for _,field := range this.Columns{
        values = append(values,data[field])
    }


    stmt,err := this.db.Prepare(this.sql)
    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    _,err = stmt.Exec(values...)

    if err!=nil {
        fmt.Println("Exec error:", err)
        panic(err)
    }

    stmt.Close()

}


func (this *Catdb) build_batch_insert(num int) (string){
    var sql bytes.Buffer
    var values bytes.Buffer

    sql.WriteString("insert into ")
    sql.WriteString(this.tablename)
    sql.WriteString(" ( ")
    first := true
    for _,field := range this.Columns{
        if first==false {
            sql.WriteString(",")
            values.WriteString(",")
        }

        values.WriteString("?")
        sql.WriteString("`"+field+"`")
        first = false
    }
    sql.WriteString(" ) values ")
    first = true
    for i:=0;i<num;i++{
        if first==false {
            sql.WriteString(",")
        }
        sql.WriteString("("+values.String()+")")
        first = false
    }

    this.sql = sql.String()
    return this.sql
}

func (this *Catdb) build_single_insert() (string){
    var sql bytes.Buffer
    var values bytes.Buffer

    sql.WriteString("insert into ")
    sql.WriteString(this.tablename)
    sql.WriteString(" ( ")
    first := true
    for _,field := range this.Columns{
        if first==false {
            sql.WriteString(",")
            values.WriteString(",")
        }

        values.WriteString("?")
        sql.WriteString("`"+field+"`")
        first = false
    }
    sql.WriteString(" ) values (")
    sql.WriteString(values.String())
    sql.WriteString(")")

    this.sql = sql.String()
    return this.sql
}
