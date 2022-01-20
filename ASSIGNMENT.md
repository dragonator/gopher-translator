# Gopher Translator Service

## Overview
-----------

* Gophers are friendly creatures, but it is not that easy to communicate with them.
* The reason being is that they have their own language called Gopher language, and they do not understand English.
* Your mission will be to create a program that could bridge the speaking gap between humans and gophers and translates English words and sentences into
equivalent in Gopher language.

## Assignment
-------------
* Create a program that starts an HTTP server.
* This server should be able to translate English words into words in the Gopher language.
* (OPTIONAL, but a huge plus) Write tests, but without using any third party libraries. The standard library provides everything you need.

### NOTE
> * The solution should be implemented in Golang.
> * It is necessary for your code to compile.
-------------

### Gopher Language Spec

The language that the gophers speak is a modified version of English and has a few simple rules:
* If a word starts with a vowel letter, add prefix g to the word:

  > apple -> gapple  
  > ear -> gear  
  > oak -> goak  
  > user -> guser

* If a word starts with the consonant letters xr , add the prefix ge to the beginning of the word:

  > xray -> gexray

* If a word starts with a consonant sound, move it to the end of the word and then add ogo suffix to the word. Consonant sounds can be made up of multiple
consonants i.e., a consonant cluster:

  > chair -> airchogo

* If a word starts with a consonant sound followed by qu , move it to the end of the word, and then add ogo suffix to the word:

  > square -> aresquogo

------------------------------------------------------------------------------------------------
### NOTE
> Please do not confuse the gophers as they do not understand shortened versions of words or apostrophes. So do not use words like don’t , shouldn’t , etc.
Even translated they still will not understand you so skip them in your solution.
------------------------------------------------------------------------------------------------

### Server Endpoints
* POST /word

    * Given an English word, the server should return the word’s translation in Gopher language.
    * It should accept JSON data in the format:
    
            {
                "english_word": "<a single English word>"
            }

    * And should return JSON data in the format:

            {
                "gopher_word": "<translated version of the given word>"
            }

* POST /sentence

    * Given an English sentence (in which each whitespace separated sequence counts as a single word), the server should return the sentence translation in Gopher language.
    * It should accept JSON data in the format:

            {
                "english_sentence": "<sentence of English words>"
            }

    * And should return JSON data in the format:

            {
                "gopher_sentence": "<translated version of the given sentence>"
            }

    * Assume that every sentence ends with dot, question or exclamation mark.

* (OPTIONAL) GET /history

    * Should return each English word or sentence that was given to the server from the time the server was started along with its translation in Gopher
language.
    * The output should look like the following:

            {
                "history": [
                    {
                        "apple":"gapple"
                    },
                    {
                        "user": "guser"
                    },
                    {
                        "my":"ymogo"
                    },
                    ...
                ]
            }

    * The returned array should be ordered alphabetically ascending by the English word/sentence.

------------------------------------------------------------------------------------------------
**IMPORTANT**
> Please implement an in-memory storage solution, e.g. a map will be just fine. DO NOT use any databases.
------------------------------------------------------------------------------------------------

### Deployment
* Your program should accept one command line argument —port which is the port that the server is running.
* Use Docker to configure the server to start as a container.
* (OPTIONAL) Build and run the container via docker-compose .

### Additional Details

* Present the project as a shareable git repository, e.g. (Github, Gitlab, Bitbucket, etc.).
* Add a README.md file that contains the project documentation which should also contain deployment instructions.

Good Luck! Always remember to have fun writing code in Go ;).