package webhook

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/ws"
)

func PostMessage(c *fiber.Ctx) error {
	var h Header
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}
	var resp atmessage.Response
	if h.Secret == WebhookSecret {
		var msg model.IteungMessage
		err = c.BodyParser(&msg)
		if err != nil {
			return err
		}

		// Check if the message is a login request
		if ws.IsLoginRequest(msg, WAKeyword) {
			// Check if the phone number is registered
			registered, err := isPhoneNumberRegistered(msg.Phone_number)
			if err != nil {
				// Handle the error here, for now, we'll log it
				log.Println("Error checking phone number registration:", err)
				// You might want to set a response or return an error response to the user
				resp.Response = "An error occurred while checking phone number registration."
			} else if registered {
				resp = HandlerQRLogin(msg, WAKeyword)
			} else {
				// Phone number is not registered, handle accordingly
				resp.Response = getRandomUnregisteredNumber()
			}
		} else if msg.Message == "minta test" {
			// Fetch user data from db_bot
			nama, email, err := FetchUserData(msg.Phone_number)
			if err != nil {
				// Handle the error here, for now, we'll log it
				log.Println("Error fetching user data:", err)
				resp.Response = getRandomIncorrectSecretMessage()
			} else {
				resp.Response = fmt.Sprintf("Nama: %s\nE-Mail: %s", nama, email)
			}
		} else {
			// Handle incoming message for non-login requests
			resp = HandlerIncomingMessage(msg)
		}
	} else {
		// Random response for incorrect secret code
		resp.Response = getRandomIncorrectSecretMessage()
	}
	return c.JSON(resp)
}

// Randomize response if secret code was invalid
func getRandomIncorrectSecretMessage() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of possible incorrect secret messages
	incorrectSecretMessages := []string{
		"Oops! It seems like the secret code is incorrect.",
		"Sorry, but the provided secret code is not valid.",
		"Hmm, something went wrong. Please check the secret code and try again.",
		"Error: Incorrect secret code. Please verify and resend your message.",
		"Ups! kode loginnya salah nihh.",
		"Maaf kak, kode kodenya ga valid :(",
		"Hmm, coba cek kode loginnya lagi deh...",
		"BANGG UDAH BANGG, KODENYA SALAHHH",
	}

	// Randomly select a message
	return incorrectSecretMessages[rand.Intn(len(incorrectSecretMessages))]
}

// Randomize response if number is not registered
func getRandomUnregisteredNumber() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of possible incorrect secret messages
	unregisteredNumberMessages := []string{
		"Oops! It seems like you are not registered!",
		"Sorry, but you are not registered yet!",
		"Hmm, something went wrong. Please check the secret code and try again.",
		"Duhh, sepertinya kakak masih belum daftar deh...",
		"Kakak daftar di web kami dulu yaaa....",
		"Aduh... kami ga nemuin data kakak :( Daftar dulu yaa",
		"LAH LU SIAPA BANG???!",
		"Ndeh... sia ko?",
	}

	// Randomly select a message
	return unregisteredNumberMessages[rand.Intn(len(unregisteredNumberMessages))]
}
