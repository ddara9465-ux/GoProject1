package telegram

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// SendNewAppointmentNotify - отправляет уведомление админу в Telegram
// Запускается в горутине, не блокирует пользователя
func SendNewAppointmentNotify(userID int, date, master, service string) {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_ADMIN_CHAT_ID")

	if botToken == "" || chatID == "" {
		log.Printf("Telegram не настроен: пропущена отправка")
		return
	}

	// Формируем сообщение
	message := fmt.Sprintf(`НОВАЯ ЗАПИСЬ! | ID клиента: %d | Желаемая дата: %s | Мастер: %s | Услуга: %s | Время записи: %s`,
		userID, date, master, service, time.Now().Format("15:04:05"))

	//Кодируем сообщение для безопасной передачи в URL
	encodedMessage := url.QueryEscape(message)

	//Формируем URL
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		botToken, chatID, encodedMessage)

	// Отправляем запрос (с таймаутом 30 секунд)
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		log.Printf("[ГОРУТИНА] Ошибка отправки в Telegram: %v", err)
		return
	}
	defer resp.Body.Close()

	// Проверяем результат
	if resp.StatusCode == 200 {
		log.Printf("[ГОРУТИНА] Уведомление отправлено в Telegram")
	} else {
		log.Printf("[ГОРУТИНА] Telegram ошибка: статус %d", resp.StatusCode)
	}
}
