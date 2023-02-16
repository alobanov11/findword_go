default:
	cat urls.txt | go run cmd/findword/main.go -w 'Go' -t 5

# make find w=Go t=5
find:
	cat urls.txt | go run cmd/findword/main.go -w '$(w)' -t $(t)