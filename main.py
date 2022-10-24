
import xml.etree.ElementTree as ET
import json

monster_names = ["Ghouls", "Ghosts",
                 "Vampires", "Zombies", "Witches", "Trolls"]
total_of_monsters = {"Ghouls": 0, "Ghosts": 0, "Vampires": 0,
                     "Zombies": 0,  "Witches": 0, "Trolls": 0}
monster_name_lookup = {"Ghoul": "Ghouls",
                       "Zombies": "Zombies",  "Witch": "Witches", "Troll":  "Trolls"}
text_files = ["data/bat-cave.txt", "data/scary-book.txt"]


def main():
    print("-----  Stephen's Monster finder! -----")


def search_str(file_path, word_to_find):
    count_words = 0
    with open(file_path, 'r') as fp:
        for l_no, line in enumerate(fp):
            words = line.split(' ')
            for idx, current_word in enumerate(words):
                if current_word.lower() == word_to_find.lower():
                    count_words += int(words[idx-1])

    return count_words


if __name__ == "__main__":
    main()

    # Get data from text files
    for file_idx, file_name in enumerate(text_files):
        for idx, monster in enumerate(monster_names):
            num_monsters = search_str(file_name, monster)
            total_of_monsters[monster] += num_monsters

    print("After Text Files:", total_of_monsters)

    # Get data from XML
    tree = ET.parse('data/scary-castle.xml')
    root = tree.getroot()

    for room in list(root):
        monster = room[0].text
        num_monsters = int(room[1].text)

        if total_of_monsters.get(monster_name_lookup.get(monster)) is not None:
            total_of_monsters[monster_name_lookup.get(monster)] += num_monsters

    print("After XML Files", total_of_monsters)

    # Get data  from JSON
    with open('data/scary-tomb.json', 'r') as f:
        data = json.load(f)

    for (k, v) in data.items():
        if total_of_monsters.get(monster_name_lookup.get(k)) is not None:
            if type(v) is list:
                total_of_monsters[monster_name_lookup.get(monster)] += len(v)
            else:
                # We just have an object, so count increases by 1
                total_of_monsters[monster_name_lookup.get(monster)] += 1

    print("After JSON Files", total_of_monsters)
