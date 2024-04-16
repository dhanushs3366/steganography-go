package pyscripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func runPythonScript(scriptName string, args ...string) (string, error) {
	// Get the absolute path to the directory containing the Python script
	scriptDir, err := filepath.Abs("/home/munke/university_works/projects/steganography/stegano-py/")
	if err != nil {
		return "", fmt.Errorf("error getting absolute path: %v", err)
	}

	// Set the working directory to the directory containing the Python script
	err = os.Chdir(scriptDir)
	if err != nil {
		return "", fmt.Errorf("error setting working directory: %v", err)
	}

	// Command to call the Python script with arguments
	cmd := exec.Command("python3.10", append([]string{scriptName}, args...)...)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error running Python script: %v", err)
	}

	return string(output), nil
}

func PyEncode(imagePath, text string) (string, error) {
	println(imagePath)
	return runPythonScript("encode.py", imagePath, text)
}

func PyDecode(imagePath string) (string, error) {
	return runPythonScript("decode.py", imagePath)
}
