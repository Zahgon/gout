package main

func genFormCurl() {
	_ = "STUB: not implemented"
	// 1.formdata
	return
}

// output:
// curl -X GET -F "text=good" -F "mode=A" -F "voice=@./voice" "http://127.0.0.1:1234"

func genJSONCurl() {
	_ = "STUB: not implemented"
	// 2.json body
	return
}

// output:
// curl -X GET -H "Content-Type:application/json" -d "{\"key1\":\"val1\",\"key2\":\"val2\"}" "http://127.0.0.1:1234"

func main() {
	genFormCurl()
	genJSONCurl()
}
