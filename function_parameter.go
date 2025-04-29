package main

import "fmt"
import _"strings"

func funtionParameter(){
    var data = []string{"apel", "jeruk", "anggur", "pisang"}
    var dataContainsO = filter(data, func(each string) bool {
        return contains(each, "apel")
    })

    fmt.Println("data asli \t\t:", data)
    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
}

func contains(s, substr string) bool{
    if substr == ""{
        return true;
    }
    // contains("hello", "lo");
    for i := 0; i <= len(s) - len(substr); i++ {
        match := true
        for j := 0; j < len(substr); j++ {
            if s[i+j] != substr[j]{ 
                match = false
                break
            }
        }
        if match{
            return true
        }
    }

    return false
}

type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string{
    var result []string;
    
    for _, each := range data {
        if filtered := callback(each); filtered{
            result = append(result, each)
        }
    }
    
    return result;
}





// func filter(data []string, callback func(string) bool) []string{
//     var result []string

//     for _, each := range data {
//         if filtered := callback(each); filtered{
//             result = append(result, each)
//         }
//     }

//     return result
// }

// func contains(s, substr string) bool{
//     // Jika substring kosong, dianggap selalu ditemukan
//     if substr == "" {
//         return true
//     }
    
//     for ()
//     return false;
// }
