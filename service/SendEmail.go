package service

import (
	"Robo/skimas"
	"log"
	"net/smtp"
	"strings"
)

func SendEmail(Body skimas.WhoisData) {
	username := "usuario mailtrap"
	password := "senha mailtrap"
	smtpHost := "sandbox.smtp.mailtrap.io"

	auth := smtp.PlainAuth("", username, password, smtpHost)

	from := "mail@gmail.com"
	to := []string{Body.Email}

	message := strings.Join([]string{
		"To: " + Body.Email,
		"From: " + from,
		"domain: " + Body.Domain,
		"Subject: Testando email",
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"utf-8\"",
		"",
		`<html>
		<head>
			<style>
				body {
					font-family: Inter, Arial, sans-serif;
				}
				img {
					width: 200px;
				}
			</style>
		</head>
		<body>
			<p>Olá pessoal, boa tarde!</p>
			<p>Tudo bem?</p>
			<p>Podemos alertar o(s) novo(s) concorrente(s) abaixo? É importante ter a <strong>negativação do lado de vocês também.</strong> Podem nos enviar o print, por gentileza?</p>
			<p><a href="` + Body.Domain + `">levoutec</a></p>
			<p>Ficamos no aguardo.</p>
			<p>Equipe BrandMonitor</p>
			<img src="https://uploads-ssl.webflow.com/605b962d5e846a3de31701a8/605b96ae6ffbf350f1722434_Logo%20Brand%20Monitor.svg" alt="Logo BrandMonitor">
		</body>
		</html>`,
	}, "\r\n")

	smtpUrl := smtpHost + ":25"
	err := smtp.SendMail(smtpUrl, auth, from, to, []byte(message))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("URL:", smtpUrl)
	log.Println("E-mail enviado para:", to)
}
