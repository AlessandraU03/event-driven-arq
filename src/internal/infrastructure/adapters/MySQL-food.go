package adapters

import (
	"eventdriven/src/core"
	"eventdriven/src/internal/domain/entities"
	"fmt"
	"log"
	"strconv"
)

type MySQLFoods struct {
	conn *core.Conn_MySQL
}

func NewMySQLFoods() *MySQLFoods {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLFoods{conn: conn}
}

func (mysql *MySQLFoods) Save(food *entities.Food) error {
	query := "INSERT INTO productos (nombre, descripcion, precio) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, food.Nombre, food.Descripcion, food.Precio)
	if err != nil {
		return fmt.Errorf("error al guardar la comida: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQLFoods) GetAll() ([]*entities.Food, error) {
	query := "SELECT producto_id, nombre, descripcion, precio FROM productos"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var foods []*entities.Food
	for rows.Next() {
		var id int32
		var nombre, descripcion, precioStr string
		if err := rows.Scan(&id, &nombre, &descripcion, &precioStr); err != nil {
			return nil, fmt.Errorf("error al escanear la comida: %w", err)
		}
		precio, err := strconv.ParseFloat(precioStr, 64)
		if err != nil {
			return nil, fmt.Errorf("error al convertir el precio a float64: %w", err)
		}
		
		food := &entities.Food{
			ID:          int(id),
			Nombre:      nombre,
			Descripcion: descripcion,
			Precio:      precio,
		}
		foods = append(foods, food)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre los libros: %w", err)
	}

	return foods, nil
}

func (mysql *MySQLFoods) Update(food *entities.Food) error {
	query := "UPDATE productos SET nombre = ?, descripcion = ?, precio = ? WHERE producto_id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, food.Nombre, food.Descripcion, food.Precio, food.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar la comida: %w", err)
	}
	return nil
}

func (mysql *MySQLFoods) Delete(FoodID int32) error { 
	query := "DELETE FROM productos WHERE producto_id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, FoodID)
	if err != nil {
		return fmt.Errorf("error al eliminar la comida: %w", err)
	}
	return nil
}

func (mysql *MySQLFoods) GetById(FoodID int32) ([]*entities.Food, error) {
    query := "SELECT producto_id, nombre, descripcion, precio FROM productos WHERE producto_id = ?"
    rows := mysql.conn.FetchRows(query, FoodID)
    defer rows.Close()

    var foods []*entities.Food

    for rows.Next() {
        var idFood int32
        var nombre, descripcion, precioStr string

        if err := rows.Scan(&idFood, &nombre, &descripcion, &precioStr); err != nil {
            return nil, fmt.Errorf("error al escanear la comida: %w", err)
        }

        precio, err := strconv.ParseFloat(precioStr, 64)
        if err != nil {
            return nil, fmt.Errorf("error al convertir el precio a float64: %w", err)
        }

        food := &entities.Food{
            ID:          int(idFood),
            Nombre:      nombre,
            Descripcion: descripcion,
            Precio:      precio,
        }
        foods = append(foods, food)
    }

    if len(foods) == 0 {
        return nil, fmt.Errorf("comida no encontrada con id: %d", FoodID)
    }

    return foods, nil
}

