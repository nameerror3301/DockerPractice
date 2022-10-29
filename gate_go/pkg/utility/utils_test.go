package utility

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

const TestVersion = "v0.0.1"

func TestReadConfig(t *testing.T) {
	testTable := []struct {
		name string
		args struct {
			dataConfig *ConfigModel
			dataType   string
		}
		result string
	}{
		{
			name: "JSON reader",
			args: struct {
				dataConfig *ConfigModel
				dataType   string
			}{
				dataConfig: &ConfigModel{
					Service: struct {
						Version string "json:\"version\" yaml:\"version\""
						Address string "json:\"address\" yaml:\"address\""
						Port    int    "json:\"port\" yaml:\"port\""
					}{Version: TestVersion},
				},
				dataType: string(driverJSON),
			},
			result: TestVersion,
		},
		{
			name: "YAML reader",
			args: struct {
				dataConfig *ConfigModel
				dataType   string
			}{
				dataConfig: &ConfigModel{
					Service: struct {
						Version string "json:\"version\" yaml:\"version\""
						Address string "json:\"address\" yaml:\"address\""
						Port    int    "json:\"port\" yaml:\"port\""
					}{Version: TestVersion},
				},
				dataType: string(driverYAML),
			},
			result: TestVersion,
		},
		{
			name: "TOML reader",
			args: struct {
				dataConfig *ConfigModel
				dataType   string
			}{
				dataConfig: &ConfigModel{
					Service: struct {
						Version string "json:\"version\" yaml:\"version\""
						Address string "json:\"address\" yaml:\"address\""
						Port    int    "json:\"port\" yaml:\"port\""
					}{Version: TestVersion},
				},
				dataType: string(driverTOML),
			},
			result: TestVersion,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			if err := os.Setenv(ENV_DriverConfig, testCase.args.dataType); err != nil {
				t.Fatal(err.Error())
			}

			testFile, err := os.CreateTemp("", "")
			if err != nil {
				t.Fatalf("Failed create test file %v", err)
			}

			os.Setenv("PATH_CONFIG", testFile.Name())

			// Create tmp file
			if testCase.args.dataType == string(driverJSON) || testCase.args.dataType == string(driverYAML) {
				data, err := json.Marshal(testCase.args.dataConfig)
				if err != nil {
					t.Fatalf(" %v", err)
				}

				if _, err = fmt.Fprintf(testFile, "%v", string(data)); err != nil {
					t.Fatalf("Failed write test file %v", err)
				}

				if err := testFile.Close(); err != nil {
					t.Fatalf("Failed close test file %v", err)
				}

			} else if testCase.args.dataType == string(driverTOML) {
				tmpConfig := &ConfigModel{}
				tmpConfig.Service.Version = TestVersion

				err := toml.NewEncoder(testFile).Encode(tmpConfig)
				if err != nil {
					t.Fatalf("Fatal encode %v", err)
				}

			}

			// Read tmp file
			config, err := ReadConfig()
			if err != nil {
				t.Fatalf("Filed ReadConfig %v", err)
			}

			if config.Service.Version != TestVersion {
				t.Fatal("Readed data from test file is invalid")
			}
		})
	}
}

func TestGetEnv(t *testing.T) {
	if err := os.Setenv(ENV_DriverConfig, "yaml"); err != nil {
		t.Fatal(err.Error())
	}

	driver := os.Getenv(ENV_DriverConfig)
	if driver != "yaml" {
		t.Fatal("Failed get .env driver -", driver)
	}
}
