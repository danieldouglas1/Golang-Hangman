package main
import (
    "os"
    "fmt"
    "time"
    "math/rand"
    "log"
    "os/exec"
    "bufio"
    "strings"
  )

//Clears terminal screen
func clearscreen() {
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}

// Reads in words from dictionary text file and stores in map
func readDictionary() map[int]string{
  i := 0
  m := make(map[int]string)

  //error checking
  file, err := os.Open("Dictionary.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  //reads words in file to map
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    m[i]= scanner.Text()
    i = i+1
  }

  //error checking
  if err := scanner.Err(); err !=nil {
    log.Fatal(err)
  }
  //returns map
  return m
}

//Allows user to add word to the dictionary
func addToDictionary(word string){

  //  file, err:= os.OpenFile("Dictionary.txt",os.O_WRONLY|os.O_APPEND, 0644)
  f, err := os.OpenFile("Dictionary.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil {
      panic(err)
  }

  defer f.Close()

  if _, err = f.WriteString("\n" + word); err != nil {
      panic(err)
  }

}

// Gets random word from the file
func get_random_word (m map[int]string) string  {
  var randWord string
  rand.Seed(time.Now().UnixNano())
  randWord = m[rand.Intn(len(m))]

  return randWord
}

//Outputs the different stages of Hangman
func draw_hangman(stage_of_death int) {
	switch stage_of_death {
	case 0:
		fmt.Printf("  +---+\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 1:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 2:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")


	case 3:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|   |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")


	case 4:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 5:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" /    |\n")
		fmt.Printf("      |\n")
		fmt.Printf("      |\n")
		fmt.Printf("========\n")

	case 6:
		fmt.Printf("  +---+\n")
		fmt.Printf("  |   |\n")
		fmt.Printf("  O   |\n")
		fmt.Printf("_/|\\_ |\n")
		fmt.Printf("  |   |\n")
		fmt.Printf(" / \\  |\n")
		fmt.Printf("      |\n")
		fmt.Printf("R.I.P |\n")
		fmt.Printf("========\n")
	}
}

//Hides the letters with underscores
func hideword(wordlen int) []string{

  dashes := make([]string, wordlen, 2*wordlen)
	for i := 0; i < wordlen; i++ {
		dashes[i]= "_"
	}
	return dashes
}

//reveals letters if they are in the word, also returns true if a letter
// is revealed
func reveal(word string, guess string, dash []string) ([]string, bool){
  revealed:= false
  for i := 0; i < len(word); i++ {
  	if strings.EqualFold(guess, string(word[i])){
      dash[i]= guess
      revealed=true
  }
}
  return dash,revealed
}

//Converts array of dashes and letters to string
func arrayToString(dashes []string) string{
  return strings.Join(dashes, "")
}

// Checks if letter was already used
func contains( slice map[int]string, val string) bool {
    for _, item := range slice {
        if strings.EqualFold(item, val){
            return true
        }
    }
    return false
}

//Checks if user won the game
func checkWin(guessedWord string, word string) bool{
  if(guessedWord== word){
      return true
  }
  return false

}

