## Rick and Morty REPL

This program will allow users to interact with the Rick and Morty API. You can find the API documentation in the following link: https://rickandmortyapi.com/documentation

The program allows the user to retrieve a list of characters, specific character information, saving and deleting favourite characters, and exporting as a JSON file the saved characters.

### Start the REPL
1. Type go build
2. Execute the .exe file

### Main menu commands
help : Displays a help message. \n
exit : Exit the program. \n
character : Access all characters information.

### Character menu commands
del : Type this command followed by the id of the character you want to delete 
export : Export your saved character 
import : Type this command followed by the path of the file that contains your saved characters
map : Get the a list with the next 20 characters
view : Type this command followed by the id of the character you want to see.
save : Type this command followed by the id of the character you want to save.
list : See your saved characters
help : See the list of available commands for characters
back : Go back to the main menu.
mapb : Get the a list with the previous 20 characters
exit : Exit the program.


### Examples of Character menu commands
1.
    Please enter a command to explore charaters > map
    Characters: 
    ID: 1, Name: Rick Sanchez, Status: Alive
    ID: 2, Name: Morty Smith, Status: Alive
    ID: 3, Name: Summer Smith, Status: Alive
    ID: 4, Name: Beth Smith, Status: Alive
    ID: 5, Name: Jerry Smith, Status: Alive
    ID: 6, Name: Abadango Cluster Princess, Status: Alive
    ID: 7, Name: Abradolf Lincler, Status: unknown
    ID: 8, Name: Adjudicator Rick, Status: Dead
    ID: 9, Name: Agency Director, Status: Dead
    ID: 10, Name: Alan Rails, Status: Dead
    ID: 11, Name: Albert Einstein, Status: Dead
    ID: 12, Name: Alexander, Status: Dead
    ID: 13, Name: Alien Googah, Status: unknown
    ID: 14, Name: Alien Morty, Status: unknown
    ID: 15, Name: Alien Rick, Status: unknown
    ID: 16, Name: Amish Cyborg, Status: Dead
    ID: 17, Name: Annie, Status: Alive
    ID: 18, Name: Antenna Morty, Status: Alive
    ID: 19, Name: Antenna Rick, Status: unknown
    ID: 20, Name: Ants in my Eyes Johnson, Status: unknown


2. 
    Please enter a command to explore charaters > view 1
    ==========================================
    Printing information of: Rick Sanchez
    ==========================================
    - ID: 1
    - Status: Alive
    - Species: Human
    - Gender: Male
    - Origin: Earth (C-137)
    - Location: Citadel of Ricks

3.
    Please enter a command to explore charaters > save 2
    Saving character...
    ID: 2, Name: Morty Smith saved!

4.
    Please enter a command to explore charaters > list
    Printing your characters...
    -------------------------------
    - ID: 2
    - Name: Morty Smith
    - Status: Alive
    - Species: Human
    - Gender: Male
    - Origin: unknown
    - Location: Citadel of Ricks