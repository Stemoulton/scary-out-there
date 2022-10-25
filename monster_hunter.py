import os, re, json
import xml.etree.ElementTree as ET

monsterTotals = {'ghoul': 0, 'ghost': 0, 'vampire': 0,
            'zombie': 0, 'witch': 0, 'troll': 0}

monstersList = list(monsterTotals.keys())

# Example monster: 12 vampire
def add_monster(monster):
    data = monster.split()
    #Looks for match in monsters list - used to check for monster in xml and json files
    match = next((x for x in monstersList if x in data[1]), False)
    if match:
        #Changes plural to singular
        #Example in scary-book it's 'Zombie' so doesn't need to change to add to dict but in xml it's 'Zombies' so needs to be changed
        data[1] = match
        monsterTotals[data[1]] += int(data[0])

def search_txt(location):
    with open(f'./data/{location}.txt', 'r', encoding='utf-8') as file:
        #regex that looks for any number then a monster
        monsters = re.findall(r"\d+\s(?:" + r"|".join(monstersList) + r")",file.read().lower())        
    list(map(add_monster,monsters))

def search_xml(location):
    castle = ET.parse(f'./data/{location}.xml').getroot()
    for area in castle:
        try:
            add_monster(area.find('amount').text + " " + area.find('monster').text.lower())
        except ValueError:
            continue

def search_json(location):
    with open(f'./data/{location}.json', encoding='utf-8') as json_file:
        data = json.load(json_file)
    for key in data.keys():
        try:
            #count number of ids in keys
            add_monster(str(str(data[key]).count('id')) + " " + key)
        except ValueError:
            continue

def main(folder):
    locations = os.listdir(folder)
    files = list(map(lambda x: list(os.path.splitext(x)),locations))
    for file in files:
        if file[1] == '.txt':
            search_txt(file[0])
        elif file[1] == '.xml':
            search_xml(file[0])
        elif file[1] == '.json':
            search_json(file[0])
        else:
            continue
    print('\nFinal monster count: ', monsterTotals)

main('./data')
