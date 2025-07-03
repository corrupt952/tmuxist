package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	test_helper "tmuxist/helper/test"
)

func TestLoadFile_YAML(t *testing.T) {
	// Create a temporary YAML file
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	yamlContent := `name: "test-session"
root: "/tmp/test"
attach: true
windows:
  - layout: "main-vertical"
    sync: false
    panes:
      - command: "echo hello"
      - command: "echo world"
  - panes:
      - command: "htop"`

	yamlPath := filepath.Join(tmpDir, "test.yaml")
	err = ioutil.WriteFile(yamlPath, []byte(yamlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the YAML file
	config, err := LoadFile(yamlPath)
	if err != nil {
		t.Fatal(err)
	}

	// Assert configuration values
	test_helper.AssertEquals(t, config.Name, "test-session")
	test_helper.AssertEquals(t, config.Root, "/tmp/test")
	test_helper.AssertEquals(t, *config.Attach, true)
	test_helper.AssertEquals(t, len(config.Windows), 2)
	test_helper.AssertEquals(t, config.Windows[0].Layout, "main-vertical")
	test_helper.AssertEquals(t, *config.Windows[0].Sync, false)
	test_helper.AssertEquals(t, len(config.Windows[0].Panes), 2)
	test_helper.AssertEquals(t, config.Windows[0].Panes[0].Command, "echo hello")
	test_helper.AssertEquals(t, config.Windows[0].Panes[1].Command, "echo world")
	test_helper.AssertEquals(t, len(config.Windows[1].Panes), 1)
	test_helper.AssertEquals(t, config.Windows[1].Panes[0].Command, "htop")
}

func TestLoadFile_YML(t *testing.T) {
	// Create a temporary YML file
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	ymlContent := `name: "yml-session"
root: "."
attach: false
windows:
  - panes:
      - command: "vim"`

	ymlPath := filepath.Join(tmpDir, "test.yml")
	err = ioutil.WriteFile(ymlPath, []byte(ymlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the YML file
	config, err := LoadFile(ymlPath)
	if err != nil {
		t.Fatal(err)
	}

	// Assert configuration values
	test_helper.AssertEquals(t, config.Name, "yml-session")
	test_helper.AssertEquals(t, config.Root, ".")
	test_helper.AssertEquals(t, *config.Attach, false)
	test_helper.AssertEquals(t, len(config.Windows), 1)
	test_helper.AssertEquals(t, len(config.Windows[0].Panes), 1)
	test_helper.AssertEquals(t, config.Windows[0].Panes[0].Command, "vim")
}

func TestLoadFile_TOML(t *testing.T) {
	// Create a temporary TOML file
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tomlContent := `name = "toml-session"
root = "/home/user"
attach = true

[[windows]]
layout = "tiled"
sync = true

[[windows.panes]]
command = "top"

[[windows.panes]]
command = "ls -la"`

	tomlPath := filepath.Join(tmpDir, "test.toml")
	err = ioutil.WriteFile(tomlPath, []byte(tomlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the TOML file
	config, err := LoadFile(tomlPath)
	if err != nil {
		t.Fatal(err)
	}

	// Assert configuration values
	test_helper.AssertEquals(t, config.Name, "toml-session")
	test_helper.AssertEquals(t, config.Root, "/home/user")
	test_helper.AssertEquals(t, *config.Attach, true)
	test_helper.AssertEquals(t, len(config.Windows), 1)
	test_helper.AssertEquals(t, config.Windows[0].Layout, "tiled")
	test_helper.AssertEquals(t, *config.Windows[0].Sync, true)
	test_helper.AssertEquals(t, len(config.Windows[0].Panes), 2)
	test_helper.AssertEquals(t, config.Windows[0].Panes[0].Command, "top")
	test_helper.AssertEquals(t, config.Windows[0].Panes[1].Command, "ls -la")
}

func TestLoadFile_NoExtension_DefaultsToTOML(t *testing.T) {
	// Create a temporary file without extension
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	tomlContent := `name = "no-ext-session"
root = "."
attach = false

[[windows]]
[[windows.panes]]
command = "bash"`

	noExtPath := filepath.Join(tmpDir, "config")
	err = ioutil.WriteFile(noExtPath, []byte(tomlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the file without extension (should default to TOML)
	config, err := LoadFile(noExtPath)
	if err != nil {
		t.Fatal(err)
	}

	// Assert configuration values
	test_helper.AssertEquals(t, config.Name, "no-ext-session")
	test_helper.AssertEquals(t, config.Root, ".")
	test_helper.AssertEquals(t, *config.Attach, false)
}

func TestConfigurationPath_Priority(t *testing.T) {
	// Create a temporary directory as working directory
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Test 1: When only TOML exists
	ioutil.WriteFile("tmuxist.toml", []byte("name = \"toml\""), 0644)
	path, err := ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.toml")

	// Test 2: When YAML exists (should take priority over TOML)
	ioutil.WriteFile("tmuxist.yaml", []byte("name: yaml"), 0644)
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.yaml")

	// Test 3: When YML also exists, YAML should still take priority
	ioutil.WriteFile("tmuxist.yml", []byte("name: yml"), 0644)
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.yaml")

	// Clean up YAML to test YML
	os.Remove("tmuxist.yaml")
	
	// Test 4: YML should be selected when YAML is removed
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.yml")

	// Test 5: When no config exists, default to TOML
	os.Remove("tmuxist.yml")
	os.Remove("tmuxist.toml")
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.toml")
}

func TestConfigurationPath_HiddenFiles(t *testing.T) {
	// Create a temporary directory as working directory
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Change to temp directory
	originalWd, _ := os.Getwd()
	os.Chdir(tmpDir)
	defer os.Chdir(originalWd)

	// Test 1: Hidden YAML file takes priority over non-hidden TOML
	ioutil.WriteFile("tmuxist.toml", []byte("name = \"toml\""), 0644)
	ioutil.WriteFile(".tmuxist.yaml", []byte("name: hidden-yaml"), 0644)
	path, err := ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), ".tmuxist.yaml")

	// Test 2: Non-hidden file takes priority over hidden file of same type
	ioutil.WriteFile("tmuxist.yaml", []byte("name: yaml"), 0644)
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), "tmuxist.yaml")

	// Clean up non-hidden yaml
	os.Remove("tmuxist.yaml")

	// Test 3: Hidden YAML takes priority when non-hidden YAML doesn't exist
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), ".tmuxist.yaml")

	// Test 4: Hidden YML file
	ioutil.WriteFile(".tmuxist.yml", []byte("name: hidden-yml"), 0644)
	os.Remove(".tmuxist.yaml")
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), ".tmuxist.yml")

	// Test 5: Hidden TOML file
	ioutil.WriteFile(".tmuxist.toml", []byte("name = \"hidden-toml\""), 0644)
	os.Remove(".tmuxist.yml")
	os.Remove("tmuxist.toml")
	path, err = ConfigurationPath()
	if err != nil {
		t.Fatal(err)
	}
	test_helper.AssertEquals(t, filepath.Base(path), ".tmuxist.toml")
}

