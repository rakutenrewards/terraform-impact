package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	tu "github.com/RakutenReady/terraform-impact/testutils"
	_ "github.com/RakutenReady/terraform-impact/testutils/setup"
)

func TestIsDir(t *testing.T) {
	testCases := []struct {
		Paths []string
		Want  bool
	}{
		{tu.GetStates(), true},
		{tu.GetModules(), true},
		{tu.GetInexistentDirs(), false},
		{tu.GetExistentFiles(), false},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		for _, path := range testCase.Paths {
			result := IsDir(path)

			assert.Equal(testCase.Want, result, `IsDir("%v")`, path)
		}
	}
}

func TestListDirsIn(t *testing.T) {
	testCases := []struct {
		Path string
		Want []string
	}{
		{
			"test_resources", []string{
				"other",
				"terraform",
			},
		},
		{
			"test_resources/other", []string{
				"ansible",
				"docker",
			},
		},
		{
			"test_resources/terraform/gcp/modules", []string{
				"datadog",
				"db",
				"google",
				"unused_module",
			},
		},
		{
			tu.GetInexistentFiles()[0], []string{},
		},
		{
			tu.GetInexistentDirs()[0], []string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		results := ListDirsIn(testCase.Path)
		msg := fmt.Sprintf("On: ListDirsIn(\"%v\")\nActual: %v\nExpected (as map): %v", testCase.Path, results, testCase.Want)

		assert.Len(results, len(testCase.Want), msg)
		for _, result := range results {
			assert.Containsf(testCase.Want, result, msg)
		}
	}
}

func TestListFilesIn(t *testing.T) {
	testCases := []struct {
		Path string
		Want []string
	}{
		{
			"test_resources/other", []string{
				"ardita.json",
				"bob.json",
				"charlotte.md",
			},
		},
		{
			"test_resources/other/ansible", []string{
				"nothing.yml",
				"ubuntu-bionic.yml",
			},
		},
		{
			"test_resources/terraform/aws/states/gateway", []string{
				"gw.vault",
				"main.tf",
				"outputs.tf",
				"terraform.tf",
				"terraform.tfvars",
			},
		},
		{
			"test_resources/terraform/gcp/states/company", []string{
				"main.tf",
				"versions.tf",
			},
		},
		{
			tu.GetInexistentFiles()[0], []string{},
		},
		{
			tu.GetInexistentDirs()[0], []string{},
		},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		results := ListFilesIn(testCase.Path)

		assert.ElementsMatchf(results, testCase.Want, `On: ListFilesIn("%v")`, testCase.Path)
	}
}

func TestSamePathComparesCleanedPaths(t *testing.T) {
	testCases := []struct {
		Path      string
		OtherPath string
		Want      bool
	}{
		// same
		{"a/b/c/", "a/b/c", true},
		{"./a/b/c/", "a/b/c", true},
		{"./a/b/../b/c", "a/b/c", true},
		{"./a/b/../b/c", "a/b/c", true},
		{"a/b/c/d.json", "a/b/c/d.json", true},
		{"./a/b/c/d.json", "a/b/c/d.json", true},
		{"./a/b/../b/c/../../b/c/d.json", "a/b/c/d.json", true},
		// different
		{"a/b/c/", "d/e/f", false},
		{"a/b/c/x.json", "a/b/c/d.json", false},
		{"a/b/../c/d.json", "a/b/c/d.json", false},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result1 := SamePath(testCase.Path, testCase.OtherPath)
		result2 := SamePath(testCase.OtherPath, testCase.Path)

		assert.Equalf(testCase.Want, result1, `On SamePath("%v", "%v")`, testCase.Path, testCase.OtherPath)
		assert.Equalf(testCase.Want, result2, `On SamePath("%v", "%v")`, testCase.OtherPath, testCase.Path)
	}
}

func TestTraceSymlink(t *testing.T) {
	testCases := []struct {
		Path string
		Want []string
	}{
		{
			"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tf", []string{
				"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tf",
				"test_resources/terraform/gcp/states/global_vars/terraform.tf",
				"test_resources/terraform/gcp/global/terraform.tf",
			},
		},
		{
			"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tfvars", []string{
				"test_resources/terraform/gcp/states/datadog-pg-google-service/terraform.tfvars",
				"test_resources/terraform/gcp/states/global_vars/terraform.tfvars",
				"test_resources/terraform/gcp/global/terraform.tfvars",
			},
		},
		{
			"test_resources/terraform/gcp/states/global_vars/terraform.tf", []string{
				"test_resources/terraform/gcp/states/global_vars/terraform.tf",
				"test_resources/terraform/gcp/global/terraform.tf",
			},
		},
		{
			"test_resources/terraform/gcp/states/global_vars/terraform.tfvars", []string{
				"test_resources/terraform/gcp/states/global_vars/terraform.tfvars",
				"test_resources/terraform/gcp/global/terraform.tfvars",
			},
		},
		{
			"test_resources/terraform/gcp/global/terraform.tf", []string{
				"test_resources/terraform/gcp/global/terraform.tf",
			},
		},
		{
			"test_resources/terraform/gcp/global/terraform.tfvars", []string{
				"test_resources/terraform/gcp/global/terraform.tfvars",
			},
		},
		{
			"test_resources/other/ardita.json", []string{
				"test_resources/other/ardita.json",
			},
		},
		{"test_resources/terraform/gcp", []string{}},
		{"other/dont/exist", []string{}},
		{"other/dont_exist.json", []string{}},
	}

	assert := assert.New(t)
	for _, testCase := range testCases {
		result := TraceSymlinkFile(testCase.Path)

		assert.ElementsMatchf(testCase.Want, result, `On: TraceSymlinkFiles("%v")`, testCase.Path)
	}
}
