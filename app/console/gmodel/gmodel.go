package gmodel

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"regexp"
	"strconv"
	"strings"
	"thh/arms"
	"thh/arms/config"
	"thh/arms/ehandle"
	"thh/arms/str"
)

func init() {
	appendCommand(CmdGmodel)
}

//go:embed stubs
var stubsFS embed.FS

var CmdGmodel = &cobra.Command{
	Use:   "gmodel",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runGmodel,
	//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

type modelEm struct {
	Name  string // 字段名
	Notes string // 注释
	Type  string // 字段类型
}
type genColumns struct {
	Field   string  `gorm:"column:Field"`
	Type    string  `gorm:"column:Type"`
	Key     string  `gorm:"column:Key"`
	Desc    string  `gorm:"column:Comment"`
	Null    string  `gorm:"column:Null"`
	Default *string `gorm:"column:Default"`
}

func runGmodel(cmd *cobra.Command, args []string) {
	dataSourceName := config.GetString("TMP_DATABASE_URL")
	outputRoot := config.GetString("GMODEL_OUTPUT_DIR", "./storage/tmp/model/")
	dbStd := `"dt/util/db"`
	dbStd = `"thh/helpers/db"`
	dbStd = `"thh/conf/dbconnect"`
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{PrepareStmt: false,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 全局禁用表名复数
		Logger:         logger.Default})
	if ehandle.PrIF(err) {
		return
	}

	rows, err := db.Raw("show tables").Rows()
	tbDesc := make(map[string]string)
	if ehandle.PrIF(err) {
		return
	}
	for rows.Next() {
		var table string
		ehandle.PrIF(rows.Scan(&table))
		tbDesc[table] = table
	}
	ehandle.PrIF(rows.Close())

	modelData, _ := stubsFS.ReadFile("stubs/entity.stub")
	connectData, _ := stubsFS.ReadFile("stubs/connect.stub")
	repData, _ := stubsFS.ReadFile("stubs/rep.stub")
	for tmpTableName, _ := range tbDesc {
		modelStr := string(modelData)
		connectStr := string(connectData)
		repStr := string(repData)
		emListStr := bytes.Buffer{}
		fieldListStr := bytes.Buffer{}
		importStr := bytes.Buffer{}
		importList := map[string]string{}
		//rows, err := db.Raw("show create table " + k).Rows()
		//if ehandle.PrIF(err) {
		//	continue
		//}
		//if rows.Next() {
		//	var table, CreateTable string
		//	ehandle.PrIF(rows.Scan(&table, &CreateTable))
		//	fmt.Println(CreateTable)
		//}
		//ehandle.PrIF(rows.Close())
		var list []genColumns

		// Get table annotations.获取表注释
		db.Raw("show FULL COLUMNS from " + tmpTableName).Scan(&list)
		fieldListStr.WriteString(fmt.Sprintf("const tableName = \"%v\"\n", tmpTableName))

		for _, value := range list {
			var field string
			if IsNum(string(value.Field[0])) {
				field = "Column" + str.Camel(value.Field)
			} else {
				field = str.Camel(value.Field)
			}
			if pkgname, ok := EImportsHead[getTypeName(value.Type, false)]; ok {
				importList[pkgname] = pkgname
			}

			constName := ""
			pidStr := ""
			nullStr := ""
			if value.Key == "PRI" {
				pidStr = "primaryKey;"
				constName = "pid"
			} else {
				constName = "field" + str.Camel(str.LowerCamel(value.Field))
			}
			fieldListStr.WriteString(fmt.Sprintf("const %v = \"%v\"\n", constName, value.Field))

			if value.Null == "NO" {
				nullStr = "not null;"
			}

			defaultStr := ""
			if value.Default != nil {
				defaultStr = "default:"
				if len(*value.Default) == 0 {
					defaultStr += "''"
				} else {
					defaultStr += *value.Default
				}
				defaultStr += ";"
			}
			typeString := `type:` + value.Type
			if pidStr != "" {
				typeString = `autoIncrement`
			}

			emListStr.WriteString(fmt.Sprintf("\t%v\t%v\t`gorm:\"%vcolumn:%v;%v;%v%v\" json:\"%v\"` // %v \n",
				field, getTypeName(value.Type, value.Null != "NO"), pidStr, value.Field, typeString, nullStr, defaultStr, str.LowerCamel(value.Field), value.Desc,
			))
		}
		modelStr = strings.ReplaceAll(modelStr, "#{EmList}", emListStr.String())
		modelStr = strings.ReplaceAll(modelStr, "#{TableName}", tmpTableName)
		if IsNum(string(tmpTableName[0])) {
			tmpTableName = "M" + tmpTableName
		}
		modelStr = strings.ReplaceAll(modelStr, "#{ModelName}", str.Camel(tmpTableName))

		for pkgname, _ := range importList {
			if importStr.Len() == 0 {
				importStr.WriteString("import (")
				importStr.WriteString("\n")
			}
			importStr.WriteString(pkgname)
			importStr.WriteString("\n")
		}
		if importStr.Len() != 0 {
			importStr.WriteString(")")
		}

		modelStr = strings.ReplaceAll(modelStr, "#{Import}", importStr.String())
		modelStr = strings.ReplaceAll(modelStr, "#{Field}", fieldListStr.String())
		repStr = strings.ReplaceAll(repStr, "#{ModelName}", str.Camel(tmpTableName))
		connectStr = strings.ReplaceAll(connectStr, "#{ModelName}", str.Camel(tmpTableName))
		connectStr = strings.ReplaceAll(connectStr, "#{ModelName}", str.Camel(tmpTableName))
		connectStr = strings.ReplaceAll(connectStr, "#{DBPkg}", dbStd)

		//outputRoot := "tmp/"
		modelPath := str.Camel(tmpTableName)
		_ = arms.FilePutContents(outputRoot+modelPath+"/"+modelPath+".go", []byte(modelStr), false)
		_ = arms.FilePutContents(outputRoot+modelPath+"/"+modelPath+"_rep.go", []byte(repStr), false)
		_ = arms.FilePutContents(outputRoot+modelPath+"/"+modelPath+"_connect.go", []byte(connectStr), false)
		fmt.Println(modelStr)
		fmt.Println(repStr)
	}
	//c := exec.Command("gofmt", "-w", "tmp/*")
	//err = c.Run()
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println("end")

}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// getTypeName Type acquisition filtering.类型获取过滤
func getTypeName(name string, isNull bool) string {
	// 优先匹配自定义类型

	// Precise matching first.先精确匹配
	if v, ok := TypeMysqlDicMp[name]; ok {
		return fixNullToPorint(v, isNull)
	}

	// Fuzzy Regular Matching.模糊正则匹配
	for _, l := range TypeMysqlMatchList {
		if ok, _ := regexp.MatchString(l.Key, name); ok {
			return fixNullToPorint(l.Value, isNull)
		}
	}

	panic(fmt.Sprintf("type (%v) not match in any way.maybe need to add on (https://github.com/xxjwxc/gormt/blob/master/data/view/cnf/def.go)", name))
}

