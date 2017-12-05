package gocatdb

import (
"bytes"
// "database/sql"
// "log"
// "reflect"
"fmt"
)



func (this *Catdb) insert_map(data map[string]interface{}){
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
