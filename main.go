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
	"time"
)

func main() {
	// Folder gde je exe, za log fajl
	exePath, _ := os.Executable()
	exeFolder := filepath.Dir(exePath)
	logPath := filepath.Join(exeFolder, "pdf_transformers.log")

	// Otvori log fajl (append mode)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Greška pri otvaranju log fajla:", err)
		return
	}
	defer logFile.Close()

	// Funkcija za upis u log sa timestampom
	log := func(msg string) {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logFile.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, msg))
		fmt.Println(msg) // opcionalno: ispiši i na ekran
	}

	// Separator za novi run
	log("===============================================")
	log("POČINJE NOVI RUN PDF TRANSFORMERS")
	log("===============================================")

	if len(os.Args) < 3 {
		// fmt.Println("Nedostaju argumenti: <folder_ime> <pdf_ime>")
		log("Nedostaju argumenti: <folder_ime> <pdf_ime>")
		return
	}

	folderIme := os.Args[1]
	pdfIme := os.Args[2]

	// basePath := "D:\\go_workspace\\pdf-transformers"
	basePath := "E:\\aplikacije\\tis\\ddn\\docs"

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

	log(fmt.Sprintf("Počinje konverzija PDF-a: %s", inputPath))

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
		// fmt.Println("Greška pri konverziji:", string(output))
		log(fmt.Sprintf("Greška pri konverziji PDF-a: %v", err))
		log(fmt.Sprintf("Detalji izlaza: %s", string(output)))
		return
	}

	log("Konverzija PDF-a završena uspešno.")

	// Generiši listu fajlova u output folderu
	files, err := os.ReadDir(outputFolder)
	if err != nil {
		// fmt.Println("Greška pri čitanju foldera:", err)
		log(fmt.Sprintf("Greška pri čitanju foldera %s: %v", outputFolder, err))
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
		// fmt.Println("Greška pri kreiranju JSON-a:", err)
		log(fmt.Sprintf("Greška pri kreiranju JSON-a: %v", err))
		return
	}
	defer fJSON.Close()

	enc := json.NewEncoder(fJSON)
	enc.SetIndent("", "  ")
	enc.Encode(slike)

	log(fmt.Sprintf("JSON fajl kreiran: %s", jsonPath))
	log(fmt.Sprintf("Broj generisanih PNG fajlova: %d", len(slike)))
	log("=== Kraj procesa ===")
}
