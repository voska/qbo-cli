Output demo.gif

Set Shell "bash"
Set FontSize 14
Set Width 900
Set Height 500
Set Theme "Catppuccin Mocha"
Set Padding 20

Type "qbo list invoices --where \"Balance > '0'\" --sandbox"
Sleep 500ms
Enter
Sleep 3s

Sleep 1s

Type "qbo report profit-and-loss --start-date 2025-01-01 --sandbox"
Sleep 500ms
Enter
Sleep 3s

Sleep 1s

Type "qbo get customer 1 --sandbox --json"
Sleep 500ms
Enter
Sleep 3s
