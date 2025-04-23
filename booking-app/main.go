/* Run Command
go run main.go helper.go
or
go run .
*/

package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "YagandaClub"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings [50]string
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	/* Get Variable Types
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName %T\n", conferenceTickets, remainingTickets, conferenceName)
	*/

	// fmt.Println("Welcome to", conferenceName, "our conference booking application.")
	// fmt.Printf("Welcome to %v our conference booking application.\n", conferenceName)

	// fmt.Println("We have total of", conferenceTickets, "ticket and", remainingTickets, "are still avaible.")

	/*
		for remainingTickets > 0 && len(bookings) < 50 { // statement }
		for { // this code will repeat endlessly }
		for true { // infinitive }
	*/

	firstName, lastName, email, userTickets := getUserInput()

	// strings.Split("yaganda@club.com", "@")[0]
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1) // Sets the number of goroutines to wait for (increases the counter by the provided number)
		// go keywork important (multiple thread). 10 second waiting...
		go sendTicket(userTickets, firstName, lastName, email)

		/* _ definitons for unused variables */
		// call func print firs names
		fmt.Printf("The first names of bookings are: %v\n", getFirstNames())

		// noTicketsRemaining := remainingTickets == 0
		if remainingTickets == 0 {
			// end program
			fmt.Println("Our confrence is booked out. Come back next year.")
			// break
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

	wg.Wait() // Blocks until the WaitGroup counter is 0.

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
		// firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
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

	// create a map for a user
	// var userData = make(map[string]string)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	/*
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email

		Converting Schema
		2 : binary
		8 : octal
		10 : decimal
		16 : hex(for letters)
		1 or 2 : invalid value... (error)

		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	*/

	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

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

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%b ticket for %v %v", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending ticket:\n %v\n to email address %v", ticket, email)
	fmt.Println("######################")

	wg.Done() // Decrements the WaitGroup counter bu 1 So this is called by the goroutine to indicate that it's finished.
}