//Prints a message if user has won
func userWin(){

  fmt.Println("\n")
  fmt.Println("CONGRATS, YOU WON")

  fmt.Println("YYYYYYY       YYYYYYY                                     WWWWWWWW                           WWWWWWWW iiii")
  fmt.Println("Y:::::Y       Y:::::Y                                     W::::::W                           W::::::W iiii ")
  fmt.Println("YYY:::::Y   Y:::::YYYooooooooooo   uuuuuu    uuuuuu        W:::::W           WWWWW           W:::::Wiiiiiiinnnn  nnnnnnnn")
  fmt.Println("   Y:::::Y Y:::::Y oo:::::::::::oo u::::u    u::::u         W:::::W         W:::::W         W:::::W i:::::in:::nn::::::::nn")
  fmt.Println("    Y:::::Y:::::Y o:::::::::::::::ou::::u    u::::u          W:::::W       W:::::::W       W:::::W   i::::in::::::::::::::nn ")
  fmt.Println("     Y:::::::::Y  o:::::ooooo:::::ou::::u    u::::u           W:::::W     W:::::::::W     W:::::W    i::::inn:::::::::::::::n")
  fmt.Println("      Y:::::::Y   o::::o     o::::ou::::u    u::::u            W:::::W   W:::::W:::::W   W:::::W     i::::i  n:::::nnnn:::::n")
  fmt.Println("       Y:::::Y    o::::o     o::::ou::::u    u::::u             W:::::W W:::::W W:::::W W:::::W      i::::i  n::::n    n::::n")
  fmt.Println("       Y:::::Y    o::::o     o::::ou::::u    u::::u              W:::::W:::::W   W:::::W:::::W       i::::i  n::::n    n::::n")
  fmt.Println("       Y:::::Y    o::::o     o::::ou:::::uuuu:::::u               W:::::::::W     W:::::::::W        i::::i  n::::n    n::::n")
  fmt.Println("       Y:::::Y    o:::::ooooo:::::ou:::::::::::::::uu              W:::::::W       W:::::::W        i::::::i n::::n    n::::n")
  fmt.Println("    YYYY:::::YYYY o:::::::::::::::o u:::::::::::::::u               W:::::W         W:::::W         i::::::i n::::n    n::::n")
  fmt.Println("    Y:::::::::::Y  oo:::::::::::oo   uu::::::::uu:::u                W:::W           W:::W          i::::::i n::::n    n::::n")
  fmt.Println("    YYYYYYYYYYYYY    ooooooooooo       uuuuuuuu  uuuu                 WWW             WWW           iiiiiiii nnnnnn    nnnnnn")


}


func userLose(word string) {

  fmt.Println("The word was: " + word)

  fmt.Println(" ___________.._______")
  fmt.Println("| .__________))______|")
  fmt.Println("| | / /      ||")
  fmt.Println("| |/ /       ||")
  fmt.Println("| / /        ||.-''.")
  fmt.Println("| |/         |/  _  \\ ")
  fmt.Println("| |          ||  `/,|")
  fmt.Println("| |          (\\`_.' ")
  fmt.Println("| |         .-`--'.")
  fmt.Println("| |        /Y . . Y\\ ")
  fmt.Println("| |       // |   | \\ ")
  fmt.Println("| |      //  | . |  \\")
  fmt.Println("| |     ')   |   |   (`")
  fmt.Println("| |          ||'|| ")
  fmt.Println("| |          || || ")
  fmt.Println("| |          || ||")
  fmt.Println("| |          || ||")
  fmt.Println("| |         / | | \\ ")
  fmt.Println(" '''''''''|_`-' `-' |'''| ")
  fmt.Println("|'|'''''''\\ \\      '''| ")
  fmt.Println("| |        \\ \\        | | ")




  fmt.Println("▄         ▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄         ▄       ▄            ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄")
  fmt.Println("▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░▌       ▐░▌     ▐░▌          ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌")
  fmt.Println("▐░▌       ▐░▌▐░█▀▀▀▀▀▀▀█░▌▐░▌       ▐░▌     ▐░▌          ▐░█▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀ ▐░█▀▀▀▀▀▀▀▀▀")
  fmt.Println("▐░▌       ▐░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌▐░▌          ▐░▌")
  fmt.Println("▐░█▄▄▄▄▄▄▄█░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄▄▄")
  fmt.Println("▐░░░░░░░░░░░▌▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌")
  fmt.Println(" ▀▀▀▀█░█▀▀▀▀ ▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌ ▀▀▀▀▀▀▀▀▀█░▌▐░█▀▀▀▀▀▀▀▀▀")
  fmt.Println("     ▐░▌     ▐░▌       ▐░▌▐░▌       ▐░▌     ▐░▌          ▐░▌       ▐░▌          ▐░▌▐░▌")
  fmt.Println("     ▐░▌     ▐░█▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄█░▌     ▐░█▄▄▄▄▄▄▄▄▄ ▐░█▄▄▄▄▄▄▄█░▌ ▄▄▄▄▄▄▄▄▄█░▌▐░█▄▄▄▄▄▄▄▄▄")
  fmt.Println("     ▐░▌     ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌     ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌ ")
  fmt.Println("      ▀       ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀       ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀▀▀▀▀ ")

}

