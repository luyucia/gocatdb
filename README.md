# gocatdb


轻量级orm框架
设计理念：
1.不封装查询类sql，查询类sql很多情况相对复杂，框架封装后带来的学习成本较高，不如直接写sql来得快，对于条件拼接，框架可以提供辅助工具。
2.insert update支持结构体，map，甚至直接传json

Example:
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
    cdb.BindDb(db,"sqlite3")

    cdb.Table("test").Create(data[0])


    cdb.Table("test").Insert(data)


    fmt.Println(cdb.Query("select pv,uv,city from test where id =1"))



```
### 初始化
```
cdb := gocatdb.Catdb{}
cdb.BindDb(db,"mysql")
db为数据库连接
mysql为数据库方言类型
目前支持:
- mysql
- sqlite3

```
### 根据map创建表
```
cdb.Table("test").Create(mapdata)
```
### 插入数据
```
cdb.Table("test").Insert(data)
data 可以是一个map
可以是一个map slice

```
### 查询
```
cdb.Query(sql)
返回一个map slice
```
### 修改
例:将id为2的city修改为"吉林"
```
cdb.Table("test").Update(map[string]interface{}{"city":"吉林"},"id=2")
```

### 删除
例:将id为1的数据
```
cdb.Table("test").Delete("id=1")
```
