import pandas as pd
from pathlib import Path
import re
import json
import numpy as np
from collections import Counter


class TextFile:
    """Class for text files
    """

    def __init__(self, path: Path) -> None:
        with open(path, "r") as f:
            self.data = f.read()

    def count_monsters(self) -> Counter:
        """Count monsters"""
        count_str = re.findall(
            rf"(\d+)\s+({monster_str})", self.data, flags=re.IGNORECASE
        )
        return sum((Counter({k.capitalize(): int(v)}) for v, k in count_str), Counter())


class XMLFile:
    """Class for XML files
    """

    def __init__(self, path: Path) -> None:
        self.data = pd.read_xml(path)

    def count_monsters(self) -> Counter:
        """Count monsters"""
        df = self.data[self.data.monster.str.contains(monster_str, case=False)]
        df_sum = df.groupby("monster").agg({"amount": np.sum})
        return Counter(df_sum.T.to_dict("records")[0])


class JSONFile:
    """Class for json files
    """

    def __init__(self, path: Path) -> None:
        with open(path, "r", encoding="utf-8") as f:
            self.data = json.load(f)

    def count_monsters(self) -> Counter:
        return Counter(
            {
                x: 1 if isinstance(self.data[x], dict) else len(self.data[x])
                for x in self.data.keys()
                if x.lower() in monster_str.lower()
            }
        )


class File:
    """Class for abstract file"""

    def __init__(self, path: Path) -> None:
        file = Path(path)
        ext = file.suffix
        if ext == ".txt":
            self.file_class = TextFile(path)
        elif ext == ".xml":
            self.file_class = XMLFile(path)
        elif ext == ".json":
            self.file_class = JSONFile(path)
        else:
            raise ("Invalid file type")  # type: ignore

    def count_monsters(self) -> Counter:
        """Count monsters"""
        return self.file_class.count_monsters()


def main(data_folder) -> None:
    data_dir = Path(data_folder)
    files = data_dir.glob("*.*")

    counts = Counter({})
    for f in files:
        file = File(f)
        counts += {
            monster_plural.get(k, k): v for k, v in file.count_monsters().items()
        }

    for k, v in counts.items():
        print(f"* {k} = {v}")


if __name__ == "__main__":

    MONSTERS = [
        "Ghouls",
        "Ghosts",
        "Vampires",
        "Zombies",
        "Witches",
        "Trolls",
        "Ghoul",
        "Ghost",
        "Vampire",
        "Zombie",
        "Witch",
        "Troll",
    ]
    monster_plural = {
        "Ghoul": "Ghouls",
        "Ghost": "Ghosts",
        "Vampire": "Vampires",
        "Zombie": "Zombies",
        "Witch": "Witches",
        "Troll": "Trolls",
    }

    monster_str = "|".join(MONSTERS)

    main("data/")
