package main

import (
	"fmt"
	"strings"
)

var conferenceName = "YagandaClub"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings [50]string
var bookings = []string{}

func main() {
	greetUsers()

	/* Get Variable Types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)
	*/

	// fmt.Println("Welcome to", conferenceName, "our conference booking application.")
	// fmt.Printf("Welcome to %v our conference booking application.\n", conferenceName)

	// fmt.Println("We have total of", conferenceTickets, "ticket and", remainingTickets, "are still avaible.")

	for {
		/*
			for remainingTickets > 0 && len(bookings) < 50 { // statement }
			for { // this code will repeat endlessly }
			for true { // infinitive }
		*/

		firstName, lastName, email, userTickets := getUserInput()

		// strings.Split("yaganda@club.com", "@")[0]
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			/* _ definitons for unused variables */
			// call func print firs names
			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our confrence is booked out. Come back next year.")
				break
			}
		} else {
			/*
				else if userTickets == remainingTickets {
					// do something else
				}
			*/

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			// continue

			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}

			// fmt.Println("Your input data is invalid, try again.")
		}
	}

	/*
		city := "London"

		switch city {
		case "New York":
			// execute code for booking New York conference tickets.
		case "Singapore", "Berlin", "Mexico City", "Hong Kong":
			// some code here for London & Berlin & M.C. & Singapore...
		default:
			fmt.Println("No valid city selected.")
		}
	*/
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still avaible.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(
	firstName string,
	lastName string,
	email string,
	userTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Printf("Enter your first name: ")
	fmt.Scanln(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scanln(&lastName)

	fmt.Printf("Enter your email name: ")
	fmt.Scanln(&email)

	fmt.Printf("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	/* Type the memory address of the variable.
	fmt.Println(&remainingTickets)
	*/

	/*
		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first array: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice type: %v\n", len(bookings))
	*/

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
