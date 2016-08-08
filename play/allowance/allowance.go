package main 

import(
	"fmt"
	"time"
	"math"
	)

func main() {
	seattle, _ := time.LoadLocation("America/Los_Angeles")
	detroit, _ := time.LoadLocation("America/Detroit")
	denver, _ := time.LoadLocation("America/Denver")
	jems := time.Date(2005,04,14,1,0,0,0,seattle)
	beans := time.Date(2007,04,30,11,0,0,0,detroit)
	rivers := time.Date(2016,05,22,22,0,0,0,denver)

	sinceSunday := time.Now().Weekday() * 24
	sinceMidnight := time.Now().Hour()

	jemshours := time.Since(jems).Hours()
	beanshours := time.Since(beans).Hours()
	rivershours := time.Since(rivers).Hours()

	jemsweeks := ((jemshours - float64(sinceSunday) - float64(sinceMidnight)) / 24) / 7
	beansweeks := ((beanshours - float64(sinceSunday) - float64(sinceMidnight)) / 24) / 7
	riversweeks := ((rivershours - float64(sinceSunday) - float64(sinceMidnight)) / 24) / 7

	fmt.Printf("Jem's allowance this week is $%3.2f.\n", math.Pow(10,(jemsweeks/521.79)))
	fmt.Printf("Bean's allowance this week is $%3.2f.\n", math.Pow(10,(beansweeks/521.79)))
	fmt.Printf("River's allowance this week is $%3.2f.\n", math.Pow(10,(riversweeks/521.79)))
}
