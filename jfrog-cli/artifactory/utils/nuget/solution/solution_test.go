package solution

import (
	"encoding/json"
	"fmt"
	"github.com/jfrog/jfrog-client-go/artifactory/buildinfo"
	"github.com/jfrog/jfrog-client-go/utils"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestEmptySolution(t *testing.T) {
	solution, err := Load(".", "")
	if err != nil {
		t.Error(err)
	}

	expected := &buildinfo.BuildInfo{}
	buildInfo, err := solution.BuildInfo()
	if err != nil {
		t.Error("An error occurred while creating the build info object", err.Error())
	}
	if !reflect.DeepEqual(buildInfo, expected) {
		expectedString, _ := json.Marshal(expected)
		buildInfoString, _ := json.Marshal(buildInfo)
		t.Errorf("Expecting: \n%s \nGot: \n%s", expectedString, buildInfoString)
	}
}

func TestParseSln(t *testing.T) {
	regExp, err := utils.GetRegExp(`Project\("(.*)\nEndProject`)
	if err != nil {
		t.Error(err)
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	testsDataDir := filepath.Join(pwd, "testsdata")

	tests := []struct {
		name     string
		slnPath  string
		expected []string
	}{
		{"oneproject", filepath.Join(testsDataDir, "oneproject.sln"), []string{`Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfig", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`}},
		{"multiProjects", filepath.Join(testsDataDir, "multiprojects.sln"), []string{`Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfigmulti", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfiganothermulti", "test\packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := parseSlnFile(test.slnPath, regExp)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(test.expected, result) {
				t.Error(fmt.Sprintf("Expected %s, got %s", test.expected, result))
			}
		})
	}
}

func TestParseProject(t *testing.T) {

	tests := []struct {
		name                string
		projectLine         string
		expectedCsprojPath  string
		expectedProjectName string
	}{
		{"packagename", `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagename", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, "jfrog/path/test/packagesconfig.csproj", "packagename"},
		{"withpath", `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagename", "packagesconfig/packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, "jfrog/path/test/packagesconfig/packagesconfig.csproj", "packagename"},
		{"sameprojectname", `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfig", "packagesconfig/packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, "jfrog/path/test/packagesconfig/packagesconfig.csproj", "packagesconfig"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			csprojPath, projectName, err := parseProject(test.projectLine, "jfrog/path/test")
			if err != nil {
				t.Error(err)
			}
			if csprojPath != test.expectedCsprojPath {
				t.Error(fmt.Sprintf("Expected %s, got %s", test.expectedCsprojPath, csprojPath))
			}
			if projectName != test.expectedProjectName {
				t.Error(fmt.Sprintf("Expected %s, got %s", test.expectedProjectName, projectName))
			}
		})
	}
}

func TestGetProjectsFromSlns(t *testing.T) {
	regExp, err := utils.GetRegExp(`Project\("(.*)\nEndProject`)
	if err != nil {
		t.Error(err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	testsDataDir := filepath.Join(pwd, "testsdata")
	tests := []struct {
		name             string
		solution         solution
		expectedProjects []string
	}{
		{"withoutSlnFile", solution{path: testsDataDir, slnFile: "", projects: nil}, []string{`Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfigmulti", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfiganothermulti", "test\packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`, `Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfig", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`},
		},
		{"withSlnFile", solution{path: testsDataDir, slnFile: "oneproject.sln", projects: nil}, []string{`Project("{FAE04EC0-301F-11D3-BF4B-00C04F79EFBC}") = "packagesconfig", "packagesconfig.csproj", "{D1FFA0DC-0ACC-4108-ADC1-2A71122C09AF}"
EndProject`},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.solution.getProjectsFromSlns(regExp)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(test.expectedProjects, result) {
				t.Error(fmt.Sprintf("Expected %s, got %s", test.expectedProjects, result))
			}
		})
	}
}
