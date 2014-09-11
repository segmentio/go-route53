package main

import "github.com/segmentio/go-route53"
import "github.com/mitchellh/goamz/aws"
import "encoding/json"
import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	auth, err := aws.EnvAuth()
	check(err)

	dns := route53.New(auth, aws.USWest2)
	check(err)

	res, err := dns.Zone("Z3T864J4ZMBODE").RecordsByName("foo.test.io")
	check(err)

	b, err := json.MarshalIndent(res, "", "  ")
	check(err)

	os.Stdout.Write(b)
}
