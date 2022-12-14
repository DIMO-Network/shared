package main

import (
	"errors"
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

	output = "package shared\n\ntype dataSchema struct {\n"

	for _, k := range orderedKeys {

		abt := result[k]
		attrs := strings.Split(k, ".")
		currentSubStruct = attrs[len(attrs)-1]

		if len(attrs) < prevLen {
			for i := 0; i < prevLen-len(attrs); i++ {
				output += "}\n"
			}
		}

		prevLen = len(attrs)
		switch abt.Type {
		case "branch":
			currentSubStruct = attrs[len(attrs)-1]
			output += "\n// " + abt.Description
			output += "\n" + currentSubStruct + " struct {" + "\n"

		case "attribute", "sensor", "actuator":
			output += "\n// " + abt.Description + "\n"
			output += currentSubStruct + " " + dataTypes[abt.DataType] + "\n"
		default:
			log.Println("unrecognized type: ", abt.Type)
		}

	}

	output += "\n}}"

	f, err := os.Create("dataSchema.go")
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
