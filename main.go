package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileContent := string(content)
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

type Castle struct {
	XMLName  xml.Name    `xml:"castle"`
	Hall     MonsterNode `xml:"hall"`
	Kitchen  MonsterNode `xml:"kitchen"`
	Basement MonsterNode `xml:"basement"`
	Attic    MonsterNode `xml:"attic"`
	Dungeon  MonsterNode `xml:"dungeon"`
}

type MonsterNode struct {
	Monster     string `xml:"monster"`
	Amount      int    `xml:"amount"`
	Description string `xml:"description"`
}

func MatchXMLFile(filePath string) map[string]int {
	monsters := map[string]int{
		"Ghouls":   0,
		"Ghoul":    0,
		"Ghosts":   0,
		"Ghost":    0,
		"Vampires": 0,
		"Vampire":  0,
		"Zombies":  0,
		"Zombie":   0,
		"Witches":  0,
		"Witch":    0,
		"Trolls":   0,
		"Troll":    0,
	}
	xmlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer xmlFile.Close()
	xmlContent, _ := ioutil.ReadAll(xmlFile)

	var castle Castle
	if err := xml.Unmarshal(xmlContent, &castle); err != nil {
		log.Fatal(err)
	}

	if _, ok := monsters[castle.Hall.Monster]; ok {
		monsters[castle.Hall.Monster] += castle.Hall.Amount
	}
	if _, ok := monsters[castle.Kitchen.Monster]; ok {
		monsters[castle.Kitchen.Monster] += castle.Kitchen.Amount
	}
	if _, ok := monsters[castle.Basement.Monster]; ok {
		monsters[castle.Basement.Monster] += castle.Basement.Amount
	}
	if _, ok := monsters[castle.Attic.Monster]; ok {
		monsters[castle.Attic.Monster] += castle.Attic.Amount
	}
	if _, ok := monsters[castle.Dungeon.Monster]; ok {
		monsters[castle.Dungeon.Monster] += castle.Dungeon.Amount
	}
	i := 0
	previousKey := ""
	for k := range monsters {
		if i%2 != 0 {
			monsters[previousKey] += monsters[k]
		} else {
			previousKey = k
		}
		i += 1
	}

	return monsters
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

func MergeResults(result1 map[string]int, result2 map[string]int, result3 map[string]int, result4 map[string]int) {
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

	for k, v := range result4 {
		if _, ok := result4[k]; ok {
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
	MergeResults(
		MatchTextFile("./data/bat-cave.txt"),
		MatchTextFile("./data/scary-book.txt"),
		MatchJsonFile("./data/scary-tomb.json"),
		MatchXMLFile("./data/scary-castle.xml"),
	)
}
