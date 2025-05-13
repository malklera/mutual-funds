# mutual-funds
Track  mutual funds yield daily

Donwload the page with and without JavaScript disable, so I have an idea of the
differences.

Make a example where I download the whole page as an html.

Another where I download the specific content that I want as an html.

Another where I download the specific content that I want instead of html.

Another where I put the downloaded content to an actual data structure, check
what type of data structure I want to use.

Allow the program to be run with and without arguments, so I can do different things.
+ $ mutual-funds/main.go
    Without arguments it will retrieve the data of the day.
+ $ mutual-funds/main.go -m
    With the argument -m it will show the menu to the user

Make a simple menu with the options i will want the program to give the user.
+ Show data.
    + Think about this, how do i want the data to be show, as a table of the
    history value of each element or what.
+ Export data.
    + json
    + pdf
    + html?
    + csv
    + txt
    + spread sheet
+ Add mutual fund.
+ Modify mutual fund.
+ Remove mutual fund.

Add the capability to save data persistently, update the data, and retrieve it.

Download the data of all mutual funds, that way i can keep track of wich one yield
the best, keep separate from the ones the user actually own.

After all that work, read if it is posible to save login data to compare my tracking
with the one generate by the website, only if I can encrypt my login data, and use it
safely.
