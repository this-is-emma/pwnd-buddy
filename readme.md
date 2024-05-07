# PWND-BUDDY

PWND-BUDDY is a tool to help you checked if your email address is part of a previous leak. It is super easy to use,
here are a few steps to get you started:

- Create an `.env` file with your HaveIBeenPwnd API Key

- Build the tool:

`go build main.go`

- run the execurtable

`./main -email="<YOUR-EMAIL@ADDRESS.COM>" -s`

Breakdown of the command:

- *-email* is a required flag, enter the email address you want to check inside quotes.
- *-s* tells the tool to save the result in a .txt file in the root directory.
If you don't want to save, simply omit the flag

Here is an example of the output:

![alt text](<cli output.png>)

