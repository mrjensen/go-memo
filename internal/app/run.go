package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mrjensen/go-memo/internal/app/format"
)

const (
	inputDirectory  = "templates"
	outputDirectory = "memos"
)

var outputMarkdownFiles []string

func Run() {
	outputMarkdownFiles = make([]string, 0)

	err := filepath.Walk(inputDirectory, outputMarkdown)
	exitOnError(err)

	generateTableOfContents()
}

func outputMarkdown(inputPath string, info os.FileInfo, err error) error {
	if info.IsDir() || !strings.HasSuffix(inputPath, ".go") {
		return nil
	}

	introPath := strings.ReplaceAll(inputPath, ".go", ".md")
	introOutput, err := ioutil.ReadFile(introPath)
	if err != nil && !os.IsNotExist(err) {
		exitOnError(err)
	}

	segments := strings.Split(inputPath, string(os.PathSeparator))
	segments[0] = outputDirectory
	outputPath := strings.Join(segments, string(os.PathSeparator))
	outputPath = strings.ReplaceAll(outputPath, ".go", ".md")
	outputMarkdownFiles = append(outputMarkdownFiles, outputPath)

	goPath, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	cmd := exec.Command(goPath, "run", inputPath)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	markdown := []string{
		string(introOutput),
		"\n```\n",
		string(output),
		"\n```\n\n",
		fmt.Sprintf("See file: [%[1]s](%[1]s)", fmt.Sprintf("/../../blob/master/%s", inputPath)),
	}
	data := strings.Join(markdown, "") // fmt.Sprintf("See file: %s", introOutput, output, )

	err = ioutil.WriteFile(outputPath, []byte(data), 0644)
	if err != nil && os.IsNotExist(err) {
		pathParts := strings.Split(outputPath, string(os.PathSeparator))
		directory := strings.Join(pathParts[:len(pathParts)-1], string(os.PathSeparator))
		err = os.MkdirAll(directory, 0755)
		err = ioutil.WriteFile(outputPath, []byte(data), 0644)
	}
	exitOnError(err)

	return nil
}

func generateTableOfContents() {
	f, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY, 0644)
	exitOnError(err)
	defer func() {
		err = f.Close()
		exitOnError(err)
	}()

	intro, err := ioutil.ReadFile("internal/app/README_intro.md")
	exitOnError(err)
	_, err = f.Write(intro)
	exitOnError(err)
	_, err = f.WriteString(format.TableOfContentsFromFilePaths(outputMarkdownFiles))
	exitOnError(err)
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println("failed with error:", err.Error())
		os.Exit(1)
	}
}
