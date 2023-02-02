package main

/*
Данная задача ориентирована на реальную работу с данными
в формате JSON. Для решения мы будем использовать
справочник ОКВЭД (Общероссийский классификатор видов
экономической деятельности), который был размещен на
web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится
в файле structure-20190514T0000.json, а сами данные,
которые вам потребуется декодировать, содержатся в файле
data-20190514T0100.json.
Файлы размещены в репозитории на github.com:
https://github.com/semyon-dev/stepik-go/tree/master/work_with_json

Для того, чтобы показать, что вы действительно смогли
декодировать документ вам необходимо в качестве ответа
записать сумму полей global_id всех элементов,
закодированных в файле.
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Global struct {
	GlobalID int `json:"global_id"`
}

func main() {
	// urlDownloadTest := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/structure-20190514T0000.json"
	urlDownload := "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	resp, err := http.Get(urlDownload)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	abs_fname, err := ioutil.ReadAll(resp.Body)

	var totalGlobals []Global
	json.Unmarshal(abs_fname, &totalGlobals)

	var result int
	for _, val := range totalGlobals {
		result += val.GlobalID
	}
	fmt.Println(result)

	// for index, val := range lines[0] {
	// 	if val == "0" {
	// 		fmt.Println("Result:", index+1)
	// 	}
	// }
}
