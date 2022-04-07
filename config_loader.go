package shared

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

// LoadConfig fills in all the values in the Settings from local yml file (for dev) and env vars (for deployments)
func LoadConfig[S any](filePath string) (S, error) {
	file, err := ioutil.ReadFile(filePath)
	var settings S
	// if no file found, ignore as we could be running in higher level environment. We could make this more explicit with a cli parameter w/ the filename
	if err != nil {
		fmt.Printf("looks like running in higher level env as could not read file: %s \n", filePath)
	} else {
		settings, err = loadFromYaml[S](file)
		if err != nil {
			return settings, errors.Wrap(err, "could not load yaml")
		}
	}
	err = loadFromEnvVars(&settings) // override with any env vars found

	return settings, err
}

// loadFromYaml loads yaml from text into object per type. Uses reflection to find yaml type metadata for lookup
func loadFromYaml[S any](yamlFile []byte) (S, error) {
	var settings S
	err := yaml.Unmarshal(yamlFile, &settings)
	if err != nil {
		return settings, errors.Wrap(err, "could not unmarshall yaml file to settings")
	}
	return settings, nil
}

// loadFromEnvVars uses reflection to find yaml type metadata and lookup those values in environment variables for corresponding property.
// settings must be passed in as a pointer
func loadFromEnvVars[S any](settings S) error {
	if reflect.ValueOf(settings).Kind() != reflect.Ptr {
		return errors.New("settings must be passed in as a pointer, eg. &varname")
	}
	valueOfConfig := reflect.ValueOf(settings).Elem()
	typeOfT := valueOfConfig.Type()
	var err error
	// iterate over all struct fields
	for i := 0; i < valueOfConfig.NumField(); i++ {
		field := valueOfConfig.Field(i)
		fieldYamlName := typeOfT.Field(i).Tag.Get("yaml")

		// check if env var with same field yaml name exists, if so, set the value from the env var
		if env, exists := os.LookupEnv(fieldYamlName); exists {
			var val interface{}
			switch field.Kind() {
			case reflect.String:
				val = env
			case reflect.Bool:
				val, err = strconv.ParseBool(env)
			case reflect.Int:
				val, err = strconv.Atoi(env)
			case reflect.Int64:
				val, err = strconv.ParseInt(env, 10, 64)
			}
			// now set the field with the val
			if val != nil {
				field.Set(reflect.ValueOf(val))
			}
		}
	}
	return err
}
