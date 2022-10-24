import pandas as pd
from pathlib import Path
import re
import json

class TextFile():
    """Class for text files
    """
    def __init__(self, path :Path) -> None:
        with open(path, 'r') as f:
            self.data = f.read()

    def count_monsters(self) -> int:
        """Count monsters"""
        count_str = re.findall(rf'(\d+)\s+[{monster_str}]', self.data, re.IGNORECASE)
        return sum(int(x) for x in count_str)
 
class XMLFile():
    """Class for XML files
    """
    def __init__(self, path: Path) -> None:
        self.data = pd.read_xml(path)

    def count_monsters(self) -> int:
        """"""
        return self.data[self.data.monster.str.contains(monster_str, case=False)].amount.sum()
 

class JSONFile():
    """Class for json files
    """
    def __init__(self, path :Path) -> None:
        with open(path, 'r') as f:
            self.data = json.load(f)

    def count_monsters(self) -> int:
        return sum(1 if isinstance(self.data[x], dict) else len(self.data[x]) for x in self.data.keys() if x.lower() in monster_str.lower())
 

class File():
    """Class for abstract file"""
    def __init__(self, path :Path) -> None:
        file = Path(path)
        ext = file.suffix
        if ext == '.txt':
            self.file_class = TextFile(path)
        elif ext == '.xml':
            self.file_class = XMLFile(path)
        elif ext == '.json':
            self.file_class = JSONFile(path)
        else:
            raise("Invalid file type")  # type: ignore

    def count_monsters(self) -> int:
        """Count monsters"""
        return self.file_class.count_monsters()


        
MONSTERS = ['Ghouls', 'Ghosts', 'Vampires', 'Zombies', 'Witches', 'Trolls', 'Ghoul', 'Ghost', 'Vampire', 'Zombie', 'Witch', 'Troll']

monster_str = '|'.join(MONSTERS)

data_dir = Path('data/')

files= data_dir.glob('*.*')

total_monsters = 0
for f in files:
    file = File(f)
    total_monsters += file.count_monsters()

print(f"Found {total_monsters} monsters in total")
