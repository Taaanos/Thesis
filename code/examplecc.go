package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type ExampleChaincode struct{
}

func main(){
	err := shim.Start(new(ExampleChaincode))
	if err != nil {
		fmt.Printf("Error starting example chaincode")
	}
}

// Instantiation of the chaincode
func (t *ExampleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	// Get the function name and the parameters from the request
	functionName, _:= stub.GetFunctionAndParameters()

	if functionName != "init"{
		return shim.Error("Unkown function call")
	}
	// Put data in the ledger for demo key/value
	err := stub.PutState("key", []byte("value"))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
	}

// The gateway function for invokations.
func (t *ExampleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	functionName, args := stub.GetFunctionAndParameters()


	if functionName == "init" {
		return t.Init(stub)
	} else if functionName == "read" {
		return t.read(stub, args)
	} else if functionName == "write" {
		return t.write(stub, args)
	}

	fmt.Println("invoked unknown functionName named " + functionName)
	return shim.Error("invoked unknown functionName named " + functionName)
}