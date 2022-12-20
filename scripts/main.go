package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

type fields struct {
	Allowed     []string `yaml:"allowed"`
	Comment     string   `yaml:"comment"`
	DataType    string   `yaml:"datatype"`
	Type        string   `yaml:"type"`
	UUID        string   `yaml:"uuid"`
	Description string   `yaml:"description"`
}

var dataTypes = map[string]string{
	"boolean":  "bool",
	"string":   "string",
	"string[]": "[]string",
	"uint8[]":  "[]uint8",
	"uint8":    "uint8",
	"int8":     "int8",
	"uint16":   "uint16",
	"int16":    "int16",
	"uint32":   "uint32",
	"int32":    "int32",
	"float":    "float64",
	"double":   "float64",
}

func getDataFromFile(arg string) map[string]fields {
	data := make(map[string]fields)
	yamlFile, err := ioutil.ReadFile(arg)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func writeToFile(fileName, data string) {

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
	_, err = exec.Command("go", "fmt", fileName).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(errors.New("yaml file required"))
	}
	arg := os.Args[1]

	data := getDataFromFile(arg)

	orderedKeys := make([]string, len(data))
	for k := range data {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Strings(orderedKeys)

	headers := "package shared\n\nimport(\n\"fmt\"\n\"github.com/clarketm/json\"\n)\n\n"
	dataStruct := generateStruct(orderedKeys, data)
	validChecks := validEnumCheck(orderedKeys, data)
	methods := generateMethods()
	finalOutput := headers + validChecks + dataStruct + methods
	writeToFile("data_schema_struct.go", finalOutput)
}

func makeJSONTag(varName string) string {
	return fmt.Sprintf(" `json:\"%s,omitempty\"`", lowerFirstLetter(varName))
}

func lowerFirstLetter(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}

func validEnumCheck(data []string, orderedKeys map[string]fields) string {
	isvalidchecks := `func(d DataSchemaStruct) IsValid() (bool, error){`

	n := 0
	for _, kvalidate := range data {
		fields := orderedKeys[kvalidate]

		if fields.Allowed != nil && !strings.HasSuffix(fields.DataType, "[]") {
			if n == 0 {
				isvalidchecks += "\nvalid := []string{"

				for i, l := range fields.Allowed {
					isvalidchecks += fmt.Sprintf("\"%s\"", l)
					if i != len(fields.Allowed)-1 {
						isvalidchecks += ","
					}
				}
				isvalidchecks += ",\"\"}\n\n"

				isvalidchecks += "val := d." + kvalidate
			} else {
				isvalidchecks += "\nvalid = []string{"

				for i, l := range fields.Allowed {
					isvalidchecks += fmt.Sprintf("\"%s\"", l)
					if i != len(fields.Allowed)-1 {
						isvalidchecks += ","
					}
				}
				isvalidchecks += ",\"\"}\n\nval = d." + kvalidate
			}
			isvalidchecks += fmt.Sprintf("\nif !contains(valid, val) {\nreturn false, fmt.Errorf(\"invalid value at %s\", val)\n}\n ", kvalidate+": %v")

			n++
		}

	}

	isvalidchecks += "\nreturn true, nil}\n\n"
	return isvalidchecks
}

func generateMethods() string {
	methods := `
	func NewVehicleStatus() DataSchemaStruct {
		return DataSchemaStruct{}
	}

	func (d DataSchemaStruct) Marshal() ([]byte, error) {
		if valid, err := d.IsValid(); !valid {
			return []byte{}, err
		}

		return json.Marshal(d)
	}

	func contains(s []string, str string) bool {
		for _, v := range s {
			if v == str {
				return true
			}
		}
		return false
	}
	`

	return methods
}

func generateStruct(orderedKeys []string, data map[string]fields) string {

	currentSubStruct := ""
	prevLen := 0
	parentArray := []string{"Vehicle"}
	dataStruct := "type DataSchemaStruct struct {\n"

	for _, k := range orderedKeys {
		fields := data[k]
		attrs := strings.Split(k, ".")
		currentSubStruct = attrs[len(attrs)-1]

		if len(attrs) < prevLen {
			diff := prevLen - len(attrs)
			for i := 0; i < diff; i++ {
				tag := makeJSONTag(parentArray[len(parentArray)-(i+1)])
				dataStruct += fmt.Sprintf("} %s \n", tag)
			}
			parentArray = parentArray[:len(parentArray)-diff]
		}
		if parentArray[len(parentArray)-1] != currentSubStruct && fields.Type == "branch" {
			parentArray = append(parentArray, currentSubStruct)
		}
		prevLen = len(attrs)
		switch fields.Type {
		case "branch":
			currentSubStruct = attrs[len(attrs)-1]
			dataStruct += "// " + fields.Description
			dataStruct += "\n" + currentSubStruct + " struct {" + "\n"
		case "attribute", "sensor", "actuator":
			description := "// " + fields.Description
			varAndType := "\n" + currentSubStruct + " " + dataTypes[fields.DataType]
			if len(fields.Allowed) != 0 {
				description += fmt.Sprintf(" Must be one of [\"%s\"]", strings.Join(fields.Allowed, `", "`))
			}
			tag := makeJSONTag(currentSubStruct)
			dataStruct += description + varAndType + tag + "\n"

		default:
			if fields.Type != "" {
				log.Println("unrecognized type: ", fields)
			}
		}
	}

	if parentArray != nil {
		tag := makeJSONTag(parentArray[0])
		dataStruct += fmt.Sprintf("\n} %s }", tag)
	}

	return dataStruct

}
