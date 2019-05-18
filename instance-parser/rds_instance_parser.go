package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/pricing"
)

func getProduct(nextToken string) (interface{}, interface{}, interface{}, *string, bool, error) {
	svc := pricing.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	input := &pricing.GetProductsInput{
		Filters: []*pricing.Filter{
			{
				Field: aws.String("location"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("US East (N. Virginia)"),
			},
		},
		ServiceCode:   aws.String("AmazonRDS"),
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(1),
		NextToken:     aws.String(nextToken),
	}
	result, err := svc.GetProducts(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case pricing.ErrCodeInternalErrorException:
				fmt.Println(pricing.ErrCodeInternalErrorException, aerr.Error())
			case pricing.ErrCodeInvalidParameterException:
				fmt.Println(pricing.ErrCodeInvalidParameterException, aerr.Error())
			case pricing.ErrCodeNotFoundException:
				fmt.Println(pricing.ErrCodeNotFoundException, aerr.Error())
			case pricing.ErrCodeInvalidNextTokenException:
				fmt.Println(pricing.ErrCodeInvalidNextTokenException, aerr.Error())
			case pricing.ErrCodeExpiredNextTokenException:
				fmt.Println(pricing.ErrCodeExpiredNextTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
	}
	switch result.NextToken {
	case nil:
		return nil, nil, nil, nil, false, nil
	default:
		instanceType := result.PriceList[0]["product"].(map[string]interface{})["attributes"].(map[string]interface{})["instanceType"]
		if instanceType == nil {
			nToken := result.NextToken
			return nil, nil, nil, nToken, true, errors.New("InstanceType is nil")
		}
		vcpu := result.PriceList[0]["product"].(map[string]interface{})["attributes"].(map[string]interface{})["vcpu"]
		memory := result.PriceList[0]["product"].(map[string]interface{})["attributes"].(map[string]interface{})["memory"]
		nToken := result.NextToken
		return instanceType, vcpu, memory, nToken, true, nil
	}
}

func main() {
	type dbInstanceType struct {
		Name   interface{}
		Vcpu   interface{}
		Memory interface{}
	}
	type dbInstanceTypeList []dbInstanceType
	var dbInstanceTypes dbInstanceTypeList

	var nextToken = ""
	for i := 0; ; {
		i++
		instanceType, vcpu, memory, nToken, next, err := getProduct(nextToken)
		if err != nil {
			nextToken = *nToken
			continue
		}
		if next != true {
			break
		}
		dbInstanceTypeInfo := dbInstanceType{
			Name:   instanceType,
			Vcpu:   vcpu,
			Memory: memory,
		}
		dbInstanceTypes = append(dbInstanceTypes, dbInstanceTypeInfo)
		nextToken = *nToken
	}
	results := make([]dbInstanceType, 0, len(dbInstanceTypes))
	encountered := map[interface{}]bool{}
	for i := 0; i < len(dbInstanceTypes); i++ {
		if !encountered[dbInstanceTypes[i]] {
			encountered[dbInstanceTypes[i]] = true
			results = append(results, dbInstanceTypes[i])
		}
	}

	maxLength := 0
	for i := 0; i < len(results); i++ {
		if len(results[i].Name.(string)) > maxLength {
			maxLength = len(results[i].Name.(string))
		}
	}

	instanceSpecs := ""
	for i := 0; i < len(results); i++ {
		NAME := results[i].Name.(string)
		MEMORY := results[i].Memory.(string)
		// In the API, db.r5.4xlarge has a memory of 192 GiB, but correctly it is 128 GiB
		if NAME == "db.r5.4xlarge" {
			MEMORY = "128 GiB"
		}
		MEMORY = strings.Replace(MEMORY, " GiB", "", 1)
		s, _ := strconv.ParseFloat(MEMORY, 64)
		s = s * (1024 * 1024 * 1024)
		MEMORY = strconv.FormatFloat(s, 'f', 0, 64)
		if len(NAME) < maxLength {
			NAME = NAME + strings.Repeat(" ", maxLength-len(NAME))
		}
		VCPU := results[i].Vcpu.(string)
		SPACE := strings.Repeat(" ", 4-len(VCPU))
		instanceSpec := "    " + NAME + " = {cpu_cores = " + VCPU + "," + SPACE + "memory = " + MEMORY + "}\n"
		instanceSpecs = instanceSpecs + instanceSpec
	}

	content := []byte(
		"locals {\n" +
			"  instance_types = {\n" +
			instanceSpecs +
			"  }\n" +
			"}\n",
	)
	ioutil.WriteFile(os.Args[len(os.Args)-1], content, os.ModePerm)
}
