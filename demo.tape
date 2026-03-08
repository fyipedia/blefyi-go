# blefyi-go — Go API Client Demo
Output demo.gif

Set FontSize 20
Set Width 1000
Set Height 580
Set Shell "zsh"
Set Theme "Catppuccin Mocha"
Set TypingSpeed 30ms
Set Padding 24

Type "# BLEFYI Go API Client"
Enter
Sleep 0.5s

Type "curl -s 'https://blefyi.com/api/search/?q=beacon' | jq '.results[:3]'"
Enter
Sleep 3s

Type "curl -s 'https://blefyi.com/api/chip/nrf52832/' | jq '{name, slug}'"
Enter
Sleep 3s

Sleep 1s
