//Soal Perulangan & Percabangan Brankas
package main

import "fmt"

func main() {
	const pin = "060125"
	var inputPin string

	fmt.Println("Selamat datang di Brankas Sigma!")

	for attempts := 3; attempts > 0; attempts-- {
		fmt.Print("Masukkan PIN Anda: ")
		fmt.Scanln(&inputPin) // Membaca input langsung ke variabel

		if inputPin == pin {
			fmt.Println("Akses Diterima! Brankas Terbuka")
			fmt.Print(`               *               
             *   *             
           *   *   *           
         *       *       *     
       *   *   *****   *   *   
     *   *       *       *   * 
   *       *   *****   *       * 
 *   *   *****       *****   *   *
   *   *       *****       *   *  
     *   *   *       *   *   *    
       *   *   *****   *   *      
         *****       *****        
       *   *   *****   *   *      
     *   *   *       *   *   *    
   *   *       *****       *   *  
 *   *   *****       *****   *   *
   *       *   *****   *       *  
     *   *       *       *   *    
       *   *   *****   *   *      
         *       *       *        
           *   *   *   *          
             *   *   *            
               *****              
`)
			return
		} else {
			fmt.Printf("PIN Salah! %d Percobaan Tersisa.\n", attempts-1)
		}
	}

	fmt.Println("Akses Ditolak. Anda Telah Kehabisan Masa Percobaan.")
}
