package main

import (
    ."fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    
    Println("Введите пример, используя римскую или арабскую нумерацию")
    Println("Для вычесления используйте - + / * и числа от 1 до 10 или от I до X")

    var a, b, symbol string 
    var countLatin, countArab, d, c int

                 
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    slc := strings.Split(scanner.Text(), " ")
    if len(slc) > 3 {
        Println("На вход принимаются только два числа и математический символ между ними!")
        return
    }
   
    a = slc[0]
    symbol = slc[1]
    b = slc[2]

                               
    
    arrayArab := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
      
    arrayRoman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", ""} // "XI", "XII", "XIII", 
		// "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX"}
    
    
       for idx, elem := range arrayRoman {
        if elem == a {
            d = idx + 1
           countLatin++ 
        } 
        if elem == b {
            c = idx + 1
           countLatin++ 
        }
    }      
        
                        
    for idx, elem := range arrayArab {
        if elem == a {
            d = idx + 1
           countArab++ 
        } 
        if elem == b {
            c = idx + 1
           countArab++ 
        }
    }     
           

        
    result := 0  
    if countLatin == 1 && countArab == 1 {
        Print("Вы не можете одновременно использовать римские и арабские цифры!")
    } else  if countArab == 2 || countLatin == 2 {       
        switch symbol {
        case "+":
            result = d + c
        case "-":
            result = d - c
        case "*":
            result = d * c
        case "/":
            result = d / c
         default:
            Println("Не возможно провести операцию")
            Println("Проверьте соответствует ли условие инструкциям")
        } 
    } else {
            Println("Не возможно провести операцию")
            Println("Проверьте соответствует ли условие инструкциям") 
        }
      
     
    
    if countLatin == 2 {
        romanResult := "" 
        search := result%10 - 1
        if search < 0 {
            search = 10
        }
        if result <= 0 {
            Println("В римской системе счисления отсутствуют отрицательные числа или ноль")
        } else {     
            cicleValue := (result/10)%10  
            if result > 10 && result < 40 {               
            
                for i:=0; i<(cicleValue); i++ {
                    romanResult += "X"
                }
                romanResult += arrayRoman[search]
            }else if result >=40 && result < 50 {

                romanResult += "XL"
                romanResult += arrayRoman[search]   
            
            } else if result >= 50 && result < 90 { 
                      
                romanResult += "L"
                
                cicleValue50 := ((result-50)/10)%10 

                for i:=0; i<(cicleValue50); i++ {
                romanResult += "X"
                }
                romanResult += arrayRoman[search]

            } else if result >= 90 && result < 100 {
                romanResult += "XC"
                romanResult += arrayRoman[search]

            } else if result == 100 {
                romanResult = "C"  
            } else {
                romanResult = arrayRoman[search]
            }                       
		    Print("= ")
            Println(romanResult)  
            }
        
    } else {
		Print("= ")
        Println(result)  
    }

        
}
























