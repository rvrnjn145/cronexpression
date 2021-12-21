package main

import (
	"./CronExpression"
	"log"
)

func main() {

	/*if len(os.Args)!=2{
		log.Printf("invalid number of agruments")
		return

	}*/
	//expression := "*/15 0 1,15 * 1-5 /usr/bin/find -e foo we we"
	//expression:="10/2 0 1,15 * 1-5 /usr/bin/find -e foo"
	//expression:="20-30/2 0 1,15 * 1-5 /usr/bin/find -e foo"
	expression := "1,2,4-5,10-20/2 0 1,15 * 1-5 /usr/bin/find -e foo"
	cronExpression.Init()
	_, err := cronExpression.DefaultClient.CronExpressionDescriptor(expression)

	if err != nil {
		log.Printf("failed to convert CRON expression to human readable description: %s", err)
	}
}
