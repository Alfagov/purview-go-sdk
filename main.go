package main

import "net/url"

func main() {
	test, _ := url.Parse("/test/123?com=123&com_1=123")
	println(test.Host)
	testValues := url.Values{}
	testValues.Add("com2", "123")
	testValues.Add("com_3", "123")

	test1 := test.JoinPath("al", "se", "no")
	test1.RawQuery = testValues.Encode()
	println(test1.String())

}
