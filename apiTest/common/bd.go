package common

import (
	"database/sql"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Primero definimos la información de configuración de la base de datos como una constante
const (
	userName = "root"
	password = ""
	ip       = "localhost"
	port     = "3306"
	dbName   = "test"
)

// Inicialice la conexión de la base de datos y devuelva la referencia del puntero de la conexión de la base de datos
func InitDB() *sql.DB {
	// Conexión de datos Golang: "Nombre de usuario: contraseña @ tcp (IP: número de puerto) / nombre de la base de datos? Charset = utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	// Abra la base de datos, el primero es el nombre del controlador, así que importe: _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	if err != nil {
		// Si abres el error de la base de datos, entra en pánico directamente
		panic(err)
	}
	// Establecer el número máximo de conexiones a la base de datos
	db.SetConnMaxLifetime(10)
	// Establecer el número máximo de conexiones inactivas en la base de datos
	db.SetMaxIdleConns(5)
	// Verificar conexión
	if err := db.Ping(); err != nil {
		panic(err)
	}
	// Devuelve la referencia del puntero de la conexión a la base de datos
	return db
}

/*
func InsertDB() {

	//InsertDB
	// Use herramientas para obtener conexión a la base de datos
	db := InitDB()
	// Abrir transacción
	tx, err := db.Begin()
	if err != nil {
		// La transacción no se abre, entra en pánico directamente
		panic(err)
	}
	// Preparar declaración SQL
	sql := "insert into  user (`name`, `phone`,`address`) values (?, ?, ?)"
	// Preprocesar la sentencia SQL
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	//result, err := stmt.Exec("Abe Dori", "123", f)
	result, err := stmt.Exec(data)
	if err != nil {
		// La ejecución de SQL falló, pánico directamente
		panic(err)
	}
	// Enviar transacción
	tx.Commit()
	// Devuelve el id del registro insertado
	fmt.Println(result.LastInsertId())

}



func SelectDB() {
	// Use herramientas para obtener conexión a la base de datos
	db := common.InitDB()
	// Preparar declaración SQL
	sql := "select * from tb_user"
	// Preprocesar la sentencia SQL
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		// La ejecución de SQL falló, pánico directamente
		panic(err)
	}
	var users []User
	for rows.Next() {
		var id int
		var name, password string
		err := rows.Scan(&id, &name, &password)
		if err != nil {
			// No se pudo leer el conjunto de resultados
			panic(err)
		}
		var user User
		user.SetId(id)
		user.SetName(name)
		//	user.SetPassword(password)
		users = append(users, user)
	}
	fmt.Println(users)
}

func UpdateDB() {
	// Use herramientas para obtener conexión a la base de datos
	db := common.InitDB()
	// Abrir transacción
	tx, err := db.Begin()
	if err != nil {
		// La transacción no se abre, entra en pánico directamente
		panic(err)
	}
	// Preparar declaración SQL
	sql := "update tb_user set `password` = ? where `id` = ?"
	// Preprocesar la sentencia SQL
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec("789", 1)
	if err != nil {
		// La ejecución de SQL falló, pánico directamente
		panic(err)
	}
	// Enviar transacción
	tx.Commit()
}

func DeleteDB() {
	// Use herramientas para obtener conexión a la base de datos
	db := common.InitDB()
	// Abrir transacción
	tx, err := db.Begin()
	if err != nil {
		// La transacción no se abre, entra en pánico directamente
		panic(err)
	}
	// Preparar declaración SQL
	sql := "delete from tb_user where `id` = ?"
	// Preprocesar la sentencia SQL
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec(1)
	if err != nil {
		// La ejecución de SQL falló, pánico directamente
		panic(err)
	}
	// Enviar transacción
	tx.Commit()
}
*/
