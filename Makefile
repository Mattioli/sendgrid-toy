.PHONY: send
send:
	go run cmd/sendgrid-toy/main.go send $(fromName) $(fromEmail) $(toName) $(toEmail) $(subject) $(msg)

.PHONY:
prt:
	@echo argument is $(fromEmail) $(fromName)