//Hangman gameplay
func gameplay(){

  //store letters the user guesses incorrectly
  letters_guessed := make(map[int]string)
  var count int         //number of times user guesses incorrectly
  count = 0
  var word string       //variable to hold word from dictionary
  var userguess string  //holds letter guessed by user
  var numwrong int      //number of times user guesses incorrectly


  //generates a random word
  word = get_random_word(readDictionary())
  gamewordlen := len(word)

  //prints dashes
  dashes := make([]string, gamewordlen, 2*gamewordlen)
  dashes = hideword(gamewordlen)

  //play as long as number of guesses < 6
   for numwrong < 6{
     draw_hangman(numwrong)
     fmt.Println()
     fmt.Println(strings.Trim(fmt.Sprint(dashes), "[]"))
     fmt.Println()

     fmt.Println("Your Letters Guessed")
     for _,element := range letters_guessed{
        fmt.Print(element)
        fmt.Print(" ")
    }

     fmt.Println()
     fmt.Print("Enter a guess: ")
     fmt.Scanln(&userguess)

    //if user used letter, tell them
     if contains(letters_guessed, userguess)==true {
        clearscreen()
        fmt.Println("You have already guessed this letter")
        continue
     }

     // add letters used to bank
     letters_guessed[count] = strings.ToUpper(userguess);
     count = count+1
     dashes, revealed := reveal(word, strings.ToUpper(userguess), dashes)
     fmt.Println()
     fmt.Println(arrayToString(dashes))

     fmt.Println()
     fmt.Println()

     // if a letter was not reveled, then
     if revealed == false{
        numwrong = numwrong +1
        if numwrong == 6 {
          userLose(word)
          break
        }
     }
     // Checks if User has Won
     if checkWin(arrayToString(dashes), strings.ToUpper(word))== true{
       userWin()
       break
     }
     clearscreen()
  }

}


func main() {
  var option int
  fmt.Println("██╗  ██╗ █████╗ ███╗   ██╗ ██████╗ ███╗   ███╗ █████╗ ███╗   ██╗")
  fmt.Println("██║  ██║██╔══██╗████╗  ██║██╔════╝ ████╗ ████║██╔══██╗████╗  ██║")
  fmt.Println("███████║███████║██╔██╗ ██║██║  ███╗██╔████╔██║███████║██╔██╗ ██║")
  fmt.Println("██╔══██║██╔══██║██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══██║██║╚██╗██║")
  fmt.Println("██║  ██║██║  ██║██║ ╚████║╚██████╔╝██║ ╚═╝ ██║██║  ██║██║ ╚████║")
  fmt.Println("╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝")
  fmt.Println("Welcome to Hangman")
  fmt.Println("Please Select an Option ")
  fmt.Println("1: Play Game ")
  fmt.Println("2: Add to Dictionary ")
  fmt.Println("3: Quit ")
  fmt.Println("")
  fmt.Print("Your Choice: ")
  fmt.Scanln(&option)

  switch option{
  case 1:
    gameplay()
  case 2:
    var newWord string
    fmt.Println("What word do you want to add?")
    fmt.Scanln(&newWord)
    addToDictionary(newWord)
    main()
  case 3:
    fmt.Println("Goodbye!")
    os.Exit(0)
  default:
    fmt.Println("Not Valid Input")
    fmt.Println("Try again")
    fmt.Println("")
    main()
  }
} // end main
