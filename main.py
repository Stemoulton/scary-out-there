
# Create an array with 4 numbers
monsters = ["Ghouls", "Ghosts", "Vampires", "Zombies", "Witches", "Trolls"]

def main():
    print("Stephen's Monster finder!")

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

    for idx, monster in enumerate(monsters):
        num_witches =  search_str(r'data/bat-cave.txt', monster)
        print("Total number of", monster, ":", num_witches)
    