# Equity Announcement Notifier  

## Description  
A Go-based tool that monitors websites for equity-related updates and sends notifications to a Discord channel.  

## Features  
- Tracks website changes  
- Sends Discord notifications  

## Installation  
```sh
git clone https://github.com/yourusername/equity-notifier.git  
cd equity-notifier  
go mod tidy  
go run main.go  
```  

## Configuration  
Edit `config.json`:  
```json
{
  "websites": ["https://example.com"],
  "discordWebhook": "YOUR_WEBHOOK_URL",
  "checkInterval": 60
}
```  

## License  
MIT
