package db

type Filter map[string]interface{}

type Attr map[string]struct{
    Select []string
    Limit  int
    Offset int
    OrderBy string
    GroupBy string
}


// INSERT INTO products (name, code) VALUES ("name", "code")
func insert(data map[string]interface{}) {

}

func insertMulti()  {

}
