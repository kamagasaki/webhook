package webhook

import (
	"log"
	"math/rand"
	"time"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
	"github.com/gofiber/fiber/v2"
	"github.com/whatsauth/ws"
)

func PostMessage(c *fiber.Ctx) error {
	var msg model.IteungMessage
	var resp atmessage.Response
	var h Header
	err := c.ReqHeaderParser(&h)
	if err != nil {
		return err
	}

	if msg.Message == "Babi" || msg.Message == "Anjing" || msg.Message == "goblok" {
		randm := []string{
			"Ih  kasar amat dahh",
			"I thought" + msg.Message + "was your mom!",
			"APASIIII",
			"Keren lu begitu?",
			"Jangan galak galak dong kak, aku takut tauu",
		}
		dt := &wa.TextMessage{
			To:       msg.Phone_number,
			IsGroup:  false,
			Messages: GetRandomString(randm),
		}
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

	} else if msg.Message == "cantik" || msg.Message == "ganteng" || msg.Message == "cakep" {
		dt := &wa.TextMessage{
			To:       msg.Phone_number,
			IsGroup:  false,
			Messages: getRandomThanksMsg()
		}
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")

	} else {
		randm := []string{
			"Hai Hai Haiii kamuuuui " + msg.Alias_name + "\nrofinya lagi gaadaa \n aku giseuubott salam kenall yaaaa \n Cara penggunaan WhatsAuth ada di link berikut ini ya kak...\n" + link,
			"IHHH jangan SPAAM berisik tau giseu lagi tidur",
			"Kamu ganteng tau",
			"Ihhh kamu cantik banget",
			"bro, mending beliin aku nasgor",
			"Jangan galak galak dong kak, aku takut tauu",
			"Mawar Indah hanya akan muncul dipagi hari, MAKANYA BANGUN PAGI KAK",
			"Cihuyyyy hari ini giseuu bahagiaaa banget",
			"Bercandyaaa berrcandyaaaa",
		}
		dt := &wa.TextMessage{
			To:       msg.Phone_number,
			IsGroup:  false,
			Messages: GetRandomString(randm),
		}
		resp, _ = atapi.PostStructWithToken[atmessage.Response]("Token", os.Getenv("TOKEN"), dt, "https://api.wa.my.id/api/send/message/text")
	} else {
		resp.Response = getRandomIncorrectSecretMessage()
	}
	fmt.Fprintf(w, resp.Response)
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
		"Duhh, sepertinya kakak masih belum daftyar deh...",
		"Kakak daftar di web kami dulu yaaa....",
		"Aduh... kami ga nemuin datya kakak :( Daftar dulu yaa",
		"LAH LU SIAPA BANG???!",
		"Ndeh... sia ko?",
	}

	// Randomly select a message
	return unregisteredNumberMessages[rand.Intn(len(unregisteredNumberMessages))]
}

func getRandomThanksMsg() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of possible incorrect secret messages
	ThankMessages := []string{
		"Eh?? Thank youu!",
		"Kamu jugaaa!",
		"Makasih kaaa.",
		"Maachihh ayanggg.",
		"MANG EAAAAAA",
		"Aelah bang, nagatin bot cantik. Cari cewe asli napa?",
	}

	// Randomly select a message
	return ThankMessages[rand.Intn(len(ThankMessages))]
}

func getRandomRudeResp() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Array of possible incorrect secret messages
	RudeResp := []string{
		"Ih  kasar amat dahh",
		"I thought" + msg.Message + "was your mom!",
		"APASIIII",
		"Keren lu begitu?",
		"Jangan galak galak dong kak, aku takut tauu",
	}

	// Randomly select a message
	return RudeResp[rand.Intn(len(RudeResp))]
}