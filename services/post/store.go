package post

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Loghadhith/cms/configs"
	"github.com/Loghadhith/cms/types"
	"github.com/Loghadhith/cms/utils"
)

type CreateFileRequest struct {
	Message   string `json:"message"`
	Content   string `json:"content"`
	Branch    string `json:"branch,omitempty"`
	Committer struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"committer"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) PostContent(post types.PostPayload) error {

	//User data is fetched
	user, err := utils.GetUserByEmail(s.db, post.Email)

	if err != nil {
		log.Println("get user error, ", err)
		return err
	}

	//sample url
	str := RawUrlGenerate(post, user.Username)

	//data insertion
	_, er := s.db.Exec("INSERT INTO content (uid,repo,path,type,url) VALUES ($1,$2,$3,$4,$5);",
		user.ID, post.Repo, post.Path, post.Type, str)

	if er != nil {
		log.Println("err: ", er)
		return er
	}

	//repo creation
	createUrl := fmt.Sprintf("%s/repos", configs.Envs.CreateApiGithub)

	repoData := map[string]interface{}{
		"name":    post.Repo,
		"private": false,
	}

	repoJSON, err := json.Marshal(repoData)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
		return err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", createUrl, bytes.NewBuffer(repoJSON))
	if err != nil {
		log.Fatal("Error creating request: ", err)
		return err
	}

	//check here before commit
	req.Header.Add("Authorization", "token "+user.Pat)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error creating repository: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		fmt.Println("Repository created successfully!")
	} else {
		log.Printf("Failed to create repo. Status: %s", resp.Status)
	}

	//data upload
	fileContentBase64 := base64.StdEncoding.EncodeToString([]byte(post.Data))

	uploadUrl := fmt.Sprintf("%s/repos/%s/%s/contents/%s", configs.Envs.PutApiGithub, user.Username, post.Repo, post.Path)

	createFileRequest := CreateFileRequest{
		Message: "INIT",
		Content: fileContentBase64,
		Branch:  "main",
	}

	// Committer info
	createFileRequest.Committer.Name = user.Username
	createFileRequest.Committer.Email = user.Email

	requestBody, err := json.Marshal(createFileRequest)
	if err != nil {
		log.Fatal("Error marshalling JSON: ", err)
    return err
	}

	req, err = http.NewRequest("PUT", uploadUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal("Error creating request: ", err)
    return err
	}

	req.Header.Add("Authorization", "token "+user.Pat)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")

	resp, err = client.Do(req)

	log.Println(resp)
	if err != nil {
		log.Fatal("Error making API request: ", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}

func RawUrlGenerate(post types.PostPayload, un string) string {
	url := configs.Envs.RawUrl

	rawurl := fmt.Sprintf("%s/%s/%s/refs/heads/main/%s", url, un, post.Repo, post.Path)
	return rawurl
}

func (s *Store) PostContentOnExistRepo(post types.PostPayload) error {
	return nil
}

func (s *Store) GetPostedData(mail types.ReqBody) ([]string, error) {

  log.Println(mail)
  log.Println("double")
	r, er := utils.GetPostedData(s.db, mail)
  log.Println("getpostedcompleted")
	log.Println(r)
	log.Println("Entered store function ")
	if er != nil {
		log.Println("error: ", er)
		return nil, er
	}

  return r,nil

	// Need to continue the fetch process of the repo names under the utils
	// return []string{"apple", "banana", "cherry"}, nil
}
