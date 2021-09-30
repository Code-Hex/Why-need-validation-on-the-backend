## Why need validation on the backend?

This project is an appendix for my blog [article (about validation)](https://codehex.dev/notes/about_validation).

This application has a scenario.

> Once upon a time, There was a war between the banana faction and the orange faction. One day, somebody said, “I’ve developed the voting application to prove which is better fruits!”

Now, which one will you vote for? Voting for any other fruit is a sin.

<img width="375" alt="screenshot" src="https://user-images.githubusercontent.com/6500104/135491289-bef4db50-103e-4312-a31b-59c9d0474f70.png">

## How to setup and runs this?

```
go run main.go
```

and go to `http://127.0.0.1:1323` with your browser.

## Cheat

You can vote for other fruit this way.

```
$ curl 'http://127.0.0.1:1323/vote' -H 'Content-Type: application/x-www-form-urlencoded' --data-raw 'fruit=apple'
```

## Refs

- https://stackoverflow.com/a/162579