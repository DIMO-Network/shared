package main

import (
	"bytes"
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

type about struct {
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
	"uint16":   "uint16",
	"float":    "float64",
	"int16":    "int16",
	"double":   "float64",
	"int32":    "int32",
	"uint8":    "uint8",
	"string[]": "[]string",
	"uint8[]":  "[]uint8",
	"int8":     "int8",
	"uint32":   "uint32",
}

func main() {

	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	} else {
		log.Fatal(errors.New("yaml file required"))
	}

	result := make(map[string]about)

	yamlFile, err := ioutil.ReadFile(arg)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		log.Fatal(err)
	}

	orderedKeys := make([]string, len(result))

	for k := range result {
		orderedKeys = append(orderedKeys, k)
	}

	sort.Strings(orderedKeys)

	output := ""
	currentSubStruct := ""
	prevLen := 0
	parentArray := []string{"Vehicle"}

	output = "package shared\n\ntype DataSchemaStruct struct {\n"

	for _, k := range orderedKeys {

		abt := result[k]
		attrs := strings.Split(k, ".")
		currentSubStruct = attrs[len(attrs)-1]

		if len(attrs) < prevLen {
			diff := prevLen - len(attrs)
			for i := 0; i < diff; i++ {
				tag := makeJSONTag(parentArray[len(parentArray)-(i+1)])
				output += fmt.Sprintf("} %s \n", tag)
			}
			parentArray = parentArray[:len(parentArray)-diff]
		}

		if parentArray[len(parentArray)-1] != currentSubStruct && abt.Type == "branch" {
			parentArray = append(parentArray, currentSubStruct)
		}

		prevLen = len(attrs)
		switch abt.Type {
		case "branch":
			currentSubStruct = attrs[len(attrs)-1]
			output += "// " + abt.Description
			output += "\n" + currentSubStruct + " struct {" + "\n"

		case "attribute", "sensor", "actuator":
			description := "// " + abt.Description + "\n"
			varAndType := currentSubStruct + " " + dataTypes[abt.DataType]
			tag := makeJSONTag(currentSubStruct)
			output += description + varAndType + tag + "\n"
		default:
			if abt.Type != "" {
				log.Println("unrecognized type: ", abt)
			}
		}
	}

	output += "\n}}"
	// output += fmt.Sprintf("\n} %s }", parentArray[len(parentArray)-1])
	f, err := os.Create("data_schema_struct.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(output)
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("go", "fmt", "dataSchema.go").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func makeJSONTag(varName string) string {
	return " `json:" + fmt.Sprintf(`"%s,omitempty"`, lowerCaseFirstLetter(varName)) + "` "
}

func lowerCaseFirstLetter(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	bts := []byte(s)
	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]
	return string(bytes.Join([][]byte{lc, rest}, nil))
}
