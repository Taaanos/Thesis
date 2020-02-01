func (t *ExampleChaincode) read(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if args[1] == "hello"	{
		// Get the state of the value matching the key "key" in the ledger. It should be "value"
		state, err := stub.GetState("key")
		if err != nil {
			return shim.Error("Failed to get state hello")
		}

		// return the value corresponding to the request as a response
		return shim.Success(state)
	}

	return shim.Error("unknown query action, check the second arg")
}

func (t *ExampleChaincode) write(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	if args[1] == "key" && len(args) == 3 {
		// Write the new value in the ledger
		err := stub.PutState("key", []byte(args[2]))

		if err != nil {
			return shim.Error("Failed to update state of key")
		}

		return shim.Success(nil)
	}
	return shim.Error("Unknown invoke action, check the second argument.")
}