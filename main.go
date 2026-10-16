package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func obrazNaBinarny(sciezkaObrazu, sciezkaWyjsciowaBin string) error {
	fmt.Printf("Odczytywanie pliku: %s...\n", sciezkaObrazu)
	start := time.Now()

	plikObrazu, err := os.Open(sciezkaObrazu)
	if err != nil {
		return err
	}
	defer plikObrazu.Close()

	plikBinarny, err := os.Create(sciezkaWyjsciowaBin)
	if err != nil {
		return err
	}
	defer plikBinarny.Close()

	reader := bufio.NewReader(plikObrazu)
	writer := bufio.NewWriter(plikBinarny)

	buffer := make([]byte, 1024*1024) // paczka 1 MB
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			var sb strings.Builder
			sb.Grow(n * 8)
			for i := 0; i < n; i++ {
				sb.WriteString(fmt.Sprintf("%08b", buffer[i]))
			}
			_, writeErr := writer.WriteString(sb.String())
			if writeErr != nil {
				return writeErr
			}
		}
		if err != nil {
			break
		}
	}

	writer.Flush()
	fmt.Printf("Obraz zapisany jako kod binarny (czas: %.2f s)\n", time.Since(start).Seconds())
	return nil
}

func binarnyNaObraz(sciezkaBinarna, sciezkaWyjsciowaObrazu string) error {
	fmt.Printf("Odczytywanie pliku tekstowego i składanie obrazu...\n")
	start := time.Now()

	plikBinarny, err := os.Open(sciezkaBinarna)
	if err != nil {
		return err
	}
	defer plikBinarny.Close()

	plikObrazu, err := os.Create(sciezkaWyjsciowaObrazu)
	if err != nil {
		return err
	}
	defer plikObrazu.Close()

	reader := bufio.NewReader(plikBinarny)
	writer := bufio.NewWriter(plikObrazu)

	textBuffer := make([]byte, 1024*1024*8)
	var buforBitow strings.Builder

	for {
		n, err := reader.Read(textBuffer)
		if n > 0 {
			buforBitow.WriteString(string(textBuffer[:n]))
			str := buforBitow.String()

			ilePelnych := (len(str) / 8) * 8
			doPrzetworzenia := str[:ilePelnych]
			buforBitow.Reset()
			buforBitow.WriteString(str[ilePelnych:])

			bytesToWrite := make([]byte, 0, len(doPrzetworzenia)/8)
			for i := 0; i < len(doPrzetworzenia); i += 8 {
				bajtStr := doPrzetworzenia[i : i+8]
				var b byte
				_, _ = fmt.Sscanf(bajtStr, "%b", &b)
				bytesToWrite = append(bytesToWrite, b)
			}
			writer.Write(bytesToWrite)
		}
		if err != nil {
			break
		}
	}

	writer.Flush()
	fmt.Printf("Obraz zwrocony lmao: %s (czas: %.2f s)\n", sciezkaWyjsciowaObrazu, time.Since(start).Seconds())
	return nil
}

func main() {
	fmt.Println("=== GOLANG BEZPIECZNY KONWERTER OBRAZÓW ===")
	fmt.Print("Podaj nazwę pliku (np. zdjecie.ARW lub obraz.jpg): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	nazwaOryginalu := strings.TrimSpace(input)

	if _, err := os.Stat(nazwaOryginalu); os.IsNotExist(err) {
		fmt.Printf("Błąd: Plik '%s' nie istnieje w tym folderze!\n", nazwaOryginalu)
		return
	}

	rozszerzenie := filepath.Ext(nazwaOryginalu)
	plikTymczasowyBin := "obraz_binarny.txt"
	nazwaOdtworzonego := "odtworzone_zdjecie" + rozszerzenie

	err := obrazNaBinarny(nazwaOryginalu, plikTymczasowyBin)
	if err != nil {
		fmt.Println("Błąd konwersji na binarny:", err)
		return
	}

	err = binarnyNaObraz(plikTymczasowyBin, nazwaOdtworzonego)
	if err != nil {
		fmt.Println("Błąd konwersji na obraz:", err)
		return
	}
}