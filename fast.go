package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const filePathF string = "./data/users.txt"

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	/*
		!!! !!! !!!
		обратите внимание - в задании обязательно нужен отчет
		делать его лучше в самом начале, когда вы видите уже узкие места, но еще не оптимизировалм их
		так же обратите внимание на команду в параметром -http
		перечитайте еще раз задание
		!!! !!! !!!
	*/

	file, err := os.Open(filePathF)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//scanner := bufio.NewScanner(file)
	fileContents, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileContents), "\n")
	users := make([]map[string]interface{}, 0) // , 1000)

	for _, line := range lines {
		//fmt.Println([]byte(line))
		user := make(map[string]interface{})
		// fmt.Printf("%v %v\n", err, line)
		err := json.Unmarshal([]byte(line), &user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	// var line []byte
	// for scanner.Scan() {
	// 	line = scanner.Bytes()
	// 	fmt.Println(line)
	// 	user := make(map[string]interface{})
	// 	// fmt.Printf("%v %v\n", err, line)
	// 	err := json.Unmarshal(line, &user)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	users = append(users, user)

	// }

	//fmt.Println(line)

	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers, ok := user["browsers"].([]interface{})
		if !ok {
			// log.Println("cant cast browsers")
			continue
		}

		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		// email := r.ReplaceAllString(user["email"].(string), " [at] ")
		email := strings.ReplaceAll(user["email"].(string), "@", " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user["name"], email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))

}
