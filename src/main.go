package main

import (
	"HawkEye-Go/src/Engine"
	"HawkEye-Go/src/SqlPaser"
	"fmt"
	"net/http"
)

func handleBlackData(w http.ResponseWriter, r *http.Request) {
	// Handle incoming black data
}

func handleWhiteData(w http.ResponseWriter, r *http.Request) {
	// Handle incoming white data
}

func handlePredictionData(w http.ResponseWriter, r *http.Request) {
	// Handle data for prediction
}

func main() {
	SqlPaser.TestLexer(nil)
	SqlPaser.TestParseInsertStatement(nil)
	SqlPaser.TestParseSelectStatement(nil)
	nb := Engine.NewNaiveBayes()

	// 训练
	featuresWhite := Engine.ExtractFeatures("SELECT name FROM users WHERE id = 1")
	nb.Train(featuresWhite, "White")

	featuresBlack := Engine.ExtractFeatures("DROP TABLE users")
	nb.Train(featuresBlack, "Black")

	err := nb.SaveToFile("naive_bayes_model.gob")
	if err != nil {
		panic(err)
	}

	nc, err := Engine.LoadModelFromFile("naive_bayes_model.gob")
	if err != nil {
		panic(err)
	}

	// 预测
	predictFeatures := Engine.ExtractFeatures("SELECT name FROM admins")
	result := nc.PredictProbability(predictFeatures)
	fmt.Println("是黑数据的概率 %v", result)
	// 输出预测结果
	http.HandleFunc("/blackdata", handleBlackData)
	http.HandleFunc("/whitedata", handleWhiteData)
	http.HandleFunc("/predict", handlePredictionData)

	http.ListenAndServe(":8080", nil)
}
