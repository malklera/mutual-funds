# mutual-funds
Track mutual funds yield daily

Donwload the page with and without JavaScript disable, so I have an idea of the
differences.
Most of the content is dynamically generated, will have to use chromedp, colly do not cut it.

Make an example where I download the whole page as a html.
Better to just use devtools and manually download the whole page to get an idea
of which elements I want to download.

Another where I download the specific content that I want as an html.
This is not necessary

I tried to get the specifics url from each element of the table from:
https://www.santander.com.ar/personas/inversiones/informacion-fondos#/
But i run into problems, only the first url gets saved, maybe in the future i will try
by doing it from chromedp instead of strait js from the console of devtools.

New approach:
Manually get the id, name, and complete url of each mutual fund available on the web, 
put all that data to a .json file, and make the program read from it.

Make a program to download the specific content that I want instead of html from each page.
Only get the first result i get, some of the field are repeated below the one i want.
Which data i want?
I have to check that the name i give and the one display on the page are the same.
Name of fund from json file:
Read throught the list of keys from the json, each key is the name of one fund.
Name display on the page:
<h1
data-testid="titleDetailDesktop"
class="sc-aXZVg jbWCFw sc-dISpDn dlSiwb"
>
Name fund
</h1>
Compare that the "Name fund" is the same of the key from the json, i actually access
each url by the id, if the website change the relation id/fund i need to know.

Risk? But only the first time, this do not change
<p
  data-testid="fundRiskDetailName"
  class="sc-aXZVg hLnzCR sc-gRtvSG bydNqC"
>
  Type
</p>



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
