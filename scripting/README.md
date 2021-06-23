# Scripting technical tests

At Adfinitas, we love Python for its powerful simplicity. We primarily use it for "data pipelines" in Google Cloud Dataflow and some Google Cloud Function.

## Tests

For that exercise, we will tests your knowledge to write efficient python and nodejs code, to refractor an existing codebase and read online docs.

Be careful of everything : we let bugs, memory leaks and bad codes in the code. Pay attention to every potential cases. Some changes (not written here) may be expected to be done for maintainability and readability.

## Exercises

We advice you to start from the exercise Python 1 to 3 and NodeJS 1 to 3 as they are sorted by difficulty.

### Python - 1. Write the most valuable customer in a CSV file

We already did a function `get_most_valuable_customer` (in the `lib.py` file). You will have to write a main which will call `get_most_valuable_customer` function and store its result in a CSV file as `most_valuable_customer.csv`.

### Python - 2. Encrypt emails in a CSV file

Write a small executable script which opens `some-emails.csv`, read its content and encrypt the `email` column to another file `some-emails-out.csv` with a `email_md5` column appended.

### Python - 3. Parameterize your "encrypt emails in CSV file" script

You will have to add parameters to the past script with:
 - the input filename
 - the output filename, default to `${input_filename without ".csv"}-out.csv`
 - the column to read, default to `email`
 - the column to add, default to `${column_to_read}_${algorithm_to_use}`
 - the encryption algorithm to use, default to `md5`, could be `md5` or `sha256` for example ?

### NodeJS - 1. Small fibonacci HTTP server

You will have to create a server with one route `/fibonacci` which give the Fibonacci sequence result of the `n` query parameter.

### NodeJS - 2. Proxy requests to IMDb

You will have to add routes to forward requests to the [IMDb API](https://rapidapi.com/apidojo/api/imdb8/).

The `/fibonacci` route must still work !

 - Responses must be wrapped in a `imdb` object.
 - Invalid requests must returns 404, find the best way to handle that case.

### NodeJS - 3. Catch invalid payloads

Write a route `/optin-lead` which receive payloads as :
```json
{
    "email_md5": "string ; format:md5 ; optional_without:email",
    "email": "string ; format:email ; optional_without:email_md5",
    "phone": "string ; format:phone",
    "firstname": "string",
    "lastname": "string",
    "sexe": "char ; enum:H,F,O, ; optional",
    "note": "int ; min:0 ; max:10"
}
```
You must validate incoming payload and return either `200 OK` or `400 Bad Request` with the errors you did detect on the payload. Pay attention on every cases.

## Haven't you miss anything ?

In all development there is one important thing not done here, find what and try to do it for that project.
