 package main
 import "fmt"

var firstUserChoice int
var secondUserChoice int



func main()  {
   fmt.Println("Welcome to GO METER app")
   fmt.Print("1-Milimeter to\n2-Centimeter to\n3-Meter to\n4-Kilometer to\nPlease choose an operation: ")
   fmt.Scanln(&firstUserChoice)
   if firstUserChoice <= 4 {
     secondStage()
   } else {
     fmt.Println("INVALID CHOICE")
     secondStage()
   }
}

// ask user for the second unit
func secondStage() ()  {
  fmt.Print("1-to Milimeter\n2-to Centimeter\n3-to Meter\n4-kilometer to\nPlease choose an operation: ")
  fmt.Scanln(&secondUserChoice)
  if secondUserChoice <= 4 {
    thirdStage()
  } else {
    fmt.Println("INVALID CHOICE")
    thirdStage()
  }
}

func thirdStage()  {
  var number float64
  const mm string = "Milimeters"
  const cm string = "Centimeters"
  const m string = "Meters"
  const km string = "Kilometers"
  fmt.Print("Enter the number you want to convert: ")
  fmt.Scanln(&number)
  switch firstUserChoice {
  case 1:
    if secondUserChoice == 1 {
      fmt.Println(number, mm,"Equals", number, mm)
    }
    if secondUserChoice == 2 {
      fmt.Println(number, mm,"Equals", number/10, cm)
    }
    if secondUserChoice == 3 {
      fmt.Println(number, mm,"Equals", number/1000, m)
    }
    if secondUserChoice == 4 {
      fmt.Println(number, mm,"Equals", number/1000000, km)
    }
  case 2:
    if secondUserChoice == 1 {
      fmt.Println(number, cm,"Equals", number*10, m)
    }
    if secondUserChoice == 2 {
      fmt.Println(number, cm,"Equals", number, cm)
    }
    if secondUserChoice == 3 {
      fmt.Println(number, cm,"Equals", number/100, m)
    }
    if secondUserChoice == 4 {
      fmt.Println(number, cm,"Equals", number/100000, km)
    }
  case 3:
    if secondUserChoice == 1 {
      fmt.Println(number, m,"Equals", number*1000, mm)
    }
    if secondUserChoice == 2 {
      fmt.Println(number, m,"Equals", number*100, cm)
    }
    if secondUserChoice == 3 {
      fmt.Println(number, m,"Equals", number, m)
    }
    if secondUserChoice == 4 {
      fmt.Println(number, m,"Equals", number/1000, km)
    }
  case 4:
    if secondUserChoice == 1 {
      fmt.Println(number, km,"Equals", number*1000000, mm)
    }
    if secondUserChoice == 2 {
      fmt.Println(number, km,"Equals", number*100000, cm)
    }
    if secondUserChoice == 3 {
      fmt.Println(number, km,"Equals", number*1000, mm)
    }
    if secondUserChoice == 4 {
      fmt.Println(number, km,"Equals", number, km)
    }
  }
}
