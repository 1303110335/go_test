package test2
 
import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
 
type SampleChaincode struct {
}
 
func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SampleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func main() {
    err := shim.Start(new(SampleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}


func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    fmt.Println("Entering CreateLoanApplication")
 
    if len(args) < 2 {
        fmt.Println("Invalid number of args")
        return nil, errors.New("Expected at least two arguments for loan application creation")
    }
 
    var loanApplicationId = args[0]
    var loanApplicationInput = args[1]
 
    err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
    if err != nil {
        fmt.Println("Could not save loan application to ledger", err)
        return nil, err
    }
 
    fmt.Println("Successfully saved loan application")
    return nil, nil
}


func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    fmt.Println("Entering GetLoanApplication")
 
    if len(args) < 1 {
        fmt.Println("Invalid number of arguments")
        return nil, errors.New("Missing loan application ID")
    }
 
    var loanApplicationId = args[0]
    bytes, err := stub.GetState(loanApplicationId)
    if err != nil {
        fmt.Println("Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
        return nil, err
    }
    return bytes, nil
}