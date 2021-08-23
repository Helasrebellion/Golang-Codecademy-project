package main

import "fmt"

func main() {
  var publisher, writer, artist, title string
  var year uint
  var pageNumber uint
  var grade float32

//orginal var assigns will go here
title = "Mr. GoToSleep"
writer = "Tracey Hatchet"
artist = "Jewel Tampson"
publisher = "DizzyBooks Publishing Inc."
year = 1997
pageNumber = 14
grade = 6.5



  fmt.Println(title, "written by ", writer, "drawn by ", artist)
  fmt.Println("Published by ", publisher, "released in ", year, "It has", pageNumber, "pages", "and was given a grade of", grade)

  //Assigns will go here
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, "written by ", writer, "drawn by ", artist)
  fmt.Println("Published by ", publisher, "released in ", year, "It has", pageNumber, "pages", "and was given a grade of", grade)

}
