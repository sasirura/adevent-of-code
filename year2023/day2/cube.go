package main

import (
    "fmt"
    "log"
    "os"

    "strconv"
    "strings"
)

func parseLinePart2(line string) int {

    if line == "" {
        return 0
    }

    parts := strings.Split(line, ":")

    colorParts := strings.Split(parts[1], ";")

    power := 0
    blue := 0
    red := 0
    green := 0

    for _, game := range colorParts{

        sets := strings.Split(game, ",") 
        for _, set := range sets{

            color := strings.Split(set, " ") 

            cube, err := strconv.Atoi(color[1])

            if err != nil {
                log.Fatal("this should not happend" )
            }

            if color[2] == "red" {
                if red < cube {
                    red = cube
                }
            } else if color[2] == "blue" {
                if blue < cube {
                    blue = cube
                }
            } else if green < cube {
                green = cube
            }
        }
    }

    power = green * blue * red

    return power
}



func parseLine(line string) int {

    if line == "" {
        return 0
    }

    parts := strings.Split(line, ":")
    gameNumber, err := strconv.Atoi(strings.Split(parts[0], " ")[1])

    if err != nil {
        log.Fatal("this should not happend" )
    }

    colorParts := strings.Split(parts[1], ";")


    for _, game := range colorParts{

        sets := strings.Split(game, ",") 

        for _, set := range sets{

            blue := 0
            red := 0
            green := 0

            color := strings.Split(set, " ") 

            cube, err := strconv.Atoi(color[1])

            if err != nil {
                log.Fatal("this should not happend" )
            }

            if color[2] == "red" {

                red += cube

            } else if color[2] == "blue" {

                blue += cube
            } else {

                green += cube
            }
            if blue > 14 || green > 13 || red > 12 {
                gameNumber = 0
            }

        }

    }

    return gameNumber
}

func main(){
    b, err := os.ReadFile("day2.txt")

    if err != nil {
        log.Fatal(err)
    }

    getInput := string(b)
    lines :=  strings.Split(getInput, "\n")

    total := 0
    for _, line := range lines {

        current := parseLine(line)

        total += current
    }

    fmt.Printf("value is %d\n" , total)

    part2, err := os.ReadFile("day2p2.txt")

    if err != nil {
        log.Fatal(err)
    }

    getInput2 := string(part2)
    lines2 :=  strings.Split(getInput2, "\n")

    total2 := 0
    for _, line := range lines2 {

        current := parseLinePart2(line)

        total2 += current
    }

    fmt.Printf("value is %d\n" , total2)
}

