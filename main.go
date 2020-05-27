package main

import(
    "bufio"
    "encoding/csv"
    // "encoding/json"
    "fmt"
    "io"
    "log"
	"os"
  "strings"
  "os/exec"
)



type Person struct {
    sku string `json:"firstname"`
    Name string  `json:"lastname"`
	description string `json:"description"`
	product_type string `json:"product_type"`
	stock string `json:"stock"`
	condition string `json:"condition"`
	link string `json:"link"`
	ImageURL string `json:"ImageURL"`
	brand string `json:"brand"`
	age_group string `json:"age_group"`
	color string `json:"color"`
	size string `json:"size"`
	shipping string `json:"shipping"`
	price string `json:"price"`
	id string  `json:"id"`
	custom_label_0 string `json:"custom_label_0"`
	gender string `json:"gender"`
	Sale_price string `json:Sale_price`
	Identity string `json:"Identity"`
	Imageurl string `json:"stock"`
	google_product_category string `json:"google_product_category"`
	availability string `json:"availability"`
	R1 string `json:"R1"`
	R2 string `json:"R2"`
	R3 string `json:"R3"`
	R4 string `json:"R4"`
}

// sku	Name	description	product_type	stock	condition	link	ImageURL	brand	age_group	color	size	shipping	price	id	custom_label_0	gender	Sale_price	Identity	Imageurl	google_product_category	availability	R1	R2	R3	R4	R5	R6	R7	R8	R9	R10	R11	R12	R13	R14	R15
//23-26 column
// type Address struct {
//     City  string `json:"city"`
//     State string `json:"state"`
// }

func main() {
// cmd := exec.Command("ls").Output()

c := exec.Command("python csvdownload.py -a 8W8-K66-675Z -c SCE-RIY-OTKL -pjson /Users/piyushdutta/Documents/kalaki_fashion/query.json -pcsv /Users/piyushdutta/Documents/kalaki_fashion/open.csv -t event")

if err := c.Run(); err != nil { 
    fmt.Println("Error: ", err)
}   

    csvFile, _ := os.Open("kalki_recommendations.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
    var people []Person
    for {
		line, error := reader.Read()
		// fmt.Println(record[17:22])
        if error == io.EOF {
            break
        } else if error != nil {
			log.Fatal(error)
			
		}
		
		if(line[0] == "341553"){
			fmt.Println(line[22]+ ` `+ line[23] +` `+ line[24]+ ` `+ line[25]+` `+ line[26])
			//function which will return the sku id
			fmt.Printf("Fields are: %q", strings.Fields(line[23]+` `+line[24] + ` `+line[25] + ` `+ line[26]))
			// value2 := strings.Fields(line[23]+` `+line[24] + ` `+line[25] + ` `+ line[26])
      // fmt.Println(value2)
      
    //  func (strings receommendations){
			 
		// 	 //Imageurl, price , Name and link (pick only four values and update it as user property)
		//  }

		for {
		row1, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
			log.Fatal(error)
			
		}
		if(row1[0]==line[23]){
			// value1 := line
			// fmt.Println(value1)
		} 
		
		}

		}
        people = append(people, Person{
      sku: line[0],
      Name:  line[1],
			description: line[2],
			product_type: line[3],
			stock: line[4],
			condition: line[5],
			link: line[6],
			ImageURL: line[7],
			brand: line[8],
			age_group: line[9],
			color: line[10],
			size: line[11],
			shipping: line[12],
			price: line[13],
			id: line[14],
			custom_label_0: line[15],
			gender: line[16],
			Sale_price: line[17],
			Identity: line[18],
			Imageurl: line[19],
			google_product_category: line[20],
			availability: line[21],
		})
			

    }

}

//app scripts 