
package main

import (
    ."fmt"
)

func main() {
    
    Println("Введите пример, используя римскую или арабскую нумерацию")
    Println("Для вычесления используйте - + / * и числа от 1 до 10 или от I до X")

    var a, b, symbol string 
    var countLatin, countArab, d, c int
    Scan(&a, &symbol, &b)
    
    arrayArab := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
      
    arrayRoman := [20]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", 
		 "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX"}
    
    
       for idx, elem := range arrayRoman[:10] {
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
      
                        
    romanResult := "" 
    if countLatin == 2 {
        if result <= 0 {
            Println("В римской системе счисления отсутствуют отрицательные числа или ноль")
        } else {      
        romanResult = arrayRoman[result - 1]
		Print("= ")
        Println(romanResult)  
        }
        
    } else {
		Print("= ")
        Println(result)  
    }

        
}









