package domain

import (
	"database/sql"
	"fmt"
	"reflect"
	"sql-service/model"
	"sql-service/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB  *gorm.DB
	Log *zap.Logger
}

type IStorage interface {
	Console(string) model.Response
	ConsoleDML(string) model.Response
}

func NewStorage(DB *gorm.DB, Log *zap.Logger) IStorage {
	return &DBInstance{DB: DB, Log: Log}
}

func (DB *DBInstance) Console(sql string) model.Response {

	simpleDB, _ := DB.DB.DB()

	resultSet, err := GetResultSet(sql, DB.Log, simpleDB)
	if err != nil {
		DB.Log.Error("Error while fetching result set", zap.Error(err))

		return model.Response{
			Data: "",
			Msg:  err.Error(),
		}
	}

	columns := ResultSetColumns(resultSet)

	data := DescribeResultSet(resultSet, columns, DB.Log)
	if len(data) == 0 {
		data, err := utils.ColumnsJSON(columns, DB.Log)
		if err != nil {
			DB.Log.Error("Uable to Unmarshal columns", zap.Error(err))

			return model.Response{
				Data: "",
				Msg:  err.Error(),
			}
		}

		return model.Response{
			Data: fmt.Sprintf("[%s]", data),
			Msg:  "Successfully Executed",
		}
	}

	stringify, err := utils.DataTojson(DB.Log, data)
	if err != nil {
		DB.Log.Error("error while generating JSON string", zap.Error(err))

		return model.Response{
			Data: stringify,
			Msg:  err.Error(),
		}
	}

	return model.Response{
		Data: stringify,
		Msg:  "Successfully Executed",
	}
}

func GetResultSet(sql string, log *zap.Logger,
	dbinstanse *sql.DB) (resultset *sql.Rows, err error) {

	resultset, err = dbinstanse.Query(sql)
	if err != nil {

		return resultset, err
	}

	return resultset, err
}

func ResultSetColumns(resultset *sql.Rows) (columns []string) {
	columns, _ = resultset.Columns()

	return columns
}

func DescribeResultSet(resultset *sql.Rows, columns []string, log *zap.Logger) []interface{} {
	finalData := make([]interface{}, 0)

	for resultset.Next() {
		values := make([]interface{}, len(columns))

		resultValue := make(map[string]interface{})

		values = utils.AssigningRawByte(values)

		columnType, err := resultset.ColumnTypes()
		fmt.Println("🚀 ~ file: console_dao.go ~ line 107 ~ forresultset.Next ~ data, err : ", columnType, err)

		if err := resultset.Scan(values...); err != nil {
			log.Error("Implementation Error", zap.Error(err))

			return finalData
		}

		for idx, value := range values {
			content := reflect.ValueOf(value).Interface().(*sql.RawBytes)

			data, checker := resultValue[columns[idx]]
			if checker != true {

				columntype := *columnType[idx]

				switch columntype.DatabaseTypeName() {
				case "INT":
					resultValue[columns[idx]] = int(content) // Assuming INT is 64-bit
				case "VARCHAR", "TEXT":
					resultValue[columns[idx]] = string(*content)
				case "BOOL":
					resultValue[columns[idx]] = bool(content.Bool())
				case "FLOAT", "DOUBLE":
					resultValue[columns[idx]] = float64(content.Float64())
				// Add more cases for other types as needed
				default:

				}

				continue
			}

			data = string(*content)

			log.Info("Found Data :", zap.String("Data", data))

			resultValue[columns[idx]] = data
		}

		finalData = append(finalData, resultValue)
	}

	return finalData
}

func (DB *DBInstance) ConsoleDML(sql string) model.Response {

	dbinstanse, _ := DB.DB.DB()

	result, err := dbinstanse.Exec(sql)
	if err != nil {
		DB.Log.Error("Error on DML query", zap.Error(err))

		return model.Response{
			Data: "",
			Msg:  err.Error(),
		}
	}

	affected, _ := result.RowsAffected()

	return model.Response{
		Data: "",
		Msg:  fmt.Sprintf("Affected Records : %d", affected),
	}
}
