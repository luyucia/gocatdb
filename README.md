# gocatdb

```
data := []map[string]interface{}{
        map[string]interface{}{"id":1,"uv":101.1,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":1,"uv":111.0,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":2,"uv":100.0,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":2,"uv":100.5,"pv":10,"name":"luyu","city":"长春"},
        map[string]interface{}{"id":3,"uv":100.3,"pv":10,"name":"luyu","city":"长春"},
    }


    os.Remove("./foo.db")
    db, err := sql.Open("sqlite3", "./foo.db")

    defer db.Close()
    if err!=nil {
        fmt.Println(err)
    }
    cdb := gocatdb.Catdb{}
    cdb.BindDb(db,"mysql")

    cdb.Table("test").Create(data[0]).Execute()


    cdb.Table("test").Insert(data)


    fmt.Println(cdb.Query("select pv,uv,city from test where id =1"))



```
