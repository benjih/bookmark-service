run:
	go run main.go

test:
	echo "GET bookmarks"
	echo "Empty response service has just started"
	curl --header "Content-Type: application/json" \
		--request GET \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "POST bookmark dataset"
	curl --header "Content-Type: application/json" \
		--request POST \
		 --data @bookmarks.json \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "GET bookmarks"
	curl --header "Content-Type: application/json" \
		--request GET \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "DELETE bookmarks by URL"
	curl --header "Content-Type: application/json" \
		--request DELETE \
		--data '{"urls": ["https://somesite.com"]}' \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "GET bookmarks"
	curl --header "Content-Type: application/json" \
		--request GET \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "DELETE bookmarks by URL"
	curl --header "Content-Type: application/json" \
		--request DELETE \
		--data '{"urls": ["https://somesites.com"]}' \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "GET bookmarks"
	curl --header "Content-Type: application/json" \
		--request GET \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "DELETE bookmarks by URL"
	curl --header "Content-Type: application/json" \
		--request DELETE \
		--data '{"urls": ["https://somesite2.com", "https://somesite4.com", "https://somesite3.com"]}' \
		http://localhost:3000/api/bookmarks/
	echo ""
	echo "GET bookmarks"
	curl --header "Content-Type: application/json" \
		--request GET \
		http://localhost:3000/api/bookmarks/