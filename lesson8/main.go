package main

import (
	"context"
	"flag"
	"log"
	"time"
)

func main() {
	var df dirFiles

	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	err := df.scanDir()
	for err != nil {
		log.Printf(`The directory "%s" doesn't exist, please, try again.`, df.dir)
		err = df.scanDir()
	}

	err = df.walkDir()
	if err != nil {
		log.Fatal(checkError("Couldn't walk the directory."))
	}

	err = df.findDuplicates()
	if err != nil {
		log.Fatal(checkError("Error finding duplicates in the directory."))
	}

	err = df.deleteDuplicates(ctx)
	for err != nil {
		log.Println(err)
		err = df.deleteDuplicates(ctx)
	}

	err = df.copyOriginals(ctx)
	for err != nil {
		log.Println(err)
		err = df.copyOriginals(ctx)
	}
}
