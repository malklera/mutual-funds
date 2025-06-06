# mutual-funds
Track mutual funds yield daily

Download the page with and without JavaScript disable, so I have an idea of the
differences.
Most of the content is dynamically generated, will have to use chromedp, colly do not cut it.

Make an example where I download the whole page as a html.
Better to just use devtools and manually download the whole page to get an idea
of which elements I want to download.

Another where I download the specific content that I want as an html.
This is not necessary

I tried to get the specifics url from each element of the table from:
https://www.santander.com.ar/personas/inversiones/informacion-fondos#/
But I run into problems, only the first url gets saved, maybe in the future I will try
by doing it from chromedp instead of strait js from the console of devtools.

New approach:
Manually get the id, name, and complete url of each mutual fund available on the web, 
put all that data to a .json file, and make the program read from it.
Discard this, i do not care about the id, i got manually the name and url of each
mutual fund, and put it on a []byte variable inside of baseFunds.go, this will take
care of creating the funds.json file to be used by the program

Make a program to download the specific content that I want instead of html from each page.
Only get the first result I get, some of the field are repeated below the one I want.
Which data I want?
Name
Url
Risk
Value []float64
I have to check that the name I give and the one display on the page are the same.

Name display on the page:

<h1
  data-testid="titleDetailDesktop"
  class="sc-aXZVg jbWCFw sc-dISpDn dlSiwb"
>
Name fund
</h1>

Compare that the "Name fund" is the same of the one from the json,
if the website change the relation url/fund I need to know.
For now it it show a log.Fatalf()
TODO: In the future i want to be taken to the menu to modify mutual funds, i still
do not have this implemented
Only show something if they are not the same.

Risk? But only the first time, this do not change, or did it change? i do not know,
I will check each time, if the risk is different from the json, i update it, deal
with the risk not existing the first time i run the program
<p
  data-testid="fundRiskDetailName"
  class="sc-aXZVg hLnzCR sc-gRtvSG bydNqC"
>
  Type
</p>

Cuotapartes
<h3
  data-testid="currentShareValueType"
  class="sc-aXZVg dMxiJX"
>
  $ 14.563,697
</h3>
Get ride of "$ " and convert the rest to a float64, see how go deal with "." and ",",
this will appended to the slices if it already exist or initialize the first run
I just thought of something, i want to know when i get the value, the value slice has to
be a map instead, where i have date and value
I manually erase "$ " and ".", then reempalce the "," by a ".", then i can convert
the string to a float64

Another where I put the downloaded content to an actual data structure, check
what type of data structure I want to use.

I choose to make a couple of struct to dealt with all the info.
Then write it down as a .json file

TODO: when running saveValues() I have to check the date of the last element
of each fund, if the current date is the same as the last element of the Value slice
then I check if they are different I update that and not append a new one.
Where do i run this check??
I want  to run this program each time i boot my pc, then run it agan after the
market close, but some day i may not be able to do so, think how and where do the
checks.

Allow the program to be run with and without arguments, so I can do different things.
+ $ mutual-funds/main.go
    With no argument it will show the menu to the user
+ $ mutual-funds/main.go -u
    Without arguments -u it will retrieve the data of the day.

Make a simple menu with the options i will want the program to give the user.
First menu will be

+ My Funds.
+ All Funds.

Both will call upon the following menu, the difference will be that "My Funds"
need its own .json file where i have a list of funds, only the name, when i call
upon the functions below i can do it with either with the slice of funds or an empty slice
if the slice is empty the options below act upon the whole funds.json, if the slice is
non empty show and export will go through funds.json and only grab the ones in the
slice, on modify instead it will modify my-funds.json instead

DONE: The same way i create the first .json file, i have to create this myFunds.json
file the first time empty

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
+ Modify mutual fund.
  + Add mutual fund.
  + Remove mutual fund.
  + Modify mutual fund.

Add the capability to save data persistently, update the data, and retrieve it.
This is done, what i have to do is actually showing the data, but first finish the menu
and its options.
DONE: menu()

Work on each operation for the data, show, export, modify
showData: DONE

modifyData:
Make a menu for this  part, put it on menu.go
  What operations can i do here?
  myFunds.json
    Add a fund.
    Delete a fund.
    Modify a fund.
      Change name
      Change shares

  funds.json
    Add a fund.
    Delete a fund.
    Modify a fund.
      Change name
      Change url

The other fields the user is not allowed to touch
      
Download the data of all mutual funds, that way i can keep track of which one yield
the best, keep separate from the ones the user actually own.
DONE


TODO: Check if the cuantity of share i own increase with time or what

TODO: Manage all the errors, i want to go back a menu or retry on all of then,
leave this for last

After all that work, read if it is possible to save login data to compare my tracking
with the one generate by the website, only if I can encrypt my login data, and use it
safely.
