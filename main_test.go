package main

import (
	"testing"
	"time"
)

func TestExecuteCommandWithTimeout(t *testing.T) {
	// Test successful command
	result, err := executeCommandWithTimeout("echo 'test'", 5*time.Second)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result != "test" {
		t.Errorf("Expected 'test', got '%s'", result)
	}

	// Test timeout
	_, err = executeCommandWithTimeout("sleep 10", 1*time.Second)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

func TestCheckDependency(t *testing.T) {
	// Test environment variable check
	status, version := checkDependency("echo $PATH")
	if status == "Missing" {
		t.Error("Expected PATH to be found")
	}
	if version == "" {
		t.Error("Expected PATH value, got empty string")
	}

	// Test missing command
	status, version = checkDependency("nonexistentcommand12345")
	if status != "Missing" {
		t.Errorf("Expected 'Missing', got '%s'", status)
	}
}

func TestGetDependencies(t *testing.T) {
	deps := getDependencies()
	if len(deps) == 0 {
		t.Error("Expected dependencies, got empty slice")
	}

	// Check that we have core dependencies
	hasGo := false
	hasFyne := false
	for _, dep := range deps {
		if dep.Name == "Go" {
			hasGo = true
		}
		if dep.Name == "Fyne CLI" {
			hasFyne = true
		}
	}

	if !hasGo {
		t.Error("Expected Go dependency not found")
	}
	if !hasFyne {
		t.Error("Expected Fyne CLI dependency not found")
	}
}

func TestConfig(t *testing.T) {
	config := Config{
		Verbose:    true,
		JSON:       false,
		Categories: []string{"core"},
		Timeout:    5 * time.Second,
		OutputFile: "test.txt",
	}

	if !config.Verbose {
		t.Error("Expected Verbose to be true")
	}
	if config.JSON {
		t.Error("Expected JSON to be false")
	}
	if len(config.Categories) != 1 {
		t.Error("Expected 1 category")
	}
	if config.Timeout != 5*time.Second {
		t.Error("Expected timeout to be 5 seconds")
	}
	if config.OutputFile != "test.txt" {
		t.Error("Expected output file to be test.txt")
	}
}
