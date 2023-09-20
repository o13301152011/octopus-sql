package Engine

import (
	"encoding/gob"
	"os"
)

type NaiveBayes struct {
	classCounts        map[string]int
	featureValueCounts map[string]map[string]map[string]int
	totalSamples       int
}

func NewNaiveBayes() *NaiveBayes {
	return &NaiveBayes{
		classCounts:        make(map[string]int),
		featureValueCounts: make(map[string]map[string]map[string]int),
		totalSamples:       0,
	}
}

// 使用数据训练分类器
func (nb *NaiveBayes) Train(features map[string]string, label string) {
	nb.totalSamples++
	nb.classCounts[label]++

	for feature, value := range features {
		if _, exists := nb.featureValueCounts[feature]; !exists {
			nb.featureValueCounts[feature] = make(map[string]map[string]int)
		}
		if _, exists := nb.featureValueCounts[feature][value]; !exists {
			nb.featureValueCounts[feature][value] = make(map[string]int)
		}
		nb.featureValueCounts[feature][value][label]++
	}
}

// 预测给定特征数据的类别
func (nb *NaiveBayes) Predict(features map[string]string) string {
	maxProb := float64(-1)
	bestClass := ""

	for class, classCount := range nb.classCounts {
		prob := float64(classCount) / float64(nb.totalSamples)
		for feature, value := range features {
			if _, exists := nb.featureValueCounts[feature]; !exists {
				continue
			}
			if _, exists := nb.featureValueCounts[feature][value]; !exists {
				continue
			}
			prob *= float64(nb.featureValueCounts[feature][value][class]+1) / float64(classCount+len(nb.featureValueCounts[feature]))
		}
		if prob > maxProb {
			maxProb = prob
			bestClass = class
		}
	}

	return bestClass
}

func (nb *NaiveBayes) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := gob.NewEncoder(file)
	return encoder.Encode(nb)
}

// 预测给定特征数据的类别的概率
func (nb *NaiveBayes) PredictProbability(features map[string]string) float64 {
	blackProb := float64(nb.classCounts["Black"]) / float64(nb.totalSamples)
	whiteProb := float64(nb.classCounts["White"]) / float64(nb.totalSamples)

	for feature, value := range features {
		if _, exists := nb.featureValueCounts[feature]; !exists {
			continue
		}
		if _, exists := nb.featureValueCounts[feature][value]; !exists {
			continue
		}
		blackProb *= float64(nb.featureValueCounts[feature][value]["Black"]+1) / float64(nb.classCounts["Black"]+len(nb.featureValueCounts[feature]))
		whiteProb *= float64(nb.featureValueCounts[feature][value]["White"]+1) / float64(nb.classCounts["White"]+len(nb.featureValueCounts[feature]))
	}

	// 返回被分类为"Black"的概率
	return blackProb / (blackProb + whiteProb)
}

func LoadModelFromFile(filename string) (*NaiveBayes, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := gob.NewDecoder(file)
	nb := &NaiveBayes{}
	err = decoder.Decode(nb)
	if err != nil {
		return nil, err
	}
	return nb, nil
}
