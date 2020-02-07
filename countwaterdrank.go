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
)

func main() {

    fmt.Println("hi there! Welcome Count the Number of Water Cups you drank App\n");
    fmt.Println("Pls press 'a' if you want to update the # of glasses or 'v' to view your consumption\n");

    reader := bufio.NewReader(os.Stdin);
    char, _, err := reader.ReadRune();

    if err != nil {
        fmt.Println(err);
    }
    
    switch char {
    case 'a':
        fmt.Println("Drank One more cup\n");
        break;
    case 'v':
        fmt.Println("Unsupported Functionality\n"); 
        break;
    }
    
/***
	length, str := getstring("kannan b!");
	fmt.Println("str = . len(str) = ", length, str);
	_, str1 := getstring("Kannan");
	fmt.Println("str = ", str1);
***/

}

/*****
func  getstring(str string)  (int, string) {
   return len(str), str;
}
****/
