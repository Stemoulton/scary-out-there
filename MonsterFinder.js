const fs = require("fs");

const scaryTomb = require("./data/scary-tomb.json");
const batCaveMonsters = fs
  .readFileSync("./data/bat-cave.txt", "utf8")
  .toLowerCase()
  .split(/[\s,]+/);
const scaryBookMonsters = fs
  .readFileSync("./data/scary-book.txt", "utf8")
  .toLowerCase()
  .split(/[\s,]+/);
const scaryCastle = fs
  .readFileSync("./data/scary-castle.xml", "utf8")
  .toLowerCase()
  .split(/<[^>]*>/g);

const monsters = new Map([
  ["ghoul", 0],
  ["ghouls", 0],
  ["ghost", 0],
  ["ghosts", 0],
  ["vampire", 0],
  ["vampires", 0],
  ["zombie", 0],
  ["zombies", 0],
  ["witch", 0],
  ["witches", 0],
  ["troll", 0],
  ["trolls", 0],
]);

const addMonstersToMap = (monsterMap, monster, numberOfMonsters) => {
  if (numberOfMonsters > 0) {
    monsterMap.set(
      monster,
      parseFloat(monsterMap.get(monster)) + parseFloat(numberOfMonsters)
    );
  }
};

batCaveMonsters.forEach((value, index) => {
  if (monsters.has(value)) {
    addMonstersToMap(monsters, value, batCaveMonsters[index - 1]);
  }
});

scaryBookMonsters.forEach((value, index) => {
  if (monsters.has(value)) {
    addMonstersToMap(monsters, value, scaryBookMonsters[index - 1]);
  }
});

Object.keys(scaryTomb).forEach((key) => {
  let currentMonster = key.toLowerCase();

  if (monsters.has(currentMonster)) {
    let numberOfMonsters;

    if (scaryTomb[key].length === undefined) {
      numberOfMonsters = 1;
    } else {
      numberOfMonsters = scaryTomb[key].length;
    }

    addMonstersToMap(monsters, currentMonster, numberOfMonsters);
  }
});

scaryCastle.forEach((value, index) => {
  if (monsters.has(value)) {
    //this is robust and will never fail
    addMonstersToMap(monsters, value, scaryCastle[index + 2]);
  }
});

const totalMonsters = new Map([
  ["Ghouls", (monsters.get("ghoul") + monsters.get("ghouls"))],
  ["Ghosts", monsters.get("ghost") + monsters.get("ghosts")],
  ["Vampires", monsters.get("vampire") + monsters.get("vampires")],
  ["Zombies", monsters.get("zombie") + monsters.get("zombies")],
  ["Witch", monsters.get("witch") + monsters.get("witches")],
  ["Troll", monsters.get("troll") + monsters.get("trolls")],
]);

console.log(totalMonsters);
