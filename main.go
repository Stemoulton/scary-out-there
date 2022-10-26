package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MonsterNode struct {
	Monster     string `xml:"monster"`
	Amount      string `xml:"amount"`
	Description string `xml:"description"`
}

func MatchTextFile(filePath string) map[string]int {
	monsters := map[string]int{
		"Ghouls":   0,
		"Ghosts":   0,
		"Vampires": 0,
		"Zombies":  0,
		"Witches":  0,
		"Trolls":   0,
	}
	fileMonsters := monsters
	fileContent := ReadTextFile(filePath)
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*\s((?:G|g)houls|(?:G|g)hosts|(?:V|v)ampires|(?:Z|z)ombi(?:e|es)|(?:W|w)itches|(?:T|t)roll(?: |s))`)
	submatchall := re.FindAllString(fileContent, -1)
	for _, element := range submatchall {
		monster := strings.Fields(element)
		switch monster[1] {
		case "witches":
			monster[1] = "Witches"
		case "Troll":
			monster[1] = "Trolls"
		case "Zombie":
			monster[1] = "Zombies"
		}

		value, exists := fileMonsters[monster[1]]
		if exists {
			amount, _ := strconv.Atoi(monster[0])
			fileMonsters[monster[1]] = value + amount
		}

	}
	return fileMonsters
}

func MatchJsonFile(filePath string) map[string]int {
	monsters := map[string]int{
		"Ghouls":   0,
		"Ghosts":   0,
		"Vampires": 0,
		"Zombies":  0,
		"Witches":  0,
		"Trolls":   0,
	}
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	result := map[string]interface{}{}
	json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		switch key {
		case "Troll":
			monsters["Trolls"] = NestedCheck(value)
			break
		case "Witch":
			monsters["Witches"] = NestedCheck(value)
			break
		case "Vampire":
			monsters["Vampires"] = NestedCheck(value)
			break
		case "Ghoul":
			monsters["Ghouls"] = NestedCheck(value)
			break
		case "Ghost":
			monsters["Ghosts"] = NestedCheck(value)
			break
		case "Zombie":
			monsters["Zombies"] = NestedCheck(value)
			break

		}
	}
	return monsters

}

func ReadTextFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}

func NestedCheck(value interface{}) int {
	_, ok := value.(map[string]interface{})
	if ok {
		return 1
	}
	firstLastNames := map[string]string{}
	for _, v := range value.([]interface{}) {
		monsterInterface, exists := v.(map[string]interface{})
		if exists {
			firstname, ok := monsterInterface["name"].(string)
			if ok {
				lastname, ok := monsterInterface["lastname"].(string)
				if ok {
					firstLastNames[firstname] = lastname
				}
			}
		}
	}

	return len(firstLastNames)
}

func MergeResults(result1 map[string]int, result2 map[string]int, result3 map[string]int) {
	final := make(map[string]int)
	for k, v := range result1 {
		if _, ok := result1[k]; ok {
			final[k] += v
		}
	}

	for k, v := range result2 {
		if _, ok := result2[k]; ok {
			final[k] += v
		}
	}

	for k, v := range result3 {
		if _, ok := result3[k]; ok {
			final[k] += v
		}
	}

	fmt.Println("Number of Ghouls:", final["Ghouls"])
	fmt.Println("Number of Ghosts:", final["Ghosts"])
	fmt.Println("Number of Vampires:", final["Vampires"])
	fmt.Println("Number of Trolls:", final["Trolls"])
	fmt.Println("Number of Zombies:", final["Zombies"])
	fmt.Println("Number of Witches:", final["Witches"])
}

func main() {
	result1 := MatchTextFile("./data/bat-cave.txt")
	result2 := MatchTextFile("./data/scary-book.txt")
	result3 := MatchJsonFile("./data/scary-tomb.json")
	MergeResults(result1, result2, result3)
}
