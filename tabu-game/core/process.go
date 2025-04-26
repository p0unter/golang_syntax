package core

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Ask          string   `json:"ask"`
	Answer       string   `json:"ans"`
	TabuKeywords []string `json:"tabkeys"`
}

func ReadData() []Question {
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatal(err)
	}
	var questions []Question
	err = json.Unmarshal(data, &questions)
	if err != nil {
		log.Fatal(err)
	}
	return questions
}

func PersonersVS() {
	questions := ReadData()
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	for {
		randq := questions[rand.Intn(len(questions))]

		fmt.Printf("\nKelime: %v\n", randq.Answer)
		fmt.Println("Yasaklı Kelimeler:")

		for _, keyword := range randq.TabuKeywords {
			fmt.Println("-", keyword)
		}

		fmt.Println("\nDevam etmek için Enter tuşuna basın, çıkmak için 'q' yazıp Enter'a basın.")
		scanner.Scan()
		input := scanner.Text()
		if input == "q" || input == "Q" {
			fmt.Println("Oyun sonlandırıldı.")
			os.Exit(0)
		}
	}
}

func RobotVS() {
	questions := ReadData()
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

	for {
		q := questions[rand.Intn(len(questions))]

		fmt.Println("\nRobot VS - İpucu:")
		fmt.Println(q.Ask)

		fmt.Println("Tahmininizi yazınız (Çıkmak için 'q' yazıp Enter'a basın):")

		scanner.Scan()
		answer := scanner.Text()

		if answer == "q" || answer == "Q" {
			fmt.Println("Uygulama kapatılıyor...")
			os.Exit(0)
		}

		if strings.EqualFold(answer, q.Answer) {
			fmt.Println("Tebrikler! Doğru cevap.")
		} else {
			fmt.Printf("Yanlış cevap. Doğru cevap: %s\n", q.Answer)
		}

		fmt.Println("Devam etmek için Enter'a basın, çıkmak için 'q' yazıp Enter'a basın.")
		scanner.Scan()
		cont := scanner.Text()
		if cont == "q" || cont == "Q" {
			fmt.Println("Uygulama kapatılıyor...")
			os.Exit(0)
		}
	}
}
