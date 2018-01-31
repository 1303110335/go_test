package test2
 
import "fmt"

func main() {
	var personalInfo PersonalInfo
	 personalInfo = PersonalInfo{"Varun", "Ojha", "dob", "varun@gmail.com", "9999999999"}
	 bytes, err := json.Marshal (&personalInfo)
	 if err != nil {
	        fmt.Println("Could not marshal personal info object", err)
	        return nil, err
	 }
	 err = stub.PutState("key", bytes)
}

