package main

import (
    "encoding/json"
    "fmt"
    "math"
    "math/rand"
    "os"
    "time"
)

type UnitFormat struct {
    Unit []GuestFormat
}

type GuestFormat struct {
    Name string `json:"name"`
    Gender string `json:"gender"`
}

type Table struct {
    Seats int
    Guests []GuestFormat
}

func (a *UnitFormat) UnmarshalJSON(b []byte) error {
    var s []GuestFormat
    if err := json.Unmarshal(b, &s); err != nil {
        return err
    }
    a.Unit = s
    return nil
}

func round(float float64) float64 {
    return math.Floor(float + 0.5)
}

func shuffle(slice []UnitFormat) []UnitFormat {
    for i := len(slice) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        slice[i], slice[j] = slice[j], slice[i]
    }
    return slice
}

func flatten(slice []UnitFormat) []GuestFormat {
    result := make([]GuestFormat, 0)
    for i := range slice {
        unit := slice[i].Unit
        result = append(result, unit[0])
        if len(unit) > 1 {
            result = append(result, unit[1])
        }
    }
    return result
}

/*
func grab(amount int, guestList []GuestFormat) []GuestFormat {
    var totBoys, totGirls, totAll int

    tableGuests := make([]GuestFormat, amount)

    for i := 0; totAll < amount; i++ {
        unit := groups[i].Unit
        totAll += len(unit)
    }

    fmt.Println(totBoys, totGirls, totAll, len(tableGuests))

    return tableGuests
}
*/

/*
func sitDownPlease(units []UnitFormat, seats int) []Table {
    tables := make([]Table, 0)
    munits := make([]UnitFormat, len(units))
    copy(munits, units)

    for {
        if len(munits) == 0 {
            break
        }

        // keep creating tables until we don't have any more fools
        var allocated []GuestFormat

        for i := 0; i < seats; {
            if len(munits) == 0 {
                break
            }
            unit := munits[0].Unit

            for _, guest := range unit {
                fmt.Println("Added guest", guest.Name)
                allocated = append(allocated, guest)
                i++
            }

            munits = append(munits[:0], munits[1:]...)
        }

        tables = append(tables, Table{Seats: len(allocated), Guests: allocated})
    }


    return tables
}
*/

func sitDownPlease(units []UnitFormat, seats int) []Table {
    tables := make([]Table, 0)
    munits := make([]UnitFormat, len(units))
    copy(munits, units)

    sorted := 0
    tableCount := 0

    for {
        if len(munits) == 0 {
            break
        }

        var allocated []GuestFormat

        fmt.Println("New table", tableCount+1)
        i := 0
        for {
            fmt.Println("..Current length", len(munits))
            fmt.Println("..Current index", i)
            fmt.Println("..Units", munits)
            fmt.Println("..Seats left", seats - len(allocated))

            if len(allocated) == seats {
                break
            }

            if i == len(munits) {
                break
            }

            unit := munits[i]
            if (len(allocated) + len(unit.Unit)) > seats {
                fmt.Println("/!\\ Hold up /!\\")
                fmt.Println("We're about to run out of chairs, find another unit, bro!")
                fmt.Println("Tried to add", unit.Unit)
                fmt.Print()
                i++
                continue
            }

            for _, guest := range unit.Unit {
                allocated = append(allocated, guest)
                sorted++
                fmt.Println("....Added guest", guest.Name, sorted)
                fmt.Println("....Seats left", seats - len(allocated))
            }

            fmt.Println("..Remove", munits[i])
            munits = append(munits[:i], munits[i+1:]...)
        }

        tableCount++
        fmt.Println()
        tables = append(tables, Table{Seats: len(allocated), Guests: allocated})
    }

    return tables
}

func main() {
    rand.Seed( time.Now().UTC().UnixNano())

    file, err := os.Open("guest-list.json")

    if err != nil {
        panic(err)
    }
    defer file.Close()

    var units = make([]UnitFormat, 0)
    err = json.NewDecoder(file).Decode(&units)

    if err != nil {
        panic(err)
    }

    var guestCount float64
    var maleCount int
    var femaleCount int

    for i := 0; i < len(units); i++ {
        unit := units[i].Unit
        for k := 0; k < len(unit); k++ {
            guestCount += 1

            switch unit[k].Gender {
                case "male": maleCount += 1
                case "female": femaleCount += 1
            }

            fmt.Println(guestCount, unit[k].Name)
        }
    }

    guestsPerTable := 6

    fmt.Println("\nStatistics")
    fmt.Println("----------")

    fmt.Println("Boys\t\t\t", maleCount)
    fmt.Println("Girls\t\t\t", femaleCount)
    fmt.Println("")
    fmt.Println("Guests per table:\t", guestsPerTable)
    fmt.Println("")

    //guestList := flatten(units)
    units = shuffle(units)

    tables := sitDownPlease(units, guestsPerTable)

    for i, table := range tables {
        boys := 0
        girls := 0

        fmt.Println()
        fmt.Printf("Table %d", i + 1)
        fmt.Println()

        for k, guest := range table.Guests {
            fmt.Printf("\t%d. %s (%s)\n", k + 1, guest.Name, guest.Gender)
            switch guest.Gender {
                case "male": boys++
                case "female": girls++
            }
        }

        fmt.Printf("M: %d, F: %d", boys, girls)
        fmt.Println()
    }
}
