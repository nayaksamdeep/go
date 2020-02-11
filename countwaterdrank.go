/*
 * This program measures the number of glasses of water a person drinks a day
 * It will an user to input the number of glasses he has drank
 * It will also show him the total number of cups he has consumed for the day
 */

package main

import  (
     "fmt"
     "bufio"
     "os"
     "log"
     "time"
     "github.com/asdine/storm"
)


type Entry struct {
        ID       int `storm:"id,increment"`
	Cups     int   
	Day      time.Time
}

func main() {

    db, err := storm.Open("test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    fmt.Println("hi there! Welcome to Hydrometer where you can track the Number of Water Cups you drank\n");
    fmt.Println("Pls press 'A' if you want to update the # of glasses or 'V' to view your consumption\n");

    reader := bufio.NewReader(os.Stdin);
    char, _, err := reader.ReadRune();

    if err != nil {
        fmt.Println(err);
    }
    
    switch char {
    case 'A':
        // Entries for water consumption
        err = addEntry(db)
        if err != nil {
                log.Fatal(err)
        }
        break;
    case 'V':
        var today []Entry
        var total int

        err = db.Range("Day", time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 1), &today)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println(today)

        /*
         * TODO - Figure out the total number of records using Storm
         */
        total = today[1].Cups + today[0].Cups

        fmt.Println(total)

        break;
    default:
        fmt.Print("Pls type in valid input\n");
    }
    
}

func addEntry(db *storm.DB) error {
    /*
     * Check if the user wants to update the count for today or wants to correct the count for previous day
     */
     fmt.Println("Do you want to update the count for today (press 0) or previous day (press 1 for yesterday, 2 for Day before etc\n"); 

    var day, cups int

    fmt.Scanf("%d", &day)

    /*
     * Enter the number of cups
     */
     fmt.Println("Enter the number of cups you drank\n"); 

    fmt.Scanf("%d", &cups)


    /*
     * Save the number into DB
     */ 
    entry := Entry{Cups: cups, Day:time.Now().AddDate(0, 0, -day) }
    err := db.Save(&entry)
    if err != nil {
          return fmt.Errorf("could not save entry, %v", err)
    }
    fmt.Println("entry saved")
    return nil
}

