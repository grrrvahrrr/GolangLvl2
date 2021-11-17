//Lesson 6 HW
package main

import (
	"GolangLvl2/lesson6/lesson6/secrets"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
	"sync"
)

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	if !strings.Contains(strings.ToLower(body.Message.Text), "marco") {
		return
	}

	var totalPolos int

	if _, err := os.Stat("./poloStats"); err == nil {
		file, err := os.Open("./poloStats")
		if err != nil {
			log.Println("Couldn't open file")
		}
		defer func() {
			err := file.Close()
			if err != nil {
				log.Println("Couldn't close file")
			}
		}()

		fileinfo, err := file.Stat()
		if err != nil {
			log.Println(err)
			return
		}
		filesize := fileinfo.Size()
		buffer := make([]byte, filesize)

		_, err = file.Read(buffer)
		if err != nil {
			log.Println(err)
			return
		}
		totalPolos, err = strconv.Atoi(string(buffer))
		if err != nil {
			log.Println("Couldn't read from file")
		}
	}

	// If the text contains marco, call the `sayPolo` function
	if err := sayPolo(body.Message.Chat.ID, totalPolos); err != nil {
		log.Println("error in sending reply:", err)
		return
	}
	log.Println("reply sent")
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func writeToPoloStats(s string) error {
	file, err := os.Create("./poloStats")
	if err != nil {
		return err
	}
	defer func() error {
		err := file.Close()
		if err != nil {
			return err
		}
		return nil
	}()
	stats := []byte(s)
	datawriter := bufio.NewWriter(file)
	for _, data := range stats {
		_, _ = datawriter.WriteString(string(data))
	}
	datawriter.Flush()

	return nil
}

// sayPolo takes a chatID and sends "polo" to them
func sayPolo(chatID int64, totalPolos int) error {
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		mu.Lock()
		totalPolos++
		mu.Unlock()

		wg.Done()
	}()
	wg.Wait()

	totalPolosString := strconv.Itoa(totalPolos)
	if _, err := os.Stat("./poloStats"); errors.Is(err, os.ErrNotExist) {
		err = writeToPoloStats(totalPolosString)
		if err != nil {
			log.Println("Couldn't write to file", err)
		}
	} else {
		err := os.Remove("./poloStats")
		if err != nil {
			log.Println("Couldn't remove file")
		}
		err = writeToPoloStats(totalPolosString)
		if err != nil {
			log.Println("Couldn't write to file", err)
		}
	}

	var response string
	if totalPolos == 1 {
		response = "Polo!!\nIn my lifetime I found Polo " + totalPolosString + " time"
	} else {
		response = "Polo!!\nIn my lifetime I found Polo " + totalPolosString + " times"
	}

	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   response,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	address := "https://api.telegram.org/bot" + secrets.BotID + "/sendMessage"
	res, err := http.Post(address, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//Task 1
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))

	//Task 2
	//task2.WorkerPull()

	//Task 3
	//task3.NumCount()
}
