run:
	go run ./app/main.go

hello:
	echo "katerina, Te amo! 😍"

git-push:
	@read -p "Elias, enter your commit message: " name; \
		git add .; \
		git commit -m "$$name"; \
		git push origin main; \
			echo "Yeap Elias your code has been pushed! i love k😍"