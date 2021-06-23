# Data technical tests

At Adfinitas, we work intensively with data. We use Google Cloud BigQuery and GC Datastore as our primary datahouse.

## Tests

For that exercise, we will tests your knowledge to write efficient SQL queries and your capability to understand a database you didn't write.

You will find a SQLite database named `chinook.db`. Open it with your prefered tool. For example, we use JetBrains tools : Datagrip for SQL.

For every exercise, you have to give a SQL file with the query you wrote and a CSV file with the result. Please, give explicit names for each.

## Exercises

We advice you to start from the exercise 1 to 6 as they are sorted by difficulty.

### 1. Invoices of may 2013

You will have to query on `invoices` table to return all invoices made on may 2013.

### 2. Income for each month

Compute for each months the total income from `invoices` table ordered from the most recent to the less recent.

### 3. Top 5 lucrative months all time

We have benefits when our income is greater than 49.50$ (our hosting fees).

Thanks to the past query, find the top 5 most lucrative beneficial months.

### 4. Top 5 customers

Find the top 5 buyers (whom spent the most in our shop) all time who are not from USA ordered by the `Email`.
Return the `Email` and the `Total` (addition of all `Total` of its invoices).

### 5. Top 10 selling artists all time

Find the top 10 artists who sells the most tracks of all time.

### 6. Total of sold tracks and top selling artist per genre

For each genre, compute the total sold tracks and the top selling artist.
