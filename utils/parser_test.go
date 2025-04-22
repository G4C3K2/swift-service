package utils_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/G4C3K2/swift-service/utils"
)

func TestParseCSV_ValidFile(t *testing.T) {
	content := `SWIFT CODE,COUNTRY ISO2 CODE,NAME
AAA,AA,Alpha
BBB,BB,Beta
CCC,CC,Gamma
`
	tmpFile, err := os.CreateTemp("", "test_valid.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	expected := []map[string]string{
		{"SWIFT CODE": "AAA", "COUNTRY ISO2 CODE": "AA", "NAME": "Alpha"},
		{"SWIFT CODE": "BBB", "COUNTRY ISO2 CODE": "BB", "NAME": "Beta"},
		{"SWIFT CODE": "CCC", "COUNTRY ISO2 CODE": "CC", "NAME": "Gamma"},
	}

	result, err := utils.ParseCSV(tmpFile.Name())
	if err != nil {
		t.Errorf("ParseCSV returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseCSV returned incorrect data. Got: %+v, Expected: %+v", result, expected)
	}
}

func TestParseCSV_EmptyFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "test_empty.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	result, err := utils.ParseCSV(tmpFile.Name())
	if err != nil {
		t.Errorf("ParseCSV returned an error for an empty file: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ParseCSV should return an empty slice for an empty file. Got: %+v", result)
	}
}

func TestParseCSV_FileWithHeadersOnly(t *testing.T) {
	content := `SWIFT CODE,COUNTRY ISO2 CODE,NAME
`
	tmpFile, err := os.CreateTemp("", "test_headers_only.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	result, err := utils.ParseCSV(tmpFile.Name())
	if err != nil {
		t.Errorf("ParseCSV returned an error for a file with only headers: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ParseCSV should return an empty slice for a file with only headers. Got: %+v", result)
	}
}

func TestParseCSV_MissingColumn(t *testing.T) {
	content := `SWIFT CODE,NAME
AAA,Alpha
BBB,Beta
`
	tmpFile, err := os.CreateTemp("", "test_missing_column.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	tmpFile.Close()

	expected := []map[string]string{
		{"SWIFT CODE": "AAA", "NAME": "Alpha"},
		{"SWIFT CODE": "BBB", "NAME": "Beta"},
	}
	result, err := utils.ParseCSV(tmpFile.Name())
	if err != nil {
		t.Errorf("ParseCSV returned an error for a file with a missing column: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ParseCSV returned incorrect data for a file with a missing column. Got: %+v, Expected: %+v", result, expected)
	}
}

func TestParseCSV_NonExistingFile(t *testing.T) {
	_, err := utils.ParseCSV("non_existent.csv")
	if err == nil {
		t.Errorf("ParseCSV should return an error for a non-existent file")
	}
	if !os.IsNotExist(err) {
		t.Errorf("ParseCSV returned the wrong error type for a non-existent file. Got: %v, Expected: os.ErrNotExist", err)
	}
}
