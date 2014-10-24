gophraser
=========

Microservice to generate random adjective-animal phrases.

Available at http://gophraser.herokuapp.com/

This is a Go port of the Java original at http://github.com/jostly/phraser

Usage
=====

Output from the service will depend on the Accept header. The following media types can be produced:

`text/plain`: The phrase as text, with no terminating line feed

`application/json`: The phrase as a json object

        {
            "adjective": "<some adjective>",
            "animal": "<some animal>"
        }

`text/html`: A page designed for viewing in browsers

Alliterative phrase
-------------------

The words start with the same letter, randomly chosen. Access at base URL: `/`, example: http://gophraser.herokuapp.com/


Random phrase
-------------

Each word is randomly chosen independently. Access at URL `/random`, example: http://gophraser.herokuapp.com/random


Phrase starting with specific letter
------------------------------------

Each word is randomly chosen among those starting with the specified letter. Access at URL `/{letter}`, example: http://gophraser.herokuapp.com/p

There are currently no animals or adjectives starting with the letter x, so the service will not accept that as a starting letter.

Heroku deploy notes
===================

To deploy on Heroku, you need to create the app with a buildpack:
`heroku create -b https://github.com/kr/heroku-buildpack-go.git`
Then push to heroku like always:
`git push heroku master`
