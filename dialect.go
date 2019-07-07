package gocatdb

type Dialect struct{
	Dbtype string

}


func (this *Dialect) GetType(v interface{}) string {
	if this.Dbtype=="sqlite3" {
		return this.getSqliteType(v)
	}else if this.Dbtype=="mysql"{
		return this.getMysqlType(v)
	}
	return ""
}

func (this *Dialect) getSqliteType(i interface{}) string{
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
func (this *Dialect) getMysqlType(i interface{}) string{
	switch i.(type){
	case int:
		return "INT(11)"
	case float64:
		return "DOUBLE"
	case string:
		return "VARCHAR(255)"
	case bool:
		return "TINYINT"
	case []byte:
		return "BLOB"
	// case time.Time:
	// return "datetime"
	default:
		return "text"
	}
}
