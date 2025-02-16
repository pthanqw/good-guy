package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/openuniland/good-guy/configs"
	"github.com/openuniland/good-guy/external/facebook"
	"github.com/openuniland/good-guy/external/types"
	"github.com/rs/zerolog/log"
)

type FacebookUS struct {
	cfg *configs.Configs
}

func NewFacebookUseCase(cfg *configs.Configs) facebook.UseCase {
	return &FacebookUS{cfg: cfg}
}

func (us *FacebookUS) SendMessage(ctx context.Context, id string, message interface{}) error {
	data := map[string]interface{}{
		"recipient": map[string]string{
			"id": id,
		},
		"message":        message,
		"messaging_type": "MESSAGE_TAG",
		"tag":            "ACCOUNT_UPDATE",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Error().Err(err).Msg("error marshal data")
		return err
	}

	url := fmt.Sprintf("https://graph.facebook.com/v14.0/me/messages?access_token=%s", us.cfg.FBConfig.FBVerifyToken)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Err(err).Msg("error create new request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error send request")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error().Err(err).Msg("error response from Facebook API")

		return errors.New("error response from Facebook API")
	}

	return nil
}

func (us *FacebookUS) SendTextMessage(ctx context.Context, id string, text string) error {
	message := map[string]string{
		"text": text,
	}

	return us.SendMessage(ctx, id, message)
}

func (us *FacebookUS) SendImageMessage(ctx context.Context, id string, url string) error {
	message := map[string]interface{}{
		"attachment": map[string]interface{}{
			"type": "image",
			"payload": map[string]string{
				"url": url,
			},
		},
	}

	return us.SendMessage(ctx, id, message)
}

func (us *FacebookUS) SendButtonMessage(ctx context.Context, id string, input *types.SendButtonMessageRequest) error {

	message := map[string]interface{}{
		"attachment": map[string]interface{}{
			"type": "template",
			"payload": map[string]interface{}{
				"template_type": "generic",
				"elements": []map[string]interface{}{
					{
						"title":     input.Title,
						"image_url": input.ImageUrl,
						"subtitle":  input.Subtitle,
						"buttons": []map[string]interface{}{
							{
								"type":                 "web_url",
								"url":                  input.Url,
								"title":                input.BtnText,
								"messenger_extensions": true,
								"webview_height_ratio": "tall",
							},
						},
					},
				},
			},
		},
	}

	return us.SendMessage(ctx, id, message)
}
