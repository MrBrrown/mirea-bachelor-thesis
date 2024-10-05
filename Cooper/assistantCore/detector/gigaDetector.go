package detector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	t "example.com/coomper/assistantCore/commands"
)

type GigaDetector struct {
	token Token
}

var authURL = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
var requestURL = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"

var promptTemplate = `У меняесть такой набор команд:
0 Информация об оборудовании <Имя оборудования>
1 Какой процесс выполняет оборудование <>
2 Информация по текущему процессу <>
3 Параметры в текущем процессе <>
4 Информация по параметру <> 
5 Значение параметров <>
6 Неизвестный запрос <>

Каждый пользователь может задать вопрос в своей манере и не точно так как я описал выше, твоя задача классифицировать запрос который я тебе отправлю. В твоем ответе должен быть только один номер из списка.  В ответ ты должен прислать сообщение в формате json:
{"command": номер из списка}
Пример: 
Запрос: Привет! Что это за станок?
Ответ от тебя:{"command": 0}
Запрос: Можешь подробно расскакзать что сейчас происходит в станке?
Ответ от тебя:{"command": 1}
Запрос: Какое значение имеет пятый параметр из списка?
Ответ от тебя: {"command": 5}
Запрос: Какая сегодня погода?
Ответ от тебя: {"command": 6}

Вот мой запрос: %s`

func (d *GigaDetector) Process(command string) (t.Command, error) {
	type result struct {
		Choices []struct {
			FinishReason string `json:"finish_reason"`
			Message      struct {
				Content struct {
					Command int `json:"command"`
				} `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	var err error

	if !d.token.IsActive(time.Now().Unix()) {
		d.token, err = requestUid()
		if err != nil {
			return t.Unknown, err
		}
	}

	promt := fmt.Sprintf(promptTemplate, command)
	body := fmt.Sprintf(`{
		"model": "GigaChat",
		"messages": [
			{
			"role": "user",
			"content": "%s"
			}
		],
		"temperature": 1,
  		"top_p": 0.1,
  		"n": 1,
  		"stream": False,
  		"max_tokens": 512,
  		"repetition_penalty": 1
	}`, promt)

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return t.Unknown, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", d.token.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return t.Unknown, err
	}
	defer resp.Body.Close()

	res := result{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return t.Unknown, err
	}

	if len(res.Choices) > 0 {
		return t.Command(res.Choices[0].Message.Content.Command), nil
	}

	return t.Unknown, nil
}

func requestUid() (Token, error) {
	payload := strings.NewReader("scope=GIGACHAT_API_PERS")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, authURL, payload)
	if err != nil {
		return Token{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("RqUID", "6f0b1291-c7f3-43c6-bb2e-9f3efb2dc98e")
	req.Header.Add("Authorization", "Basic NzIyNDgyZWItMGI0Ni00MDlhLWI2MzctOTUyMzQxNTE5ZGVkOjIyZWUyMDVhLTk0ODUtNDM1NC04MzY0LTA5YTNjZTEyMDY4ZA==")

	res, err := client.Do(req)
	if err != nil {
		return Token{}, err
	}
	defer res.Body.Close()

	token := Token{}
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return Token{}, err
	}

	return token, nil
}
