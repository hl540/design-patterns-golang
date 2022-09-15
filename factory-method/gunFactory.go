package factory_method

import "fmt"

// 简单工厂，根据gunType返回不太的IGun实现
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
