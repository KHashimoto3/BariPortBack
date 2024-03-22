package bariport

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

//dynamo db接続関数
func connect(dbName string) dynamo.Table {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})
	return db.Table(dbName)
}

//各テーブルのデータ構造
//companiesテーブル
type Company struct {
	Id 		string `dynamo:"id"`
	Name 	string `dynamo:"name"`
}

//projectsテーブル
type Project struct {
	Id 			string `dynamo:"id"`
	CompanyId 	string `dynamo:"companyId"`
	Name 		string `dynamo:"name"`
	Description string `dynamo:"description"`
}

//companyのデータを取得
func GetCompany(companyId string) (Company, error) {
	table := connect("companies")
	var company Company
	err := table.Get("id", companyId).One(&company)
	if err != nil {
		return company, err
	}
	return company, nil
}

func GetProjects() ([]Project, error) {
	table := connect("projects")
	var projects []Project
	err := table.Scan().All(&projects)
	if err != nil {
		return projects, err
	}
	return projects, nil
}
