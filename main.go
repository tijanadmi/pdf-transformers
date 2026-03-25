// package main

// import (
// 	"fmt"
// 	"os/exec"
// )

// func main() {
// 	pdfPath := "primer.pdf"
// 	outputPrefix := "slike/strana"

// 	cmd := exec.Command(
// 		"D:\\poppler\\Library\\bin\\pdftoppm.exe",
// 		"-png",
// 		"-r", "300",
// 		pdfPath,
// 		outputPrefix,
// 	)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Println("Greška:", err)
// 		fmt.Println("Detalji:", string(output))
// 		return
// 	}

// 	fmt.Println("OK:", string(output))
// }

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Nedostaju argumenti: <folder_ime> <pdf_ime>")
		return
	}

	folderIme := os.Args[1]
	pdfIme := os.Args[2]

	// basePath := "D:\\go_workspace\\pdf-transformers"
	basePath := "E:\\aplikacije\\tis\\ddn"

	// putanja do ulaznog PDF-a (folder korisnika + pdfIme)
	inputPath := filepath.Join(basePath, folderIme, "scan", pdfIme)
	// fmt.Println(inputPath)

	// folder za izlazne slike: basePath + folderIme + "slike"
	outputFolder := filepath.Join(basePath, folderIme, "scan", "slike")
	// fmt.Println(outputFolder)

	// Kreiraj output folder ako ne postoji
	os.MkdirAll(outputFolder, os.ModePerm)
	pdfBaseName := strings.TrimSuffix(pdfIme, filepath.Ext(pdfIme))

	// prefix za PNG fajlove
	outputPrefix := filepath.Join(outputFolder, pdfBaseName+"_strana")

	// Konverzija PDF u slike
	cmd := exec.Command(
		// "D:\\poppler\\Library\\bin\\pdftoppm.exe",
		"E:\\aplikacije\\poppler\\Library\\bin\\pdftoppm.exe",
		"-png",
		"-r", "300",
		inputPath,
		outputPrefix,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Greška pri konverziji:", string(output))
		return
	}

	// Generiši listu fajlova u output folderu
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		fmt.Println("Greška pri čitanju foldera:", err)
		return
	}

	var slike []string
	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".png" {
			slike = append(slike, f.Name())
		}
	}

	// Napravi JSON fajl
	// ime JSON fajla po PDF-u
	jsonPath := filepath.Join(outputFolder, pdfBaseName+".json")
	fJSON, err := os.Create(jsonPath)
	if err != nil {
		fmt.Println("Greška pri kreiranju JSON-a:", err)
		return
	}
	defer fJSON.Close()

	enc := json.NewEncoder(fJSON)
	enc.SetIndent("", "  ")
	enc.Encode(slike)

	fmt.Println("Konverzija uspešna! JSON fajl kreiran:", jsonPath)
}