func TestLoadFile_HiddenFiles(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Test loading hidden YAML file
	hiddenYamlContent := `name: "hidden-session"
root: "/tmp/hidden"
attach: false
windows:
  - panes:
      - command: "ls -la"`

	hiddenYamlPath := filepath.Join(tmpDir, ".tmuxist.yaml")
	err = ioutil.WriteFile(hiddenYamlPath, []byte(hiddenYamlContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the hidden YAML file
	config, err := LoadFile(hiddenYamlPath)
	if err != nil {
		t.Fatal(err)
	}

	// Assert configuration values
	test_helper.AssertEquals(t, config.Name, "hidden-session")
	test_helper.AssertEquals(t, config.Root, "/tmp/hidden")
	test_helper.AssertEquals(t, *config.Attach, false)
	test_helper.AssertEquals(t, len(config.Windows), 1)
	test_helper.AssertEquals(t, config.Windows[0].Panes[0].Command, "ls -la")
}

func TestLoadFile_InvalidYAML(t *testing.T) {
	// Create a temporary invalid YAML file
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	invalidYaml := `name: "test"
  invalid indentation here
windows:`

	yamlPath := filepath.Join(tmpDir, "invalid.yaml")
	err = ioutil.WriteFile(yamlPath, []byte(invalidYaml), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Try to load the invalid YAML file
	_, err = LoadFile(yamlPath)
	if err == nil {
		t.Fatal("Expected error for invalid YAML, but got nil")
	}
}

func TestLoadFile_InvalidTOML(t *testing.T) {
	// Create a temporary invalid TOML file
	tmpDir, err := ioutil.TempDir("", "tmuxist-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	invalidToml := `name = "test"
[invalid section without closing
windows = []`

	tomlPath := filepath.Join(tmpDir, "invalid.toml")
	err = ioutil.WriteFile(tomlPath, []byte(invalidToml), 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Try to load the invalid TOML file
	_, err = LoadFile(tomlPath)
	if err == nil {
		t.Fatal("Expected error for invalid TOML, but got nil")
	}
}