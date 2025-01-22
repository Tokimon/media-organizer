package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"media-organizer/backend/colors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func contentHash(content string) string {
	hasher := sha1.New()
	hasher.Write([]byte(content))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func extractSVGContent(svg string) (string, string) {
	svgRe := regexp.MustCompile(`^(?:.|\n|\r)*<svg.*?(viewBox=".*?").*?>[^<]*|(\s|\n|\r)*</svg>(?:.|\n|\r)*$`)
	spcRe := regexp.MustCompile(`>[^<]*<`)

	viewBox := svgRe.FindStringSubmatch(svg)

	svg = svgRe.ReplaceAllString(svg, "")
	svg = spcRe.ReplaceAllString(svg, "><")

	return svg, viewBox[1]
}

func removeDir(path string) {
	err := os.RemoveAll(path)

	if err != nil {
		fmt.Println(fmt.Errorf("%s Failed to remove dir: %s\n%w", colors.Red("⨉"), colors.Gray(path), err))
	} else {
		fmt.Println(fmt.Sprintf("%s Removed: %s", colors.Green("✓"), colors.Strikethrough(colors.Gray(path))))
	}
}

func writeFile(path string, content string) {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		fmt.Println(fmt.Errorf("%s Failed to create path: %s\n%w", colors.Red("⨉"), dir, err))
		return
	}

	err = os.WriteFile(path, []byte(content), 0644)

	if err != nil {
		fmt.Println(fmt.Errorf("%s Failed to write file: %s\n%w", colors.Red("⨉"), colors.Gray(path), err))
	} else {
		fmt.Println(fmt.Sprintf("%s Written: %s", colors.Green("✓"), colors.Gray(path)))
	}
}

func main() {
	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		fmt.Println(fmt.Errorf(`Failed to get CWD: %w`, cwdErr))
		return
	}

	pathIn := filepath.Join(cwd, os.Args[1])
	svgGlob := filepath.Join(pathIn, "*.svg")

	fmt.Println("Searching for icons:")
	fmt.Println("--------------------")
	fmt.Println(colors.Yellow(pathIn))
	fmt.Println("")

	files, err := filepath.Glob(svgGlob)

	if err != nil {
		fmt.Println(fmt.Errorf("Failed to find files: %s\n%w", svgGlob, err))
		return
	}

	if len(files) == 0 {
		fmt.Println(fmt.Sprintf(`%s Found no svg files at "%s"`, colors.Red("×"), pathIn))
		return
	}

	fmt.Println("Parsing icons:")
	fmt.Println("--------------")

	var names []string
	var symbols []string
	var uses []string

	for _, filePath := range files {
		bytes, err := os.ReadFile(filePath)

		if err != nil {
			fmt.Println(fmt.Errorf(`Failed to read file: %s\n%w`, filePath, err))
			continue
		}

		filename := strings.TrimSuffix(filepath.Base(filePath), ".svg")
		names = append(names, filename)

		svgContent, viewBox := extractSVGContent(string(bytes))

		symbols = append(symbols, fmt.Sprintf("<symbol id=\"sym-%s\" %s>%s</symbol>", filename, viewBox, svgContent))
		uses = append(uses, fmt.Sprintf("<use id=\"%s\" href=\"#sym-%s\" />", filename, filename))

		fmt.Println("-", colors.Yellow(filename))
	}

	tsTmpl := `export const iconNames = [
	'%s'
] as const;

export type IconNames = typeof iconNames[number];`

	svgTmpl := `<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
	<style>use:not(:target) { display: none; }</style>
	<defs>
		%s
	</defs>
	%s
</svg>`

	cssTmpl := `[class*="icon-"] {
	display: inline-block;
	width: 2.4rem;
	background: currentColor;
	aspect-ratio: 1 / 1;
	mask-size: contain;
	mask-repeat: no-repeat;
	mask-position: center;
}

%s`

	fmt.Println("\nWriting files:")
	fmt.Println("--------------")

	pathOut := filepath.Join(cwd, os.Args[2])

	svgContent := fmt.Sprintf(svgTmpl, strings.Join(symbols, "\n\t\t"), strings.Join(uses, "\n\t"))
	svgHash := contentHash(svgContent)
	svgFileName := fmt.Sprintf("icons-%s.svg", svgHash)

	var iconLines []string

	for _, name := range names {
		iconLines = append(iconLines, fmt.Sprintf(".icon-%s { mask-image: url('./%s#%s'); }", name, svgFileName, name))
	}

	cssContent := fmt.Sprintf(cssTmpl, strings.Join(iconLines, "\n"))

	removeDir(pathOut)

	writeFile(filepath.Join(pathOut, svgFileName), svgContent)
	writeFile(filepath.Join(pathOut, "icons.css"), cssContent)
	writeFile(filepath.Join(pathOut, "icons.ts"), fmt.Sprintf(tsTmpl, strings.Join(names, "',\n\t'")))

	fmt.Println("")
}
