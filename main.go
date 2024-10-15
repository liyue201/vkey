package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	groth16_bn254 "github.com/consensys/gnark/internalx/backend/bn254/groth16"
	"os"
)

func LoadVk2(filename string) {
	d, err := os.ReadFile(filename + ".vk")
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(d)
	vk := groth16_bn254.VerifyingKey{}
	_, err = vk.ReadFrom(buffer)
	if err != nil {
		panic(err)
	}

	// 将验证密钥序列化为 JSON 格式
	vkJSON, err := json.MarshalIndent(&vk, "", "    ")
	if err != nil {
		panic(err)
	}
	// 将 JSON 数据写入文件
	err = os.WriteFile(filename+".json", vkJSON, 0644)
	if err != nil {
		panic(err)
	}

	// solidity
	b := bytes.Buffer{}
	vk.ExportSolidity(&b)
	os.WriteFile(filename+".sol", b.Bytes(), 0766)
}

func loadVk(filename string) (*groth16_bn254.VerifyingKey, error) {
	d, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(d)
	vk := groth16_bn254.VerifyingKey{}
	_, err = vk.ReadFrom(buffer)
	if err != nil {
		return nil, err
	}
	return &vk, err
}

func saveJson(vk *groth16_bn254.VerifyingKey, filename string) error {
	vkJSON, err := json.MarshalIndent(&vk, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, vkJSON, 0644)
}

func saveSolidity(vk *groth16_bn254.VerifyingKey, filename string) error {
	b := bytes.Buffer{}
	vk.ExportSolidity(&b)
	return os.WriteFile(filename, b.Bytes(), 0766)
}

func main() {
	vk := flag.String("vk", "", "verifying key")
	sol := flag.String("sol", "", "solidity verifier")
	json := flag.String("json", "", "verifying key in json format")

	flag.Parse()

	if *vk == "" {
		fmt.Println("verifying key not found")
		return
	}
	vkey, err := loadVk(*vk)
	if err != nil {
		fmt.Printf("loadVk: %v\n", err)
		return
	}
	if *sol != "" {
		err := saveSolidity(vkey, *sol)
		if err != nil {
			fmt.Printf("saveSolidity: %v\n", err)
		}
	}
	if *json != "" {
		err := saveJson(vkey, *json)
		if err != nil {
			fmt.Printf("saveSolidity: %v\n", err)
		}
	}
}
