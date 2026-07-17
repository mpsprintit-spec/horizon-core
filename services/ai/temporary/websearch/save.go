package websearch

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func SaveKnowledge(result *WikipediaResult) error {

	fmt.Println()
	fmt.Println("========================================================")
	fmt.Println("HORIZON SAVE KNOWLEDGE")
	fmt.Println("========================================================")
	fmt.Println()

	// ======================================================
	// Build YAML
	// ======================================================

	var yaml strings.Builder

	yaml.WriteString("# ==============================================\n")
	yaml.WriteString("# Horizon Dictionary\n")
	yaml.WriteString("# Generated Automatically\n")
	yaml.WriteString("# ==============================================\n\n")

	yaml.WriteString("lemma: ")
	yaml.WriteString(result.Word)
	yaml.WriteString("\n")

	yaml.WriteString("title: ")
	yaml.WriteString(result.Title)
	yaml.WriteString("\n")

	yaml.WriteString("description: |\n")

	for _, line := range strings.Split(result.Description, "\n") {

		yaml.WriteString("  ")
		yaml.WriteString(strings.TrimSpace(line))
		yaml.WriteString("\n")

	}

	yaml.WriteString("\n")

	yaml.WriteString("meaning: |\n")

	for _, line := range strings.Split(result.Summary, "\n") {

		yaml.WriteString("  ")
		yaml.WriteString(strings.TrimSpace(line))
		yaml.WriteString("\n")

	}

	yaml.WriteString("\n")

	yaml.WriteString("language: Indonesian\n")
	yaml.WriteString("status: verified\n")

	yaml.WriteString("source:\n")
	yaml.WriteString("  name: Wikipedia Indonesia\n")
	yaml.WriteString("  url: ")
	yaml.WriteString(result.URL)
	yaml.WriteString("\n")

	yaml.WriteString("related_words: []\n")
	yaml.WriteString("examples: []\n")
	yaml.WriteString("cross_references: []\n")

	// ======================================================
	// Create Folder
	// ======================================================

	directory := filepath.Join(
		"knowledge",
		"language",
		"Indonesian",
		"dictionary",
	)

	if err := os.MkdirAll(directory, 0755); err != nil {
		return err
	}

	filename := filepath.Join(
		directory,
		result.Word+".yaml",
	)

	fmt.Println("Directory : ", directory)
	fmt.Println("File      : ", filename)
	fmt.Println()
  
  // ======================================================
	// Save YAML File
	// ======================================================

	fmt.Println("Writing dictionary file...")

	if err := os.WriteFile(filename, []byte(yaml.String()), 0644); err != nil {
		return err
	}

	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	fmt.Println("Dictionary successfully saved.")
	fmt.Println("Size :", info.Size(), "bytes")
	fmt.Println()

	// ======================================================
	// Git Add
	// ======================================================

	fmt.Println("Running Git Add...")

	cmd := exec.Command("git", "add", ".")

	output, err := cmd.CombinedOutput()

	if err != nil {

		fmt.Println(string(output))

		return err

	}

	fmt.Println("Git Add Success")
	fmt.Println()

	// ======================================================
	// Git Commit
	// ======================================================

	fmt.Println("Running Git Commit...")

	message := "Knowledge: " + result.Word

	cmd = exec.Command(
		"git",
		"commit",
		"-m",
		message,
	)

	output, err = cmd.CombinedOutput()

	if err != nil {

		fmt.Println(string(output))

		return err

	}

	fmt.Println("Git Commit Success")
	fmt.Println()

	// ======================================================
	// Git Push
	// ======================================================

	fmt.Println("Running Git Push...")

	cmd = exec.Command(
		"git",
		"push",
	)

	output, err = cmd.CombinedOutput()

	if err != nil {

		fmt.Println(string(output))

		return err

	}

	fmt.Println("Git Push Success")
	fmt.Println()

  	// ======================================================
	// Finish
	// ======================================================

	fmt.Println("========================================================")
	fmt.Println("HORIZON KNOWLEDGE SAVED")
	fmt.Println("========================================================")
	fmt.Println()

	fmt.Println("Word")
	fmt.Println("----------------------------------------")
	fmt.Println(result.Word)
	fmt.Println()

	fmt.Println("Title")
	fmt.Println("----------------------------------------")
	fmt.Println(result.Title)
	fmt.Println()

	fmt.Println("Location")
	fmt.Println("----------------------------------------")
	fmt.Println(filename)
	fmt.Println()

	fmt.Println("Source")
	fmt.Println("----------------------------------------")
	fmt.Println(result.URL)
	fmt.Println()

	fmt.Println("Git Status")
	fmt.Println("----------------------------------------")
	fmt.Println("✓ Git Add")
	fmt.Println("✓ Git Commit")
	fmt.Println("✓ Git Push")
	fmt.Println()

	fmt.Println("Status")
	fmt.Println("----------------------------------------")
	fmt.Println("Knowledge successfully added to Horizon.")
	fmt.Println()

	fmt.Println("========================================================")
	fmt.Println("END OF SAVE PROCESS")
	fmt.Println("========================================================")
	fmt.Println()

	return nil

}