func fixNullToPorint(name string, isNull bool) string {
	if isNull {
		if strings.HasPrefix(name, "uint") {
			return "*" + name
		}
		if strings.HasPrefix(name, "int") {
			return "*" + name
		}
		if strings.HasPrefix(name, "float") {
			return "*" + name
		}
		if strings.HasPrefix(name, "date") {
			return "*" + name
		}
		if strings.HasPrefix(name, "time") {
			return "*" + name
		}
		if strings.HasPrefix(name, "bool") {
			return "*" + name
		}
		if strings.HasPrefix(name, "string") {
			return "*" + name
		}
	}

	return name
}

var EImportsHead = map[string]string{
	"stirng":         `"string"`,
	"time.Time":      `"time"`,
	"gorm.Model":     `"gorm.io/gorm"`,
	"fmt":            `"fmt"`,
	"datatypes.JSON": `"gorm.io/datatypes"`,
	"datatypes.Date": `"gorm.io/datatypes"`,
}

var TypeMysqlDicMp = map[string]string{
	"smallint":            "int16",
	"smallint unsigned":   "uint16",
	"int":                 "int",
	"int unsigned":        "uint",
	"bigint":              "int64",
	"bigint unsigned":     "uint64",
	"mediumint":           "int32",
	"mediumint unsigned":  "uint32",
	"varchar":             "string",
	"char":                "string",
	"date":                "datatypes.Date",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int8",
	"tinyint unsigned":    "uint8",
	"tinyint(1)":          "int", // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "int", // tinyint(1) 默认设置成bool
	"json":                "datatypes.JSON",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
	"numeric":             "float64",
	"smalldatetime":       "time.Time", //sqlserver
	"nvarchar":            "string",
	"real":                "float32",
	"binary":              "[]byte",
}

var TypeMysqlMatchList = []struct {
	Key   string
	Value string
}{
	{`^(tinyint)[(]\d+[)] unsigned`, "uint8"},
	{`^(smallint)[(]\d+[)] unsigned`, "uint16"},
	{`^(int)[(]\d+[)] unsigned`, "uint32"},
	{`^(bigint)[(]\d+[)] unsigned`, "uint64"},
	{`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(tinyint)[(]\d+[)]`, "int8"},
	{`^(smallint)[(]\d+[)]`, "int16"},
	{`^(int)[(]\d+[)]`, "int"},
	{`^(bigint)[(]\d+[)]`, "int64"},
	{`^(char)[(]\d+[)]`, "string"},
	{`^(enum)[(](.)+[)]`, "string"},
	{`^(varchar)[(]\d+[)]`, "string"},
	{`^(varbinary)[(]\d+[)]`, "[]byte"},
	{`^(blob)[(]\d+[)]`, "[]byte"},
	{`^(binary)[(]\d+[)]`, "[]byte"},
	{`^(decimal)[(]\d+,\d+[)]`, "float64"},
	{`^(mediumint)[(]\d+[)]`, "int16"},
	{`^(mediumint)[(]\d+[)] unsigned`, "uint16"},
	{`^(double)[(]\d+,\d+[)]`, "float64"},
	{`^(float)[(]\d+,\d+[)]`, "float64"},
	{`^(datetime)[(]\d+[)]`, "time.Time"},
	{`^(bit)[(]\d+[)]`, "[]uint8"},
	{`^(text)[(]\d+[)]`, "string"},
	{`^(integer)[(]\d+[)]`, "int"},
	{`^(timestamp)[(]\d+[)]`, "time.Time"},
	{`^(geometry)[(]\d+[)]`, "[]byte"},
	{`^(set)[(][\s\S]+[)]`, "string"},
}
