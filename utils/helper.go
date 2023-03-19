package utils

import (
	"database/sql"
	"encoding/json"
	"strings"

	"go.uber.org/zap"
)

func SQLType(sql string) (types string) {
	sqlReader := strings.Split(strings.Trim(sql, " "), " ")

	switch strings.ToUpper(sqlReader[0]) {
	case "INSERT", "UPDATE", "DELETE", "MERGE", "LOCK", "UNLOCK", "EXECUTE":
		types = "DML"
	case "SELECT", "DESC":
		types = "DQL"
	case "CREATE", "ALTER", "DROP", "TRUNCATE", "RENAME":
		types = "DDL"
	case "GRANT", "REVOKE":
		types = "DCL"
	case "COMMIT", "ROLLBACK", "SAVEPOINT":
		types = "TCL"
	case "SET", "SHOW", "USE", "HELP", "EXPLAIN":
		types = "PLUGIN"
	default:
		types = ""
	}

	return types
}

func AssigningRawByte(values []interface{}) []interface{} {
	for idx := range values {
		values[idx] = new(sql.RawBytes)
	}

	return values
}

func ColumnsJSON(columns []string, log *zap.Logger) (string, error) {
	data := make(map[string]string)

	for idx := range columns {
		data[columns[idx]] = ""
	}

	stringify, err := json.Marshal(data)
	if err != nil {
		log.Error("Error while converting json string")

		return string(stringify), err
	}

	return string(stringify), err
}

func DataTojson(log *zap.Logger, data []interface{}) (string, error) {
	stringify, err := json.Marshal(data)
	if err != nil {
		log.Error("Error while stringify")

		return string(stringify), err
	}

	return string(stringify), err
}
