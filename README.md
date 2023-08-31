# long-term-project

## ideas to aplay during development
* skip volume data in basic xtb import for some time, and then add it along with db migration

## 2023-07-27
inspired by this video:
https://www.youtube.com/watch?v=VTCP8RzBx6w
I decided to change the scope of the project to an application in which I will be able to monitor my trading on the forex market.  
**Little History:**  
Becoming a professional trader is my unfulfilled dream. I had a few opportunities to do it professionally, but my results never allowed me to make a living out of it.
## new project description
so as my history is as is, and Piotrak at movie said that we should do something that we like, i decided to do something that i like.
i would like to have some tool that will help me to monitor my trading on forex market, yes i still trade, but only for fun with very small amount of money.
other thing is that i not what start with POC with bad code, i want start with "kind of" good code (apply my current knowledge) and simulate normal development process with new features, refactoring, etc.
It turned out that writing poor quality code intentionally is difficult and can be very easy to bend, for example, using only http get :D

### First Epic
As a trader i would like to take csv from  my broker and import it to my app, so i can see my trades in my app. Want to be able see some basic statistics like:
* how many trades i have
* how many trades i won and lost
* what is my win ratio
* what is my average win and average lost in pips
* what instrument i trade and what is my win ratio for each instrument
* time of day when i trade and what is my win ratio for each time of day


## loose thoughts and ideas
* road map
* trunk base development with feature flags and somehow simulate parallel development
* kubernetes
* some cloud provider (aws, gcp, azure)
* maybe one microservice in other language (python, nodejs, rust)

## 2021-07-10
**TL,DR:**   
i am starting with POC (very messy code) and then i will refactor it to good quality code(hopefully).


This is a _long-term project_ to learn and practice software development on on-line shop example, like ebay, amazon, allegro, etc.
i think set up some kind of app is much simpler then maintain it, add new features, refactor it, etc.
so i will try to do it... try to imagine that we do not have any on-line shop (we have back to 90` :D ) and we want to create one.  
first with very basic features to allow to connection some sellers and buyers. will try to do this as POC - so some quick way to deliver something.
then i will try to refactor it to good quality code. not sure want this exactly mean...  
but for now i know that most challenging part is developing app witch is running on production, and i want to try to simulate it.

