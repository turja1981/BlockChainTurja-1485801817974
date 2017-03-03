/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	//"strconv"
	//"strings"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	//"regexp"
)

var logger = shim.NewLogger("ClaimChaincode")

const   CLAIM_SOURCE      	=  "claimsource"
const   INSURANCE_COMPANY 	=  "insuranceCompany"
const   ADJUSTER  			=  "adjuster"
const   BANK 				=  "bank"



// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


type Vehicle struct {
	Make            string `json:"make"`
	Model           string `json:"model"`
	VIN             string `json:"vin"`
	Year           	string `json:"year"`
//	LicensePlate	string `json:"licensePlate"`
}


type Loss struct {
	LossType            	string `json:"lossType"`
	LossDateTime            string `json:"lossDate"`
	LossDescription     	string `json:"lossDescription"`
	LossAddress         	string `json:"lossAddress"`
	LossCity            	string `json:"lossCity"`
	LossState	    		string `json:"lossState"`

}

type Insured struct {
	FirstName              string `json:"firstName"`
	LastName           	   string `json:"lastName"`
	PhoneNo         	   string `json:"phoneNo"`
	Email           	   string `json:"email"`
	Dob             	   string `json:"dobb"`
	DrivingLicense         string `json:"DrivingLicense"`
}

type Property struct {
	PropertyAddress            	string `json:"PropertyAddress"`
	PropertyCity         		string `json:"PropertyCity"`
	PropertyState           	string `json:"PropertyState"`
	PropertyZip             	string `json:"PropertyZip"`
	IfRoofDamaged       		string `json:"ifRoofDamaged"`
	IfLightingCausedFire       	string `json:"ifLightingCausedFire"`
}


type Claim struct {
	 
	ClaimId	    		string		`json:"claimId"` 
	PolicyNo			string		`json:"policyNo"` 
	ClaimNo	    		string		`json:"claimNo"`
	EstmLossAmount		string		`json:"estmLossAmount"` 
	LossDetails 		Loss 		`json:"lossDetails"`
	InsuredDetails 		Insured 	`json:"insuredDetails"`
	VehicleDetails 		Vehicle 	`json:"vehicleDetails"`
//	PropertyDetails 	Property 	`json:"propertyDetails"`
}





// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	
	var claim Claim
	
	bytes, err := json.Marshal(claim)
	
	if err != nil { return nil, errors.New("Error creating V5C_Holder record") }

	err = stub.PutState("claimNo", bytes)
	
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	var c Claim

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	} else if function == "add_fnol" {
        return t.add_fnol(stub, c)
   	} else {
		
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	 claimNo := args[1]
	// Handle different functions
	if function == "retrieve_Claim" { 
		    c , err := t.retrieve_Claim(stub, claimNo) 
			bytes, err := json.Marshal(c)                         //read a variable
	        return bytes , err
    	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *SimpleChaincode) add_fnol(stub shim.ChaincodeStubInterface, claimObj Claim) ([]byte, error) {
   
    var err error
    fmt.Println("running add_fnol()")

    _ ,err = t.save_changes(stub, claimObj)
     
    if err != nil {
        return nil, err
    }
    return nil, nil
}

func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    var key, jsonResp string
    var err error

    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
    }

    key = args[0]
    valAsbytes, err := stub.GetState(key)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
        return nil, errors.New(jsonResp)
    }

    return valAsbytes, nil
}

//==============================================================================================================================
// save_changes - Writes to the ledger the Vehicle struct passed in a JSON format. Uses the shim file's
//				  method 'PutState'.
//==============================================================================================================================
func (t *SimpleChaincode) save_changes(stub shim.ChaincodeStubInterface, c Claim) (bool, error) {

	bytes, err := json.Marshal(c)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error converting vehicle record: %s", err); return false, errors.New("Error converting claim record") }

	err = stub.PutState(c.ClaimId, bytes)

	if err != nil { fmt.Printf("SAVE_CHANGES: Error storing vehicle record: %s", err); return false, errors.New("Error storing claim record") }

	return true, nil
}

func (t *SimpleChaincode) retrieve_Claim(stub shim.ChaincodeStubInterface, claimNo string) (Claim, error) {

	var c Claim

	bytes, err := stub.GetState(claimNo);

	if err != nil {	fmt.Printf("RETRIEVE_claimId: Failed to invoke vehicle_code: %s", err); return c, errors.New("RETRIEVE_V5C: Error retrieving vehicle with v5cID = " + claimNo) }

	err = json.Unmarshal(bytes, &c);

    if err != nil {	fmt.Printf("RETRIEVE_claimId: Corrupt vehicle record "+string(bytes)+": %s", err); return c, errors.New("RETRIEVE_V5C: Corrupt vehicle record"+string(bytes))	}

	return c, nil
}
